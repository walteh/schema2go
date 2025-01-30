package generator

import (
	"context"
	"go/format"
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

// testCase represents a single schema to struct conversion test
type testCase struct {
	name           string
	input          string
	expectedOutput string
}

func normalizeWhitespace(s string) string {
	// Replace all newlines and multiple spaces with a single space
	s = strings.Join(strings.Fields(s), " ")
	// Remove spaces around braces and parentheses
	s = strings.ReplaceAll(s, "{ ", "{")
	s = strings.ReplaceAll(s, " }", "}")
	s = strings.ReplaceAll(s, "( ", "(")
	s = strings.ReplaceAll(s, " )", ")")
	return s
}

// runTestCase is a helper function to run a single test case
func runTestCase(t *testing.T, tc testCase) {
	t.Helper() // marks this as a helper function for better test output
	t.Run(tc.name, func(t *testing.T) {
		gen := New(Options{
			PackageName: "models",
		})

		output, err := gen.Generate(context.Background(), tc.input)
		if err != nil {
			t.Fatalf("Failed to generate code: %v", err)
		}

		// format the go code

		formattedGot, err := format.Source([]byte(output))
		if err != nil {
			t.Fatalf("Failed to format code: %v", err)
		}

		formattedWant, err := format.Source([]byte(tc.expectedOutput))
		if err != nil {
			t.Fatalf("Failed to format code: %v", err)
		}

		diff.RequireKnownValueEqual(t, string(formattedWant), string(formattedGot))
	})
}

func TestBasicSchemaToStruct(t *testing.T) {
	runTestCase(t, testCase{
		name: "basic_types",
		input: `{
			"$schema": "http://json-schema.org/draft-07/schema#",
			"title": "BasicExample",
			"type": "object",
			"required": ["id"],
			"properties": {
				"id": { 
					"type": "string"
				},
				"count": { 
					"type": "integer",
					"default": 0
				},
				"enabled": { 
					"type": "boolean",
					"default": false
				},
				"ratio": {
					"type": "number"
				}
			}
		}`,
		expectedOutput: `// Code generated by schema2go. DO NOT EDIT.
// 🏗️ Generated from JSON Schema

package models

import (
	"encoding/json"
	"gitlab.com/tozd/go/errors"
)

type BasicExample struct {
	ID string ` + "`json:\"id\"`" + ` // Required
	Count *int ` + "`json:\"count,omitempty\"`" + ` // Default: 0
	Enabled *bool ` + "`json:\"enabled,omitempty\"`" + ` // Default: false
	Ratio *float64 ` + "`json:\"ratio,omitempty\"`" + `
}

// NewBasicExample creates a new BasicExample with default values
func NewBasicExample() *BasicExample {
	count := 0
	enabled := false
	return &BasicExample{
		Count: &count,
		Enabled: &enabled,
	}
}

// Validate ensures all required fields are present and valid
func (x *BasicExample) Validate() error {
	if x.ID == "" {
		return errors.New("id is required")
	}
	return nil
}`,
	})
}

func TestSchemaDocumentation(t *testing.T) {
	runTestCase(t, testCase{
		name: "documentation",
		input: `{
			"$schema": "http://json-schema.org/draft-07/schema#",
			"title": "DocumentedExample",
			"description": "A thoroughly documented example schema",
			"type": "object",
			"properties": {
				"field1": {
					"type": "string",
					"description": "A well documented string field",
					"examples": ["example1", "example2"]
				},
				"nested": {
					"type": "object",
					"description": "A nested object with its own documentation",
					"properties": {
						"subField": {
							"type": "integer",
							"description": "A documented integer field",
							"minimum": 0,
							"maximum": 100
						}
					}
				}
			}
		}`,
		expectedOutput: `package models

// DocumentedExample represents a thoroughly documented example schema
type DocumentedExample struct {
	// A well documented string field
	// Examples: ["example1", "example2"]
	Field1 *string ` + "`json:\"field1,omitempty\"`" + `

	// A nested object with its own documentation
	Nested *DocumentedExample_Nested ` + "`json:\"nested,omitempty\"`" + `
}

// DocumentedExample_Nested represents the nested object in DocumentedExample
type DocumentedExample_Nested struct {
	// A documented integer field
	// Range: [0, 100]
	SubField *int ` + "`json:\"subField,omitempty\"`" + `
}`,
	})
}

func TestOneOfSchemaToStruct(t *testing.T) {
	runTestCase(t, testCase{
		name: "oneof_types",
		input: `{
			"$schema": "http://json-schema.org/draft-07/schema#",
			"title": "OneOfExample",
			"type": "object",
			"properties": {
				"identifier": {
					"oneOf": [
						{ "type": "string" },
						{ "type": "integer" }
					]
				},
				"status": {
					"oneOf": [
						{ "type": "boolean" },
						{ "type": "string", "enum": ["pending", "failed"] }
					]
				}
			}
		}`,
		expectedOutput: `package models

type OneOfExample struct {
	Identifier OneOfExample_Identifier_OneOf ` + "`json:\"identifier,omitempty\"`" + `
	Status OneOfExample_Status_OneOf ` + "`json:\"status,omitempty\"`" + `
}

type OneOfExample_Identifier_OneOf struct {
	StringValue *string ` + "`json:\"string_value,omitempty\"`" + `
	IntegerValue *int ` + "`json:\"integer_value,omitempty\"`" + `
}

type OneOfExample_Status_OneOf struct {
	BooleanValue *bool ` + "`json:\"boolean_value,omitempty\"`" + `
	EnumValue *OneOfExample_Status_OneOf_Enum ` + "`json:\"enum_value,omitempty\"`" + `
}

type OneOfExample_Status_OneOf_Enum string

const (
	OneOfExample_Status_OneOf_Enum_Pending OneOfExample_Status_OneOf_Enum = "pending"
	OneOfExample_Status_OneOf_Enum_Failed  OneOfExample_Status_OneOf_Enum = "failed"
)

func (o *OneOfExample_Identifier_OneOf) UnmarshalJSON(data []byte) error {
	return nil // TODO: Implement
}

func (o OneOfExample_Identifier_OneOf) MarshalJSON() ([]byte, error) {
	return nil, nil // TODO: Implement
}

func (o *OneOfExample_Status_OneOf) UnmarshalJSON(data []byte) error {
	return nil // TODO: Implement
}

func (o OneOfExample_Status_OneOf) MarshalJSON() ([]byte, error) {
	return nil, nil // TODO: Implement
}`,
	})
}

func TestAnyOfSchemaToStruct(t *testing.T) {
	runTestCase(t, testCase{
		name: "anyof_types",
		input: `{
			"$schema": "http://json-schema.org/draft-07/schema#",
			"title": "AnyOfExample",
			"type": "object",
			"properties": {
				"value": {
					"anyOf": [
						{ "type": "string" },
						{ "type": "number" },
						{ "type": "boolean" }
					]
				},
				"nested": {
					"type": "object",
					"properties": {
						"field": {
							"anyOf": [
								{ "$ref": "#/definitions/RefType" },
								{ "type": "string" }
							]
						}
					}
				}
			},
			"definitions": {
				"RefType": {
					"type": "object",
					"properties": {
						"name": { "type": "string" }
					}
				}
			}
		}`,
		expectedOutput: `package models

type RefType struct {
	Name *string ` + "`json:\"name,omitempty\"`" + `
}

type AnyOfExample struct {
	Value AnyOfExample_Value ` + "`json:\"value,omitempty\"`" + `
	Nested *AnyOfExample_Nested ` + "`json:\"nested,omitempty\"`" + `
}

type AnyOfExample_Value struct {
	StringValue_AnyOf  *string  ` + "`json:\"string_value,omitempty\"`" + `
	NumberValue_AnyOf  *float64 ` + "`json:\"number_value,omitempty\"`" + `
	BooleanValue_AnyOf *bool    ` + "`json:\"boolean_value,omitempty\"`" + `
}

type AnyOfExample_Nested struct {
	Field AnyOfExample_Nested_Field ` + "`json:\"field,omitempty\"`" + `
}

type AnyOfExample_Nested_Field struct {
	RefValue_AnyOf   *RefType ` + "`json:\"ref_value,omitempty\"`" + `
	StringValue_AnyOf *string  ` + "`json:\"string_value,omitempty\"`" + `
}

func (a *AnyOfExample_Value) UnmarshalJSON(data []byte) error {
	return nil // TODO: Implement
}

func (a AnyOfExample_Value) MarshalJSON() ([]byte, error) {
	return nil, nil // TODO: Implement
}

func (a *AnyOfExample_Nested_Field) UnmarshalJSON(data []byte) error {
	return nil // TODO: Implement
}

func (a AnyOfExample_Nested_Field) MarshalJSON() ([]byte, error) {
	return nil, nil // TODO: Implement
}`,
	})
}

func TestAllOfSchemaToStruct(t *testing.T) {
	runTestCase(t, testCase{
		name: "allof_basic",
		input: `{
			"$schema": "http://json-schema.org/draft-07/schema#",
			"title": "AllOfExample",
			"type": "object",
			"allOf": [
				{
					"type": "object",
					"properties": {
						"name": { 
							"type": "string"
						}
					}
				},
				{
					"type": "object",
					"properties": {
						"age": { 
							"type": "integer"
						}
					}
				}
			]
		}`,
		expectedOutput: `package models

type AllOfExample struct {
	Name_AllOf *string ` + "`json:\"name,omitempty\"`" + `
	Age_AllOf  *int    ` + "`json:\"age,omitempty\"`" + `
}

func (a *AllOfExample) UnmarshalJSON(data []byte) error {
	return nil // TODO: Implement custom unmarshaling for allOf fields
}

func (a AllOfExample) MarshalJSON() ([]byte, error) {
	return nil, nil // TODO: Implement custom marshaling for allOf fields
}`,
	})
}

func TestAllOfWithRefsSchemaToStruct(t *testing.T) {
	runTestCase(t, testCase{
		name: "allof_with_refs",
		input: `{
			"$schema": "http://json-schema.org/draft-07/schema#",
			"title": "AllOfWithRefsExample",
			"type": "object",
			"allOf": [
				{ "$ref": "#/definitions/PersonInfo" },
				{ "$ref": "#/definitions/EmployeeInfo" }
			],
			"definitions": {
				"PersonInfo": {
					"type": "object",
					"properties": {
						"name": { 
							"type": "string"
						},
						"age": { 
							"type": "integer"
						}
					}
				},
				"EmployeeInfo": {
					"type": "object",
					"properties": {
						"employeeId": { 
							"type": "string"
						},
						"department": { 
							"type": "string"
						}
					}
				}
			}
		}`,
		expectedOutput: `package models

type PersonInfo struct {
	Name *string ` + "`json:\"name,omitempty\"`" + `
	Age  *int    ` + "`json:\"age,omitempty\"`" + `
}

type EmployeeInfo struct {
	EmployeeID *string ` + "`json:\"employeeId,omitempty\"`" + `
	Department *string ` + "`json:\"department,omitempty\"`" + `
}

type AllOfWithRefsExample struct {
	PersonInfo_AllOf   ` + "`json:\",inline\"`" + `
	EmployeeInfo_AllOf ` + "`json:\",inline\"`" + `
}

func (a *AllOfWithRefsExample) UnmarshalJSON(data []byte) error {
	return nil // TODO: Implement custom unmarshaling for allOf fields
}

func (a AllOfWithRefsExample) MarshalJSON() ([]byte, error) {
	return nil, nil // TODO: Implement custom marshaling for allOf fields
}`,
	})
}

func TestRequiredFieldsSchemaToStruct(t *testing.T) {
	runTestCase(t, testCase{
		name: "required_fields",
		input: `{
			"$schema": "http://json-schema.org/draft-07/schema#",
			"title": "RequiredExample",
			"type": "object",
			"required": ["id", "name"],
			"properties": {
				"id": { 
					"type": "string"
				},
				"name": { 
					"type": "string"
				},
				"description": { 
					"type": "string"
				}
			}
		}`,
		expectedOutput: `package models

type RequiredExample struct {
	ID string ` + "`json:\"id\"`" + ` // Required

	Name string ` + "`json:\"name\"`" + ` // Required

	Description *string ` + "`json:\"description,omitempty\"`" + `
}

// Validate ensures all required fields are present
func (r *RequiredExample) Validate() error {
	if r.ID == "" {
		return errors.New("id is required")
	}
	if r.Name == "" {
		return errors.New("name is required")
	}
	return nil
}`,
	})
}

func TestNestedObjectSchemaToStruct(t *testing.T) {
	runTestCase(t, testCase{
		name: "nested_object",
		input: `{
			"$schema": "http://json-schema.org/draft-07/schema#",
			"title": "UserProfile",
			"type": "object",
			"required": ["id", "address"],
			"properties": {
				"id": { 
					"type": "string"
				},
				"address": {
					"type": "object",
					"required": ["street", "city"],
					"properties": {
						"street": { 
							"type": "string"
						},
						"city": { 
							"type": "string"
						},
						"state": { 
							"type": "string"
						},
						"coordinates": {
							"type": "object",
							"properties": {
								"latitude": { "type": "number" },
								"longitude": { "type": "number" }
							}
						}
					}
				}
			}
		}`,
		expectedOutput: `package models

type Coordinates struct {
	Latitude  *float64 ` + "`json:\"latitude,omitempty\"`" + `
	Longitude *float64 ` + "`json:\"longitude,omitempty\"`" + `
}

type Address struct {
	Street string ` + "`json:\"street\"`" + ` // Required
	City   string ` + "`json:\"city\"`" + `   // Required
	State  *string ` + "`json:\"state,omitempty\"`" + `

	Coordinates *Coordinates ` + "`json:\"coordinates,omitempty\"`" + `
}

// Validate ensures all required fields are present
func (a *Address) Validate() error {
	if a.Street == "" {
		return errors.New("street is required")
	}
	if a.City == "" {
		return errors.New("city is required")
	}
	return nil
}

type UserProfile struct {
	ID      string  ` + "`json:\"id\"`" + `      // Required
	Address Address ` + "`json:\"address\"`" + ` // Required
}

// Validate ensures all required fields are present
func (u *UserProfile) Validate() error {
	if u.ID == "" {
		return errors.New("id is required")
	}
	if err := u.Address.Validate(); err != nil {
		return errors.Errorf("validating address: %w", err)
	}
	return nil
}`,
	})
}

func TestStringEnumSchemaToStruct(t *testing.T) {
	runTestCase(t, testCase{
		name: "string_enum",
		input: `{
			"$schema": "http://json-schema.org/draft-07/schema#",
			"title": "ColorConfig",
			"type": "object",
			"properties": {
				"primaryColor": {
					"type": "string",
					"enum": ["red", "green", "blue"]
				},
				"theme": {
					"type": "string",
					"enum": ["light", "dark"],
					"default": "light"
				}
			}
		}`,
		expectedOutput: `package models

type ColorType string

const (
	ColorTypeRed   ColorType = "red"
	ColorTypeGreen ColorType = "green"
	ColorTypeBlue  ColorType = "blue"
)

type ThemeType string

const (
	ThemeTypeLight ThemeType = "light"
	ThemeTypeDark  ThemeType = "dark"
)

type ColorConfig struct {
	PrimaryColor *ColorType ` + "`json:\"primaryColor,omitempty\"`" + `
	Theme        *ThemeType ` + "`json:\"theme,omitempty\"`" + ` // Default: light
}

// Validate ensures enum values are valid
func (c *ColorConfig) Validate() error {
	if c.PrimaryColor != nil {
		switch *c.PrimaryColor {
		case ColorTypeRed, ColorTypeGreen, ColorTypeBlue:
		default:
			return errors.New("invalid primary color")
		}
	}
	if c.Theme != nil {
		switch *c.Theme {
		case ThemeTypeLight, ThemeTypeDark:
		default:
			return errors.New("invalid theme")
		}
	}
	return nil
}`,
	})
}

func TestIntegerEnumSchemaToStruct(t *testing.T) {
	runTestCase(t, testCase{
		name: "integer_enum",
		input: `{
			"$schema": "http://json-schema.org/draft-07/schema#",
			"title": "HttpConfig",
			"type": "object",
			"properties": {
				"port": {
					"type": "integer",
					"enum": [80, 443, 8080, 8443]
				},
				"status": {
					"type": "integer",
					"enum": [200, 404, 500]
				}
			}
		}`,
		expectedOutput: `package models

type PortType int

const (
	PortType80   PortType = 80
	PortType443  PortType = 443
	PortType8080 PortType = 8080
	PortType8443 PortType = 8443
)

type StatusType int

const (
	StatusType200 StatusType = 200
	StatusType404 StatusType = 404
	StatusType500 StatusType = 500
)

type HttpConfig struct {
	Port   *PortType   ` + "`json:\"port,omitempty\"`" + `
	Status *StatusType ` + "`json:\"status,omitempty\"`" + `
}

// Validate ensures enum values are valid
func (h *HttpConfig) Validate() error {
	if h.Port != nil {
		switch *h.Port {
		case PortType80, PortType443, PortType8080, PortType8443:
		default:
			return errors.New("invalid port")
		}
	}
	if h.Status != nil {
		switch *h.Status {
		case StatusType200, StatusType404, StatusType500:
		default:
			return errors.New("invalid status")
		}
	}
	return nil
}`,
	})
}

func TestArrayOfReferencesSchemaToStruct(t *testing.T) {
	runTestCase(t, testCase{
		name: "array_of_refs",
		input: `{
			"$schema": "http://json-schema.org/draft-07/schema#",
			"title": "TeamConfig",
			"type": "object",
			"properties": {
				"members": {
					"type": "array",
					"items": { "$ref": "#/definitions/Member" },
					"description": "List of team members with their roles"
				},
				"projects": {
					"type": "array",
					"items": { "$ref": "#/definitions/Project" },
					"description": "List of team projects and their status"
				}
			},
			"definitions": {
				"Member": {
					"type": "object",
					"required": ["id"],
					"properties": {
						"id": { 
							"type": "string",
							"description": "Unique member identifier"
						},
						"role": { 
							"type": "string",
							"description": "Member's role in the team"
						}
					}
				},
				"Project": {
					"type": "object",
					"required": ["name"],
					"properties": {
						"name": { 
							"type": "string",
							"description": "Project name"
						},
						"active": { 
							"type": "boolean",
							"description": "Whether the project is active",
							"default": true
						}
					}
				}
			}
		}`,
		expectedOutput: `package models

// Member represents a team member
type Member struct {
	// Unique member identifier
	ID string ` + "`json:\"id\"`" + ` // Required

	// Member's role in the team
	Role *string ` + "`json:\"role,omitempty\"`" + `
}

// Validate ensures all required fields are present
func (m *Member) Validate() error {
	if m.ID == "" {
		return errors.New("id is required")
	}
	return nil
}

// Project represents a team project
type Project struct {
	// Project name
	Name string ` + "`json:\"name\"`" + ` // Required

	// Whether the project is active
	Active *bool ` + "`json:\"active,omitempty\"`" + ` // Default: true
}

// Validate ensures all required fields are present
func (p *Project) Validate() error {
	if p.Name == "" {
		return errors.New("name is required")
	}
	return nil
}

// TeamConfig represents team configuration with members and projects
type TeamConfig struct {
	// List of team members with their roles
	Members []Member ` + "`json:\"members,omitempty\"`" + `

	// List of team projects and their status
	Projects []Project ` + "`json:\"projects,omitempty\"`" + `
}

// Validate ensures all nested objects are valid
func (t *TeamConfig) Validate() error {
	for i, member := range t.Members {
		if err := member.Validate(); err != nil {
			return errors.Errorf("validating member %d: %w", i, err)
		}
	}
	for i, project := range t.Projects {
		if err := project.Validate(); err != nil {
			return errors.Errorf("validating project %d: %w", i, err)
		}
	}
	return nil
}`,
	})
}

func TestPatternPropertiesSchemaToStruct(t *testing.T) {
	runTestCase(t, testCase{
		name: "pattern_properties",
		input: `{
			"$schema": "http://json-schema.org/draft-07/schema#",
			"title": "DynamicConfig",
			"type": "object",
			"description": "Configuration with dynamic field names following specific patterns",
			"patternProperties": {
				"^S_": { 
					"type": "string",
					"description": "String fields must start with S_ prefix (e.g., S_name, S_description)"
				},
				"^N_": { 
					"type": "number",
					"description": "Number fields must start with N_ prefix (e.g., N_count, N_ratio)"
				}
			},
			"additionalProperties": false
		}`,
		expectedOutput: `package models

// DynamicConfig represents a configuration with dynamic field names following specific patterns
type DynamicConfig struct {
	// String fields must start with S_ prefix (e.g., S_name, S_description)
	StringFields_Pattern map[string]string ` + "`json:\"-\"`" + `

	// Number fields must start with N_ prefix (e.g., N_count, N_ratio)
	NumberFields_Pattern map[string]float64 ` + "`json:\"-\"`" + `
}

func (d *DynamicConfig) UnmarshalJSON(data []byte) error {
	// TODO: Implement pattern-based unmarshaling
	// 1. Parse raw JSON into map[string]interface{}
	// 2. For each key-value pair:
	//    - If key starts with S_, validate value is string and add to StringFields
	//    - If key starts with N_, validate value is number and add to NumberFields
	//    - Otherwise, return error for invalid pattern
	return nil
}

func (d DynamicConfig) MarshalJSON() ([]byte, error) {
	// TODO: Implement pattern-based marshaling
	// 1. Create output map combining both field maps
	// 2. Validate all keys follow required patterns
	// 3. Marshal to JSON
	return nil, nil
}

func (d *DynamicConfig) Validate() error {
	// Validate string field patterns
	for key := range d.StringFields_Pattern {
		if !strings.HasPrefix(key, "S_") {
			return errors.Errorf("invalid string field key %q: must start with S_", key)
		}
	}

	// Validate number field patterns
	for key := range d.NumberFields_Pattern {
		if !strings.HasPrefix(key, "N_") {
			return errors.Errorf("invalid number field key %q: must start with N_", key)
		}
	}
	return nil
}`,
	})
}

func TestTypeNamingConventions(t *testing.T) {
	runTestCase(t, testCase{
		name: "type_naming",
		input: `{
			"$schema": "http://json-schema.org/draft-07/schema#",
			"title": "Config",
			"type": "object",
			"properties": {
				"inline": {
					"type": "object",
					"properties": {
						"value": {
							"oneOf": [
								{ "type": "string" },
								{ "type": "integer" }
							]
						}
					}
				},
				"referenced": {
					"$ref": "#/definitions/ReferencedType"
				},
				"mixedArray": {
					"type": "array",
					"items": {
						"oneOf": [
							{ "$ref": "#/definitions/ReferencedType" },
							{ "type": "object", "properties": { "inline": { "type": "string" } } }
						]
					}
				}
			},
			"definitions": {
				"ReferencedType": {
					"type": "object",
					"properties": {
						"value": {
							"oneOf": [
								{ "type": "string" },
								{ "type": "integer" }
							]
						}
					}
				}
			}
		}`,
		expectedOutput: `package models

type Config struct {
	Inline *Config_Inline ` + "`json:\"inline,omitempty\"`" + `
	Referenced *ReferencedType ` + "`json:\"referenced,omitempty\"`" + `
	MixedArray []Config_MixedArray ` + "`json:\"mixedArray,omitempty\"`" + `
}

type Config_Inline struct {
	Value Config_Inline_Value ` + "`json:\"value,omitempty\"`" + `
}

type Config_Inline_Value struct {
	StringValue_OneOf *string ` + "`json:\"string_value,omitempty\"`" + `
	IntegerValue_OneOf *int ` + "`json:\"integer_value,omitempty\"`" + `
}

type ReferencedType struct {
	Value ReferencedType_Value ` + "`json:\"value,omitempty\"`" + `
}

type ReferencedType_Value struct {
	StringValue_OneOf *string ` + "`json:\"string_value,omitempty\"`" + `
	IntegerValue_OneOf *int ` + "`json:\"integer_value,omitempty\"`" + `
}`,
	})
}
