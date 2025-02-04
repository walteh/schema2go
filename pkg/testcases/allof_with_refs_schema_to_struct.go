package testcases

import (
	"github.com/google/gnostic/jsonschema"
)

type allof_with_refs_schema_to_struct struct{}

func init() {
	registerTestCase(&allof_with_refs_schema_to_struct{})
}

func (t *allof_with_refs_schema_to_struct) Name() string {
	return myfilename()
}

func (t *allof_with_refs_schema_to_struct) JSONSchema() string {
	return `{
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
	}`
}

func (t *allof_with_refs_schema_to_struct) RawSchema() *jsonschema.Schema {
	return &jsonschema.Schema{
		Schema: strPtr("http://json-schema.org/draft-07/schema#"),
		Title:  strPtr("AllOfWithRefsExample"),
		Type:   typePtr("object"),
		AllOf: &[]*jsonschema.Schema{
			{Ref: strPtr("#/definitions/PersonInfo")},
			{Ref: strPtr("#/definitions/EmployeeInfo")},
		},
		Definitions: &[]*jsonschema.NamedSchema{
			{
				Name: "PersonInfo",
				Value: &jsonschema.Schema{
					Type: typePtr("object"),
					Properties: &[]*jsonschema.NamedSchema{
						{
							Name: "name",
							Value: &jsonschema.Schema{
								Type: typePtr("string"),
							},
						},
						{
							Name: "age",
							Value: &jsonschema.Schema{
								Type: typePtr("integer"),
							},
						},
					},
				},
			},
			{
				Name: "EmployeeInfo",
				Value: &jsonschema.Schema{
					Type: typePtr("object"),
					Properties: &[]*jsonschema.NamedSchema{
						{
							Name: "employeeId",
							Value: &jsonschema.Schema{
								Type: typePtr("string"),
							},
						},
						{
							Name: "department",
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

func (t *allof_with_refs_schema_to_struct) GoCode() string {
	return ``
}
