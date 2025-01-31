package main

import (
	"bytes"
	"fmt"
	"go/format"
	"os"
	"path/filepath"
	"text/template"

	"github.com/walteh/schema2go/pkg/generator/testcases"
)

func generateTestCases() {
	tcs, err := testcases.LoadTestCases()
	if err != nil {
		panic(err)
	}
	funcNames := []string{}

	for _, tc := range tcs {
		buf := bytes.NewBuffer(nil)

		if tc.JSONSchema() == "" {
			funcNames = append(funcNames, `
			func `+tc.Name()+`(t *testing.T) {
				t.Skip("no json-schema")
			}
			`)
			continue
		}

		buf.WriteString("func " + tc.Name() + "(t *testing.T) {\n")
		rawSchemaModelFunc, err := GenerateRawSchemaTest(&tc)
		if err != nil {
			panic(err)
		}

		if rawSchemaModelFunc != "" {
			buf.WriteString(rawSchemaModelFunc)
			buf.WriteString("\n")
		}

		staticSchemaFunc, err := GenerateStaticSchemaTest(&tc)
		if err != nil {
			panic(err)
		}
		if staticSchemaFunc != "" {
			buf.WriteString(staticSchemaFunc)
			buf.WriteString("\n")
		}

		goCodeFunc, err := GenerateGoCodeTest("checkGoCode", &tc)
		if err != nil {
			panic(err)
		}
		if goCodeFunc != "" {
			buf.WriteString(goCodeFunc)
			buf.WriteString("\n")
		}

		buf.WriteString("}\n\n")

		funcNames = append(funcNames, buf.String())

	}

	buf := bytes.NewBuffer(nil)
	buf.WriteString("package generator_test\n\n")
	buf.WriteString("import (\n")
	buf.WriteString("\"testing\"\n")
	buf.WriteString("\"github.com/walteh/schema2go/pkg/generator/testcases\"\n")
	buf.WriteString("\"github.com/walteh/schema2go/pkg/generator\"\n")
	// buf.WriteString("\"github.com/walteh/schema2go/gen/mockery\"\n")
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
		buf.WriteString(rawSchemaModelFunc)
		buf.WriteString("\n\n")
	}

	data := buf.Bytes()

	// format the code
	formatted, err := format.Source(data)
	if err != nil {
		panic(fmt.Sprintf("Failed to format code: %v ... writing raw code: %s", err, string(data)))
	}

	if err := os.WriteFile(filepath.Join("..", "gen_test.go"), formatted, 0644); err != nil {
		panic(fmt.Sprintf("Failed to write file: %v", err))
	}
}

func GenerateGoCodeTest(funcName string, tc *testcases.TestCase) (string, error) {
	if tc.GoCode() == "" {
		return `
		t.Run("go-code", func(t *testing.T) {
			t.Fatalf("no go code")
		})
		`, nil
	}

	tmpl := `
		t.Run("go-code", func(t *testing.T) {
			tc := testcases.LoadAndParseTestCase("{{ .Name }}")

			got := mustLoadSchemaModel(t, tc.JSONSchema())

			{{.FuncName}}(t, got, tc.GoCode())
		})
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

func GenerateStaticSchemaTest(tc *testcases.TestCase) (string, error) {

	if tc.StaticSchema() == "" {
		return `
		t.Run("static-schema", func(t *testing.T) {
			t.Fatalf("no static schema")
		})
		`, nil
	}

	tmpl := `
		t.Run("static-schema", func(t *testing.T) {

			staticWant := {{ .StaticSchema }}

			tc := testcases.LoadAndParseTestCase("{{ .Name }}")

			got := mustLoadSchemaModel(t, tc.JSONSchema())

			staticGot := generator.NewStaticSchema(got)

			diff.RequireKnownValueEqual(t, staticWant, staticGot)
		})
	`
	tmpld := template.Must(template.New("wantSchemaModel").Parse(tmpl))
	buf := bytes.NewBuffer(nil)
	err := tmpld.ExecuteTemplate(buf, "wantSchemaModel", map[string]any{
		"Name":         tc.Name(),
		"StaticSchema": tc.StaticSchema(),
	})
	if err != nil {
		return "", err
	}
	return buf.String(), nil

}

func GenerateRawSchemaTest(tc *testcases.TestCase) (string, error) {
	if tc.RawSchema() == "" {
		return `
		t.Run("raw-schema", func(t *testing.T) {
			t.Fatalf("no raw schema")
		})
		`, nil
	}

	tmpl := `
		t.Run("raw-schema", func(t *testing.T) {

			want := &generator.SchemaModel{
				SourceSchema: {{ .RawSchema }},
			}

			tc := testcases.LoadAndParseTestCase("{{ .Name }}")

			got := mustLoadSchemaModel(t, tc.JSONSchema())

			diff.RequireKnownValueEqual(t, want.SourceSchema, got.SourceSchema)
		})
	`
	tmpld := template.Must(template.New("wantSchemaModel").Parse(tmpl))
	buf := bytes.NewBuffer(nil)
	err := tmpld.ExecuteTemplate(buf, "wantSchemaModel", map[string]any{
		"Name":      tc.Name(),
		"RawSchema": tc.RawSchema(),
	})
	if err != nil {
		return "", err
	}
	return buf.String(), nil
}
