package generator_test

import (
	"context"
	_ "embed"
	"go/format"
	"strings"
	"testing"

	"github.com/google/gnostic/jsonschema"
	"github.com/stretchr/testify/assert"
	"github.com/walteh/schema2go/pkg/diff"
	"github.com/walteh/schema2go/pkg/generator"
)

func ptr[T any](v T) *T {
	return &v
}

func typ(v string) *jsonschema.StringOrStringArray {
	return jsonschema.NewStringOrStringArrayWithString(v)
}

/*
                     JSON Schema
                          |
                          v
      +----------------------------------+
      |        schema2go Generator       |
      |  +----------------------------+  |
      |  |    Parse JSON Schema      |  |
      |  +----------------------------+  |
      |              |                  |
      |              v                  |
      |  +----------------------------+ |
      |  |   Generate Go Structs     | |
      |  +----------------------------+ |
      +----------------------------------+
                      |
                      v
                   Go Code

Implementation Patterns:
----------------------
1. Pointer Fields:
   - Use pointers (*T) for optional fields (omitempty)
   - Use non-pointers for required fields
   - Use pointers for fields with default values
   Example:
   type Example struct {
       Required   string   ` + "`json:\"required\"`" + `         // Non-pointer
       Optional   *string  ` + "`json:\"optional,omitempty\"`" + ` // Pointer
       WithDefault *bool   ` + "`json:\"default,omitempty\"`" + ` // Pointer with default
   }

2. Validation:
   - Required fields -> Check for zero values
   - Enum fields -> Switch statement validation
   - Nested objects -> Recursive validation
   - Arrays -> Validate each element
   Example:
   func (e *Example) Validate() error {
       if e.Required == "" {
           return errors.New("required field is empty")
       }
       if err := e.Nested.Validate(); err != nil {
           return errors.Errorf("validating nested: %w", err)
       }
       return nil
   }

3. Default Values:
   - Constructor functions (New[Type]) for types with defaults
   - Default values set in constructor, not in struct
   Example:
   func NewExample() *Example {
       defaultValue := true
       return &Example{
           WithDefault: &defaultValue,
       }
   }

4. Special Types:
   - OneOf/AnyOf -> Custom marshaling required
   - AllOf -> Embedded fields with custom marshaling
   - Pattern Properties -> Map types with custom marshaling
   Example:
   func (o *OneOfType) UnmarshalJSON(data []byte) error {
       // Custom unmarshaling logic
       return nil
   }

Naming Conventions:
------------------
1. Referenced Types (defined in "definitions"):
   ReferencedType                 // Base type name
   ReferencedType_Value          // Nested field
   ReferencedType_Value_OneOf    // Special type (oneOf/anyOf/allOf)

2. Inline Types (nested in properties):
   ParentType_Field              // Nested object
   ParentType_Field_Value        // Nested field in nested object
   ParentType_Field_Value_OneOf  // Special type in nested object

3. Fields with Special Types:
   normalField            // Regular field
   enumField_Enum        // Enum field
   stringValue_OneOf     // OneOf field
   numberValue_AnyOf     // AnyOf field
   personInfo_AllOf      // AllOf field
   dynamicField_Pattern  // Pattern field

4. Pattern:
   [Parent_Path_To_Field][_TypeModifier]
   - Parent_Path_To_Field: Follows object nesting
   - TypeModifier: _OneOf, _AnyOf, _AllOf, _Enum, _Pattern

Examples:
---------
type Config struct {
    // Regular field
    Name string

    // Referenced type (from definitions)
    Info *PersonInfo

    // Inline nested object
    Address *Config_Address

    // Field with special type
    Status Config_Status_OneOf
}

type Config_Status_OneOf struct {
    StringValue_OneOf  *string
    IntegerValue_OneOf *int
}

type PersonInfo struct {
    // Referenced types don't need parent prefix
    Name string
    Age  int
}

TODO(schema2go): Implementation phases:
1. ⏳ Basic type conversion (string, int, bool)
2. ⏳ OneOf/AnyOf support
3. ⏳ AllOf merging
4. ⏳ References and definitions
5. ⏳ Custom type mappings
*/

// runTestCase is a helper function to run a single test case
func checkGoCode(t *testing.T, input, expectedOutput string) {
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

	model, err := generator.NewSchemaModel(input)
	if err != nil {
		t.Fatalf("Failed to parse schema: %v", err)
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

func checkSchemaMock(t *testing.T, got generator.Schema, want generator.Schema) {
	t.Helper()

	type mockField struct {
		Description         string
		Type                string
		IsRequired          bool
		IsEnum              bool
		EnumTypeName        string
		EnumValues          []*generator.EnumValue
		DefaultValue        *string
		DefaultValueComment *string
		ValidationRules     []*generator.ValidationRule
	}

	type mockStruct struct {
		Description         string
		Fields              []mockField
		HasAllOf            bool
		HasCustomMarshaling bool
		HasDefaults         bool
		HasValidation       bool
	}

	type mockSchema struct {
		Enums   []*generator.EnumModel
		Imports []string
		Package string
		Structs []mockStruct
	}

	assert.Equal(t, got.Package(), want.Package())
	assert.Equal(t, len(got.Structs()), len(want.Structs()))
	assert.Equal(t, len(got.Enums()), len(want.Enums()))
	assert.Equal(t, len(got.Imports()), len(want.Imports()))
	for i, gotStruct := range got.Structs() {
		wantStruct := want.Structs()[i]
		assert.Equal(t, gotStruct.Description(), wantStruct.Description(), "Struct description mismatch")
		assert.Equal(t, len(gotStruct.Fields()), len(wantStruct.Fields()), "Struct field count mismatch")
		assert.Equal(t, gotStruct.HasAllOf(), wantStruct.HasAllOf(), "Struct HasAllOf mismatch")
		assert.Equal(t, gotStruct.HasCustomMarshaling(), wantStruct.HasCustomMarshaling(), "Struct HasCustomMarshaling mismatch")
		assert.Equal(t, gotStruct.HasDefaults(), wantStruct.HasDefaults(), "Struct HasDefaults mismatch")
		assert.Equal(t, gotStruct.HasValidation(), wantStruct.HasValidation(), "Struct HasValidation mismatch")
		for j, gotField := range gotStruct.Fields() {
			wantField := wantStruct.Fields()[j]
			assert.Equal(t, gotField.Description(), wantField.Description(), "Field description mismatch")
			assert.Equal(t, gotField.Type(), wantField.Type(), "Field type mismatch")
			assert.Equal(t, gotField.IsRequired(), wantField.IsRequired(), "Field IsRequired mismatch")
			assert.Equal(t, gotField.DefaultValue(), wantField.DefaultValue(), "Field DefaultValue mismatch")
			assert.Equal(t, gotField.ValidationRules(), wantField.ValidationRules(), "Field ValidationRules mismatch")
			assert.Equal(t, gotField.IsEnum(), wantField.IsEnum(), "Field IsEnum mismatch")
		}
	}
}

func checkRawSchema(t *testing.T, got, want *jsonschema.Schema) {

}

//go:genr
