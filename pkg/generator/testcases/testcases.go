package testcases

import (
	"runtime"
	"strings"

	"github.com/google/gnostic/jsonschema"
	"github.com/walteh/schema2go/pkg/generator"
)

func myfilename() string {
	_, file, _, _ := runtime.Caller(1)
	return file
}

func ptr[T any](v T) *T {
	return &v
}

func typ(v string) *jsonschema.StringOrStringArray {
	return jsonschema.NewStringOrStringArrayWithString(v)
}

var registeredTestCases = map[string]TestCase{}

func registerTestCase(tc TestCase) {
	tc.goCode = strings.ReplaceAll(tc.goCode, "$$$", "`")
	registeredTestCases[tc.name] = tc
}

func LoadTestCases() ([]TestCase, error) {
	cases := []TestCase{}
	for _, tc := range registeredTestCases {
		cases = append(cases, tc)
	}

	return cases, nil
}

// TestCase represents a single schema to struct conversion test
type TestCase struct {
	name         string
	jsonSchema   string
	goCode       string
	staticSchema *generator.StaticSchema
	rawSchema    *jsonschema.Schema
}

func (tc *TestCase) JSONSchema() string {
	return tc.jsonSchema
}

func (tc *TestCase) GoCode() string {
	return tc.goCode
}

func (tc *TestCase) StaticSchema() *generator.StaticSchema {
	return tc.staticSchema
}

func (tc *TestCase) RawSchema() *jsonschema.Schema {
	return tc.rawSchema
}

func (tc *TestCase) Name() string {
	return tc.name
}
