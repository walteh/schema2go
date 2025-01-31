package generator

import (
	"context"
	"embed"
	_ "embed"
	"go/format"
	"path/filepath"
	"strings"
	"testing"

	"github.com/walteh/schema2go/pkg/diff"
)

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

// PROGRESS(passing): Basic type conversion with required fields and defaults
func TestBasicSchemaToStruct(t *testing.T) {
	runTestCase(t)
}

// PROGRESS(passing): Simple nested object with required fields
func TestNestedObjectSimple(t *testing.T) {
	runTestCase(t)
}

// PROGRESS(passing): String enum with default values
func TestStringEnumSchemaToStruct(t *testing.T) {
	runTestCase(t)
}

// PROGRESS(passing): Integer enum with validation
func TestIntegerEnumSchemaToStruct(t *testing.T) {
	runTestCase(t)
}

// PROGRESS(passing): Basic allOf merging
func TestAllOfSchemaToStruct(t *testing.T) {
	runTestCase(t)
}

// PROGRESS(untested): AllOf with referenced types
func TestAllOfWithRefsSchemaToStruct(t *testing.T) {
	t.Skip("Implementation in progress")
	runTestCase(t)
}

// PROGRESS(untested): OneOf type support
func TestOneOfSchemaToStruct(t *testing.T) {
	t.Skip("Implementation in progress")
	runTestCase(t)
}

// PROGRESS(untested): AnyOf type support
func TestAnyOfSchemaToStruct(t *testing.T) {
	t.Skip("Implementation in progress")
	runTestCase(t)
}

// PROGRESS(untested): Required fields validation
func TestRequiredFieldsSchemaToStruct(t *testing.T) {
	t.Skip("Implementation in progress")
	runTestCase(t)
}

// PROGRESS(untested): Optional nested objects with defaults
func TestNestedObjectWithOptional(t *testing.T) {
	t.Skip("Implementation in progress")
	runTestCase(t)
}

// PROGRESS(untested): Deep nested objects with validation
func TestNestedObjectDeep(t *testing.T) {
	t.Skip("Implementation in progress")
	runTestCase(t)
}

// PROGRESS(untested): Documentation and comments
func TestSchemaDocumentation(t *testing.T) {
	t.Skip("Implementation in progress")
	runTestCase(t)
}

// PROGRESS(untested): Array of referenced types
func TestArrayOfReferencesSchemaToStruct(t *testing.T) {
	t.Skip("Implementation in progress")
	runTestCase(t)
}

// PROGRESS(untested): Pattern properties with dynamic fields
func TestPatternPropertiesSchemaToStruct(t *testing.T) {
	t.Skip("Implementation in progress")
	runTestCase(t)
}

// PROGRESS(untested): Type naming conventions and references
func TestTypeNamingConventions(t *testing.T) {
	t.Skip("Implementation in progress")
	runTestCase(t)
}

// testCase represents a single schema to struct conversion test
type testCase struct {
	input          string
	expectedOutput string
}

// runTestCase is a helper function to run a single test case
func runTestCase(t *testing.T) {
	t.Helper() // marks this as a helper function for better test output

	tc := loadAndParseTestCase(t)

	gen := New(Options{
		PackageName: "models",
	})

	formattedWant, err := format.Source([]byte(tc.expectedOutput))
	if err != nil {
		t.Fatalf("Failed to format expected code: %v", err)
	}

	got, err := GenerateWithFormatting(context.Background(), gen, tc.input)
	if err != nil {
		if strings.Contains(err.Error(), "formatting code") {
			t.Logf("Formatting failed, trying again without formatting to show prettier output, this test will fail")
			// try again without formatting
			got, err = gen.Generate(context.Background(), tc.input)
			if err != nil {
				t.Fatalf("Failed to generate code (without formatting): %v", err)
			}
			diff.RequireKnownValueEqual(t, string(formattedWant), string(got))
			t.FailNow() // we always want to fail if formatting fails
		}
		t.Fatalf("Failed to generate code: %v", err)
	}

	diff.RequireKnownValueEqual(t, string(formattedWant), string(got))
}

//go:embed testcases/Test*.md
var embedTestCases embed.FS

func loadAndParseTestCase(t *testing.T) testCase {
	t.Helper()
	content, err := embedTestCases.ReadFile(filepath.Join("testcases", t.Name()+".md"))
	if err != nil {
		t.Fatalf("Failed to read test case file: %v", err)
	}
	return parseTestCase(t, string(content))
}

func parseTestCase(t *testing.T, text string) testCase {
	t.Helper()
	// Split the markdown into sections
	// We expect:
	// ```json
	// <input schema>
	// ```
	//
	// ```go
	// <expected output>
	// ```

	// Find the JSON section
	jsonStart := strings.Index(text, "```json\n")
	jsonEnd := strings.Index(text[jsonStart+7:], "\n```")
	if jsonStart == -1 || jsonEnd == -1 {
		panic("Could not find JSON section in test case markdown")
	}
	input := strings.TrimSpace(text[jsonStart+7 : jsonStart+7+jsonEnd])

	// Find the Go section
	goStart := strings.Index(text, "```go\n")
	goEnd := strings.Index(text[goStart+5:], "\n```")
	if goStart == -1 || goEnd == -1 {
		panic("Could not find Go section in test case markdown")
	}
	expectedOutput := strings.TrimSpace(text[goStart+5 : goStart+5+goEnd])

	return testCase{
		input:          input,
		expectedOutput: expectedOutput,
	}
}
