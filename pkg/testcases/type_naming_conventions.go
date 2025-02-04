package testcases

import (
	"github.com/google/gnostic/jsonschema"
)

type type_naming_conventions struct{}

func init() {
	registerTestCase(&type_naming_conventions{})
}

func (t *type_naming_conventions) Name() string {
	return myfilename()
}

func (t *type_naming_conventions) JSONSchema() string {
	return /*jsonc*/ `{
		"$schema": "http://json-schema.org/draft-07/schema#",
		"title": "Config",
		"type": "object",
		"properties": {
			"inline": {
				"type": "object",
				"properties": {
					"value": {
						"oneOf": [{ "type": "string" }, { "type": "integer" }]
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
						{
							"type": "object",
							"properties": { "inline": { "type": "string" } }
						}
					]
				}
			}
		},
		"definitions": {
			"ReferencedType": {
				"type": "object",
				"properties": {
					"value": {
						"oneOf": [{ "type": "string" }, { "type": "integer" }]
					}
				}
			}
		}
	}`
}

func (t *type_naming_conventions) RawSchema() *jsonschema.Schema {
	return &jsonschema.Schema{
		Schema: strPtr("http://json-schema.org/draft-07/schema#"),
		Title:  strPtr("Config"),
		Type:   typePtr("object"),
		Properties: &[]*jsonschema.NamedSchema{
			{
				Name: "inline",
				Value: &jsonschema.Schema{
					Type: typePtr("object"),
					Properties: &[]*jsonschema.NamedSchema{
						{
							Name: "value",
							Value: &jsonschema.Schema{
								OneOf: &[]*jsonschema.Schema{
									{Type: typePtr("string")},
									{Type: typePtr("integer")},
								},
							},
						},
					},
				},
			},
			{
				Name: "referenced",
				Value: &jsonschema.Schema{
					Ref: strPtr("#/definitions/ReferencedType"),
				},
			},
			{
				Name: "mixedArray",
				Value: &jsonschema.Schema{
					Type: typePtr("array"),
					Items: &jsonschema.SchemaOrSchemaArray{
						Schema: &jsonschema.Schema{
							OneOf: &[]*jsonschema.Schema{
								{Ref: strPtr("#/definitions/ReferencedType")},
								{
									Type: typePtr("object"),
									Properties: &[]*jsonschema.NamedSchema{
										{
											Name: "inline",
											Value: &jsonschema.Schema{
												Type: typePtr("string"),
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
		Definitions: &[]*jsonschema.NamedSchema{
			{
				Name: "ReferencedType",
				Value: &jsonschema.Schema{
					Type: typePtr("object"),
					Properties: &[]*jsonschema.NamedSchema{
						{
							Name: "value",
							Value: &jsonschema.Schema{
								OneOf: &[]*jsonschema.Schema{
									{Type: typePtr("string")},
									{Type: typePtr("integer")},
								},
							},
						},
					},
				},
			},
		},
	}
}

func (t *type_naming_conventions) GoCode() string {
	return /*go*/ `
package example

`
}
