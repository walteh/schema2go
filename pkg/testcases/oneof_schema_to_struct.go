package testcases

import (
	"github.com/google/gnostic/jsonschema"
)

type oneof_schema_to_struct struct{}

func init() {
	registerTestCase(&oneof_schema_to_struct{})
}

func (t *oneof_schema_to_struct) Name() string {
	return myfilename()
}

func (t *oneof_schema_to_struct) JSONSchema() string {
	return /*jsonc*/ `{
		"$schema": "http://json-schema.org/draft-07/schema#",
		"title": "OneOfExample",
		"type": "object",
		"properties": {
			"identifier": {
				"oneOf": [{ "type": "string" }, { "type": "integer" }]
			},
			"status": {
				"oneOf": [
					{ "type": "boolean" },
					{ "type": "string", "enum": ["pending", "failed"] }
				]
			}
		}
	}`
}

func (t *oneof_schema_to_struct) RawSchema() *jsonschema.Schema {
	return &jsonschema.Schema{
		Schema: strPtr("http://json-schema.org/draft-07/schema#"),
		Title:  strPtr("OneOfExample"),
		Type:   typePtr("object"),
		Properties: &[]*jsonschema.NamedSchema{
			{
				Name: "identifier",
				Value: &jsonschema.Schema{
					OneOf: &[]*jsonschema.Schema{
						{Type: typePtr("string")},
						{Type: typePtr("integer")},
					},
				},
			},
			{
				Name: "status",
				Value: &jsonschema.Schema{
					OneOf: &[]*jsonschema.Schema{
						{Type: typePtr("boolean")},
						{
							Type: typePtr("string"),
							Enumeration: &[]jsonschema.SchemaEnumValue{
								{String: strPtr("pending")},
								{String: strPtr("failed")},
							},
						},
					},
				},
			},
		},
	}
}

func (t *oneof_schema_to_struct) GoCode() string {
	return /*go*/ `
package example

`
}
