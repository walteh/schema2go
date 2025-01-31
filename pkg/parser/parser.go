package parser

import (
	"github.com/google/gnostic/jsonschema"
	"github.com/k0kubun/pp/v3"
	"gopkg.in/yaml.v3"
)

// Store YAML nodes for each schema
var schemaNodes = make(map[*jsonschema.Schema]*yaml.Node)

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

	pp.Printf("🔍 Parsing YAML node: %+v\n", node)
	schema := jsonschema.NewSchemaFromObject(&node)
	pp.Printf("📝 Created schema: %+v\n", schema)

	// Store the root node
	schemaNodes[schema] = &node

	// Store nodes for properties
	if schema.Properties != nil {
		for _, prop := range *schema.Properties {
			if prop.Value != nil {
				for i := 0; i < len(node.Content[0].Content); i += 2 {
					if i+1 >= len(node.Content[0].Content) {
						break
					}
					if node.Content[0].Content[i].Value == "properties" {
						propsNode := node.Content[0].Content[i+1]
						for j := 0; j < len(propsNode.Content); j += 2 {
							if j+1 >= len(propsNode.Content) {
								break
							}
							if propsNode.Content[j].Value == prop.Name {
								schemaNodes[prop.Value] = propsNode.Content[j+1]
								break
							}
						}
						break
					}
				}
			}
		}
	}

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

// GetEnumFromNode returns the enum values from a YAML node or nil if none
func GetEnumFromNode(node *yaml.Node) []jsonschema.SchemaEnumValue {
	if node == nil {
		return nil
	}

	// Find the enum field
	var enumNode *yaml.Node
	for i := 0; i < len(node.Content); i += 2 {
		if i+1 >= len(node.Content) {
			break
		}
		if node.Content[i].Value == "enum" {
			enumNode = node.Content[i+1]
			break
		}
	}

	if enumNode == nil || enumNode.Kind != yaml.SequenceNode {
		return nil
	}

	var result []jsonschema.SchemaEnumValue
	for _, val := range enumNode.Content {
		switch val.Tag {
		case "!!str":
			result = append(result, jsonschema.SchemaEnumValue{String: &val.Value})
		case "!!bool":
			boolVal := val.Value == "true"
			result = append(result, jsonschema.SchemaEnumValue{Bool: &boolVal})
		case "!!int":
			result = append(result, jsonschema.SchemaEnumValue{String: &val.Value})
		}
	}

	return result
}

// GetEnum returns the enum values from a schema or nil if none
func GetEnum(schema *jsonschema.Schema) []jsonschema.SchemaEnumValue {
	// Try to get enum values directly from the YAML node
	if node := schemaNodes[schema]; node != nil {
		if enums := GetEnumFromNode(node); len(enums) > 0 {
			pp.Printf("🔢 Getting enum values from schema: %+v\n", schema)
			pp.Printf("📊 Enum values: %+v\n", enums)
			return enums
		}
	}

	// Fallback to gnostic's enum values
	if schema.Enumeration == nil {
		return nil
	}
	pp.Printf("🔢 Getting enum values from schema: %+v\n", schema)
	pp.Printf("📊 Enum values: %+v\n", *schema.Enumeration)
	return *schema.Enumeration
}
