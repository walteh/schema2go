package generator_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/walteh/schema2go/pkg/diff"
	"github.com/walteh/schema2go/pkg/generator"
	"github.com/walteh/schema2go/pkg/generator/testcases"
)

const testCasesHash = "1ec3aaecb246ccf10046aaaa2a0e858bfa39a614faed8acd97109782e0a147a2"

func TestAll(t *testing.T) {

	tc, err := testcases.LoadTestCases()
	require.NoError(t, err, "failed to load test cases")
	// should load at least one test case
	require.Greater(t, len(tc), 0, "no test cases loaded")

	for _, tc := range tc {
		t.Run(tc.Name(), func(t *testing.T) {
			schema, err := generator.NewSchemaModel(tc.JSONSchema())
			require.NoError(t, err, "failed to parse schema")

			t.Run("json-schema", func(t *testing.T) {
				// nothing to do here right now
			})

			t.Run("raw-schema", func(t *testing.T) {

				want := &generator.SchemaModel{
					SourceSchema: tc.RawSchema(),
				}

				schema.RemoveYamlLineNumbers()

				diff.RequireKnownValueEqual(t, want.SourceSchema, schema.SourceSchema)
			})

			t.Run("static-schema", func(t *testing.T) {

				staticWant := tc.StaticSchema()

				staticGot := generator.NewStaticSchema(schema)

				diff.RequireKnownValueEqual(t, staticWant, staticGot)
			})

			t.Run("go-code", func(t *testing.T) {
				checkGoCode(t, schema, tc.GoCode())
			})
		})
	}

}
