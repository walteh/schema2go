package parser

import (
	"github.com/google/gnostic/jsonschema"
	"gopkg.in/yaml.v3"
)

func Ptr[T any](v T) *T {
	return &v
}

// Parse parses a JSON schema string into a gnostic Schema
func Parse(input string) (*jsonschema.Schema, error) {
	var node yaml.Node
	err := yaml.Unmarshal([]byte(input), &node)
	if err != nil {
		return nil, err
	}

	schema := jsonschema.NewSchemaFromObject(&node)

	// // Set the ID to "#" so that internal references work
	// schema.ID = Ptr("#")

	// Resolve all references and special types
	// schema.ResolveRefs()
	// schema.ResolveAllOfs()
	// schema.ResolveAnyOfs()

	return schema, nil
}

func NewType(s string) *jsonschema.Schema {
	return &jsonschema.Schema{
		Type: jsonschema.NewStringOrStringArrayWithString(s),
	}
}

func NewTypeArray(s []string) *jsonschema.Schema {
	return &jsonschema.Schema{
		Type: jsonschema.NewStringOrStringArrayWithStringArray(s),
	}
}

func NewDefinitionRef(name string) *jsonschema.Schema {
	return &jsonschema.Schema{
		Ref: Ptr("#/definitions/" + name),
	}
}

// Helper methods to make working with the schema easier

// GetRequiredFields returns the list of required fields or empty slice if none
func GetRequiredFields(schema *jsonschema.Schema) []string {
	if schema.Required == nil {
		return []string{}
	}
	return *schema.Required
}

// GetTitle returns the schema title or empty string if none
func GetTitle(schema *jsonschema.Schema) string {
	if schema.Title == nil {
		return ""
	}
	return *schema.Title
}

func GetTypeOrEmpty(schema *jsonschema.Schema) string {
	typ := GetType(schema)
	if typ == nil {
		return ""
	}
	return *typ
}

// GetType returns the schema type or empty string if none
func GetType(schema *jsonschema.Schema) *string {
	if schema.Type == nil {
		return nil
	}
	return schema.Type.String
}

func GetTypeArray(schema *jsonschema.Schema) *[]string {
	if schema.Type == nil {
		return nil
	}
	return schema.Type.StringArray
}

// GetDescription returns the schema description or empty string if none
func GetDescription(schema *jsonschema.Schema) string {
	if schema.Description == nil {
		return ""
	}
	return *schema.Description
}

// IsRequired checks if a field is required
func IsRequired(schema *jsonschema.Schema, fieldName string) bool {
	required := GetRequiredFields(schema)
	for _, r := range required {
		if r == fieldName {
			return true
		}
	}
	return false
}

// GetProperties returns a map of property name to schema
func GetProperties(schema *jsonschema.Schema) map[string]*jsonschema.Schema {
	result := make(map[string]*jsonschema.Schema)
	if schema.Properties == nil {
		return result
	}

	for _, prop := range *schema.Properties {
		if prop.Name != "" && prop.Value != nil {
			result[prop.Name] = prop.Value
		} else {
			panic("idk if this should happen, but we should handle it")
		}

	}
	return result
}

// GetDefinitions returns a map of definition name to schema
func GetDefinitions(schema *jsonschema.Schema) map[string]*jsonschema.Schema {
	result := make(map[string]*jsonschema.Schema)
	if schema.Definitions == nil {
		return result
	}

	for _, def := range *schema.Definitions {
		if def.Name != "" && def.Value != nil {
			result[def.Name] = def.Value
		} else {
			panic("idk if this should happen, but we should handle it")
		}
	}
	return result
}

// HasOneOf returns true if the schema has oneOf
func HasOneOf(schema *jsonschema.Schema) bool {
	return schema.OneOf != nil && len(*schema.OneOf) > 0
}

// HasAnyOf returns true if the schema has anyOf
func HasAnyOf(schema *jsonschema.Schema) bool {
	return schema.AnyOf != nil && len(*schema.AnyOf) > 0
}

// HasAllOf returns true if the schema has allOf
func HasAllOf(schema *jsonschema.Schema) bool {
	return schema.AllOf != nil && len(*schema.AllOf) > 0
}

// IsArray returns true if the schema is an array type
func IsArray(schema *jsonschema.Schema) bool {
	typ := GetType(schema)
	if typ == nil {
		return false
	}
	return *typ == "array"
}

// GetArrayItems returns the schema for array items or nil
func GetArrayItems(schema *jsonschema.Schema) *jsonschema.SchemaOrSchemaArray {
	if schema.Items == nil {
		return nil
	}

	return schema.Items

}
