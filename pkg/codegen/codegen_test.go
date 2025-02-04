package codegen_test

import (
	"context"
	"go/format"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/walteh/schema2go/internal/archives/generator"
	"github.com/walteh/schema2go/pkg/diff"
	"github.com/walteh/schema2go/pkg/testcases"
)

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

			t.Run("go-code", func(t *testing.T) {
				code := tc.GoCode()
				if code == "" {
					t.Skip("no go code to check")
				}
				replaced := strings.ReplaceAll(code, "$$$", "`")
				checkGoCode(t, schema, replaced)
			})
		})
	}

}

func checkGoCode(t *testing.T, model *generator.SchemaModel, expectedOutput string) {
	t.Helper() // marks this as a helper function for better test output

	ctx := context.Background()

	gen := generator.New(generator.Options{
		PackageName: "models",
	})

	// Format expected output
	formattedWant, err := format.Source([]byte(expectedOutput))
	if err != nil {
		t.Fatalf("Failed to format expected code: %v", err)
	}

	got, err := generator.GenerateWithFormatting(ctx, gen, model)
	if err != nil {
		if strings.Contains(err.Error(), "formatting code") {
			t.Logf("Formatting failed, trying again without formatting to show prettier output, this test will fail")
			// try again without formatting
			got, err = gen.Generate(ctx, model)
			if err != nil {
				t.Fatalf("Failed to generate code (without formatting): %v", err)
			}
			diff.RequireKnownValueEqual(t, normalizeCode(string(formattedWant)), normalizeCode(got))
			t.FailNow() // we always want to fail if formatting fails
		}
		t.Fatalf("Failed to generate code: %v", err)
	}

	diff.RequireKnownValueEqual(t, normalizeCode(string(formattedWant)), normalizeCode(got))
}

// normalizeCode removes comments and normalizes whitespace
func normalizeCode(code string) string {
	// Split into lines
	lines := strings.Split(code, "\n")

	// Process each line
	var result []string

	for _, line := range lines {
		// Skip empty lines and comment-only lines
		if line == "" || strings.HasPrefix(line, "//") {
			continue
		}

		// Remove inline comments
		if idx := strings.Index(line, "//"); idx >= 0 {
			line = strings.TrimSpace(line[:idx])
		}

		// Skip if line is now empty
		if line == "" {
			continue
		}

		result = append(result, line)
	}

	// Join lines back together
	return strings.Join(result, "\n")
}
