package main

import (
	"bytes"
	"fmt"
	"go/format"
	"html/template"
	"os"
	"path/filepath"

	"github.com/walteh/schema2go/pkg/generator/testcases"
)

func generateTestCases() {
	tcs, err := testcases.LoadTestCases()
	if err != nil {
		panic(err)
	}
	funcNames := []string{}

	for _, tc := range tcs {
		rawSchemaModelFunc, err := GenerateExpectedSchemaModel(&tc)
		if err != nil {
			panic(err)
		}
		funcNames = append(funcNames, rawSchemaModelFunc)

		schemaModelMockFunc, err := GenerateSchemaModelMockFunc(&tc)
		if err != nil {
			panic(err)
		}
		funcNames = append(funcNames, schemaModelMockFunc)

		goCodeFunc, err := GenerateGoCodeFunc("checkGoCode", &tc)
		if err != nil {
			panic(err)
		}
		funcNames = append(funcNames, goCodeFunc)

	}

	buf := bytes.NewBuffer(nil)
	buf.WriteString("package generator_test\n\n")
	buf.WriteString("import (\n")
	buf.WriteString("\"testing\"\n")
	buf.WriteString("\"github.com/walteh/schema2go/pkg/generator/testcases\"\n")
	buf.WriteString("\"github.com/walteh/schema2go/pkg/generator\"\n")
	buf.WriteString("\"github.com/walteh/schema2go/gen/mockery\"\n")
	buf.WriteString("\"github.com/google/gnostic/jsonschema\"\n")
	buf.WriteString("\"github.com/walteh/schema2go/pkg/diff\"\n")
	buf.WriteString(")\n\n")

	// add mustLoadSchemaModel
	buf.WriteString("func mustLoadSchemaModel(t *testing.T, input string) *generator.SchemaModel {\n")
	buf.WriteString("	schema, err := generator.NewSchemaModel(input)\n")
	buf.WriteString("	if err != nil {\n")
	buf.WriteString("		t.Fatalf(\"Failed to parse schema: %v\", err)\n")
	buf.WriteString("	}\n")
	buf.WriteString("	return schema\n")
	buf.WriteString("}\n\n")

	for _, rawSchemaModelFunc := range funcNames {
		if rawSchemaModelFunc == "" {
			continue
		}
		buf.WriteString(rawSchemaModelFunc)
		buf.WriteString("\n\n")
	}

	data := buf.Bytes()

	// format the code
	formatted, err := format.Source(data)
	if err != nil {
		panic(fmt.Sprintf("Failed to format code: %v ... writing raw code: %s", err, string(data)))
	}

	if err := os.WriteFile(filepath.Join("..", "expectations_gen_test.go"), formatted, 0644); err != nil {
		panic(fmt.Sprintf("Failed to write file: %v", err))
	}
}

func GenerateGoCodeFunc(funcName string, tc *testcases.TestCase) (string, error) {
	if tc.GenerateExpectedGoCode() == "" {
		return "", nil
	}

	tmpl := `
		func {{ .Name }}_GoCode(t *testing.T) {
			tc := testcases.LoadAndParseTestCase("{{ .Name }}")

			{{.FuncName}}(t, tc.GenerateInput(), tc.GenerateExpectedGoCode())
		}
	`
	tmpld := template.Must(template.New("wantSchemaModel").Parse(tmpl))
	buf := bytes.NewBuffer(nil)
	err := tmpld.ExecuteTemplate(buf, "wantSchemaModel", map[string]any{
		"Name":     tc.Name(),
		"FuncName": funcName,
	})
	if err != nil {
		return "", err
	}
	return buf.String(), nil
}

func GenerateSchemaModelMockFunc(tc *testcases.TestCase) (string, error) {

	if tc.GenerateExpectedSchemaModelMockGoCode() == "" {
		return "", nil
	}

	tmpl := `
		func {{ .Name }}_SchemaModelMock(t *testing.T) {

			{{ .SchemaModelMockGoCode }}

			tc := testcases.LoadAndParseTestCase("{{ .Name }}")

			got := mustLoadSchemaModel(t, tc.GenerateInput())

			staticGot := generator.NewStaticSchema(got)

			staticWant := generator.NewStaticSchema(want)

			diff.RequireKnownValueEqual(t, staticWant, staticGot)
		}
	`
	tmpld := template.Must(template.New("wantSchemaModel").Parse(tmpl))
	buf := bytes.NewBuffer(nil)
	err := tmpld.ExecuteTemplate(buf, "wantSchemaModel", map[string]any{
		"Name":                  tc.Name(),
		"SchemaModelMockGoCode": tc.GenerateExpectedSchemaModelMockGoCode(),
	})
	if err != nil {
		return "", err
	}
	return buf.String(), nil

}

func GenerateExpectedSchemaModel(tc *testcases.TestCase) (string, error) {
	if tc.GenerateExpectedSchemaModel() == "" {
		return "", nil
	}

	tmpl := `
		func {{ .Name }}_RawSchemaModel(t *testing.T) {

			want := &generator.SchemaModel{
				SourceSchema: {{ .Schema }},
			}

			tc := testcases.LoadAndParseTestCase("{{ .Name }}")

			got := mustLoadSchemaModel(t, tc.GenerateInput())

			diff.RequireKnownValueEqual(t, want.SourceSchema, got.SourceSchema)
		}
	`
	tmpld := template.Must(template.New("wantSchemaModel").Parse(tmpl))
	buf := bytes.NewBuffer(nil)
	err := tmpld.ExecuteTemplate(buf, "wantSchemaModel", map[string]any{
		"Name":   tc.Name(),
		"Schema": tc.GenerateExpectedSchemaModel(),
	})
	if err != nil {
		return "", err
	}
	return buf.String(), nil
}
