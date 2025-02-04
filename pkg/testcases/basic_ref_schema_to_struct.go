package testcases

import (
	"github.com/google/gnostic/jsonschema"
)

type basic_ref_schema_to_struct struct {
}

func init() {
	registerTestCase(&basic_ref_schema_to_struct{})
}

func (me *basic_ref_schema_to_struct) Name() string {
	return myfilename()
}

func (*basic_ref_schema_to_struct) JSONSchema() string {
	return /*jsonc*/ `{
		"$schema": "http://json-schema.org/draft-07/schema#",
		"title": "RefExample",
		"type": "object",
		"properties": {
			"person": {
				"$ref": "#/definitions/Person"
			},
			"innerPerson": {
				"type": "object",
				"properties": {
					"name": {
						"type": "string"
					},
					"age": {
						"type": "integer"
					}
				}
			}
		},
		"definitions": {
			"Person": {
				"type": "object",
				"properties": {
					"name": {
						"type": "string"
					},
					"age": {
						"type": "integer"
					}
				}
			}
		}
	}`
}

func (*basic_ref_schema_to_struct) RawSchema() *jsonschema.Schema {
	return &jsonschema.Schema{
		Schema: ptr("http://json-schema.org/draft-07/schema#"),
		Title:  ptr("RefExample"),
		Type:   typ("object"),
		Properties: &[]*jsonschema.NamedSchema{
			{
				Name: "person",
				Value: &jsonschema.Schema{
					Ref: ptr("#/definitions/Person"),
				},
			},
			{
				Name: "innerPerson",
				Value: &jsonschema.Schema{
					Type: typ("object"),
					Properties: &[]*jsonschema.NamedSchema{
						{
							Name: "name",
							Value: &jsonschema.Schema{
								Type: typ("string"),
							},
						},
						{
							Name: "age",
							Value: &jsonschema.Schema{
								Type: typ("integer"),
							},
						},
					},
				},
			},
		},
		Definitions: &[]*jsonschema.NamedSchema{
			{
				Name: "Person",
				Value: &jsonschema.Schema{
					Type: typ("object"),
					Properties: &[]*jsonschema.NamedSchema{
						{
							Name: "name",
							Value: &jsonschema.Schema{
								Type: typ("string"),
							},
						},
						{
							Name: "age",
							Value: &jsonschema.Schema{
								Type: typ("integer"),
							},
						},
					},
				},
			},
		},
	}
}

func (*basic_ref_schema_to_struct) GoCode() string {
	return /*go*/ `
package example

`
}
