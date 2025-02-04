package testcases

import (
	"github.com/google/gnostic/jsonschema"
)

type anyof_schema_to_struct struct{}

func init() {
	registerTestCase(&anyof_schema_to_struct{})
}

func (t *anyof_schema_to_struct) Name() string {
	return myfilename()
}

func (t *anyof_schema_to_struct) JSONSchema() string {
	return `{
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
	}`
}

func (t *anyof_schema_to_struct) RawSchema() *jsonschema.Schema {
	return &jsonschema.Schema{
		Schema: strPtr("http://json-schema.org/draft-07/schema#"),
		Title:  strPtr("AnyOfExample"),
		Type:   typePtr("object"),
		Properties: &[]*jsonschema.NamedSchema{
			{
				Name: "value",
				Value: &jsonschema.Schema{
					AnyOf: &[]*jsonschema.Schema{
						{Type: typePtr("string")},
						{Type: typePtr("number")},
						{Type: typePtr("boolean")},
					},
				},
			},
			{
				Name: "nested",
				Value: &jsonschema.Schema{
					Type: typePtr("object"),
					Properties: &[]*jsonschema.NamedSchema{
						{
							Name: "field",
							Value: &jsonschema.Schema{
								AnyOf: &[]*jsonschema.Schema{
									{Ref: strPtr("#/definitions/RefType")},
									{Type: typePtr("string")},
								},
							},
						},
					},
				},
			},
		},
		Definitions: &[]*jsonschema.NamedSchema{
			{
				Name: "RefType",
				Value: &jsonschema.Schema{
					Type: typePtr("object"),
					Properties: &[]*jsonschema.NamedSchema{
						{
							Name: "name",
							Value: &jsonschema.Schema{
								Type: typePtr("string"),
							},
						},
					},
				},
			},
		},
	}
}

func strPtr(s string) *string {
	return &s
}

func typePtr(s string) *jsonschema.StringOrStringArray {
	return &jsonschema.StringOrStringArray{
		String: &s,
	}
}

func (t *anyof_schema_to_struct) GoCode() string {
	return ``
}
