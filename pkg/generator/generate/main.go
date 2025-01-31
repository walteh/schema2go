package generate

import (
	"bytes"
	"go/format"
	"html/template"
	"os"
	"path/filepath"

	"github.com/walteh/schema2go/pkg/generator/testcases"
)

func main() {
	tcs, err := testcases.LoadTestCases()
	if err != nil {
		panic(err)
	}
	funcNames := []string{}

	for _, tc := range tcs {

		rawSchemaModelFunc, err := GenerateExpectedSchemaModel("checkModel", &tc)
		if err != nil {
			panic(err)
		}
		funcNames = append(funcNames, rawSchemaModelFunc)

		schemaModelMockFunc, err := GenerateSchemaModelMockFunc("checkModel", &tc)
		if err != nil {
			panic(err)
		}
		funcNames = append(funcNames, schemaModelMockFunc)

		goCodeFunc, err := GenerateGoCodeFunc("checkModel", &tc)
		if err != nil {
			panic(err)
		}
		funcNames = append(funcNames, goCodeFunc)

	}

	buf := bytes.NewBuffer(nil)
	buf.WriteString("package generator_test\n\n")
	buf.WriteString("import (\n")
	buf.WriteString("\"testing\"\n")
	buf.WriteString("\"github.com/walteh/schema2go/pkg/generator\"\n")
	buf.WriteString("\"github.com/walteh/schema2go/gen/mockery\"\n")
	buf.WriteString(")\n\n")

	for _, rawSchemaModelFunc := range funcNames {
		buf.WriteString(rawSchemaModelFunc)
		buf.WriteString("\n\n")
	}

	data := buf.Bytes()

	// format the code
	formatted, err := format.Source(data)
	if err != nil {
		panic(err)
	}

	if err := os.WriteFile(filepath.Join("..", "expectations_test.gen.go"), formatted, 0644); err != nil {
		panic(err)
	}
}

func GenerateGoCodeFunc(funcName string, tc *testcases.TestCase) (string, error) {
	tmpl := `
		func {{ .Name }}_GoCode(t *testing.T) {
			tc := LoadAndParseTestCase("{{ .Name }}")
			{{.funcName}}(t, tc.GenerateInput(), tc.GenerateExpectedGoCode())
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

func GenerateSchemaModelMockFunc(funcName string, tc *testcases.TestCase) (string, error) {

	tmpl := `
		func {{ .Name }}_SchemaModelMock(t *testing.T) {

			{{ .SchemaModelMockGoCode }}

			tc := LoadAndParseTestCase("{{ .Name }}")

			got := mustLoadSchemaModel(t, tc.GenerateInput())

			{{.FuncName}}(t, schema, got)
		}
	`
	tmpld := template.Must(template.New("wantSchemaModel").Parse(tmpl))
	buf := bytes.NewBuffer(nil)
	err := tmpld.ExecuteTemplate(buf, "wantSchemaModel", map[string]any{
		"Name":                  tc.Name(),
		"SchemaModelMockGoCode": tc.GenerateExpectedSchemaModelMockGoCode(),
		"FuncName":              funcName,
	})
	if err != nil {
		return "", err
	}
	return buf.String(), nil

}

func GenerateExpectedSchemaModel(funcName string, tc *testcases.TestCase) (string, error) {

	tmpl := `
		func {{ .Name }}_RawSchemaModel(t *testing.T) {

			want := {{ .Schema }}

			got := mustLoadSchemaModel(t, tc.GenerateInput())

			{{.FuncName}}(t, want, got)
		}
	`
	tmpld := template.Must(template.New("wantSchemaModel").Parse(tmpl))
	buf := bytes.NewBuffer(nil)
	err := tmpld.ExecuteTemplate(buf, "wantSchemaModel", map[string]any{
		"Name":     tc.Name(),
		"Schema":   tc.GenerateExpectedSchemaModel(),
		"FuncName": funcName,
	})
	if err != nil {
		return "", err
	}
	return buf.String(), nil
}
