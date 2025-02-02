package testcases

import (
	"path/filepath"
	"runtime"
	"strings"

	"github.com/google/gnostic/jsonschema"
	"github.com/walteh/schema2go/pkg/generator"
)

func myfilename() string {
	_, file, _, _ := runtime.Caller(1)
	return strings.TrimSuffix(filepath.Base(file), ".go")
}

func ptr[T any](v T) *T {
	return &v
}

func typ(v string) *jsonschema.StringOrStringArray {
	return jsonschema.NewStringOrStringArrayWithString(v)
}

var registeredTestCases = map[string]TestCase{}

func registerTestCase(tc TestCase) {
	registeredTestCases[tc.Name()] = tc
}

func LoadTestCases() ([]TestCase, error) {
	cases := []TestCase{}
	for _, tc := range registeredTestCases {
		cases = append(cases, tc)
	}

	return cases, nil
}

// TestCase represents a single schema to struct conversion test
type TestCase interface {
	JSONSchema() string
	GoCode() string
	StaticSchema() *generator.StaticSchema
	RawSchema() *jsonschema.Schema
	Name() string
}
