package testcases

import (
	"github.com/google/gnostic/jsonschema"
)

type pattern_properties_schema_to_struct struct{}

func (t *pattern_properties_schema_to_struct) Name() string {
	return myfilename()
}

func (t *pattern_properties_schema_to_struct) JSONSchema() string {
	return /*jsonc*/ `{
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
	}`
}

func (t *pattern_properties_schema_to_struct) RawSchema() *jsonschema.Schema {
	return &jsonschema.Schema{
		Schema:      strPtr("http://json-schema.org/draft-07/schema#"),
		Title:       strPtr("DynamicConfig"),
		Type:        typePtr("object"),
		Description: strPtr("Configuration with dynamic field names following specific patterns"),
		PatternProperties: &[]*jsonschema.NamedSchema{
			{
				Name: "^S_",
				Value: &jsonschema.Schema{
					Type:        typePtr("string"),
					Description: strPtr("String fields must start with S_ prefix (e.g., S_name, S_description)"),
				},
			},
			{
				Name: "^N_",
				Value: &jsonschema.Schema{
					Type:        typePtr("number"),
					Description: strPtr("Number fields must start with N_ prefix (e.g., N_count, N_ratio)"),
				},
			},
		},
		AdditionalProperties: &jsonschema.SchemaOrBoolean{
			Schema:  nil,
			Boolean: boolPtr(false),
		},
	}
}

func boolPtr(b bool) *bool {
	return &b
}

func (t *pattern_properties_schema_to_struct) GoCode() string {
	return /*go*/ `
package example

`
}
