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

		buf.WriteString("func " + tc.Name() + "(t *testing.T) {\n\n")
		if tc.JSONSchema() == "" {
			buf.WriteString("	t.Fatalf(\"no json-schema definition defined for testcases/" + tc.Name() + ".md\")\n")
		}
		buf.WriteString("	tc := testcases.LoadTestCase(\"" + tc.Name() + "\")\n")
		buf.WriteString("	require.Equal(t, testCasesHash, testcases.GetHash(), \"test cases hash mismatch, please run 'go generate ./...' to update the test cases hash\")\n\n")
		buf.WriteString("	schema, err := generator.NewSchemaModel(tc.JSONSchema())\n")
		buf.WriteString("	require.NoError(t, err, \"failed to parse schema\")\n")

		rawSchemaModelFunc, err := GenerateRawSchemaTest(&tc)
		if err != nil {
			panic(err)
		}

		jsonSchemaFunc, err := GenerateJSONSchemaTest(tc.Name(), &tc)
		if err != nil {
			panic(err)
		}

		if jsonSchemaFunc != "" {
			buf.WriteString(jsonSchemaFunc)
			buf.WriteString("\n")
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
	buf.WriteString("\"github.com/stretchr/testify/require\"\n")
	// buf.WriteString("\"github.com/walteh/schema2go/gen/mockery\"\n")
	buf.WriteString("\"github.com/google/gnostic/jsonschema\"\n")
	buf.WriteString("\"github.com/walteh/schema2go/pkg/diff\"\n")
	buf.WriteString("\"gopkg.in/yaml.v3\"\n")
	buf.WriteString(")\n\n")

	// buf.WriteString("const testCasesHash = \"" + testcases.GetHash() + "\"\n\n")

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

func GenerateJSONSchemaTest(funcName string, tc *testcases.TestCase) (string, error) {

	tmpl := `
		t.Run("json-schema", func(t *testing.T) {
			// nothing to do here right now
		})
	`
	return tmpl, nil
}

func GenerateGoCodeTest(funcName string, tc *testcases.TestCase) (string, error) {
	if tc.GoCode() == "" {
		return `
		t.Run("go-code", func(t *testing.T) {
			t.Fatalf("no go-code test case defined for ` + tc.Name() + `.md")
		})
		`, nil
	}

	tmpl := `
		t.Run("go-code", func(t *testing.T) {

			{{.FuncName}}(t, schema, tc.GoCode())
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

	if tc.StaticSchema() == nil {
		return `
		t.Run("static-schema", func(t *testing.T) {
			t.Fatalf("no static-schema test case defined for testcases/` + tc.Name() + `.md")
		})
		`, nil
	}

	tmpl := `
		t.Run("static-schema", func(t *testing.T) {

			staticWant := tc.StaticSchema()

			staticGot := generator.NewStaticSchema(schema)

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
	if tc.RawSchema() == nil {
		return `
		t.Run("raw-schema", func(t *testing.T) {
			t.Fatalf("no raw-schema test case defined for testcases/` + tc.Name() + `.md")
		})
		`, nil
	}

	tmpl := `
		t.Run("raw-schema", func(t *testing.T) {

			want := &generator.SchemaModel{
				SourceSchema: tc.RawSchema(),
			}

			schema.RemoveYamlLineNumbers()

			diff.RequireKnownValueEqual(t, want.SourceSchema, schema.SourceSchema)
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
