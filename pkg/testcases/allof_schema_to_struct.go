package testcases

import (
	"github.com/google/gnostic/jsonschema"
)

type allof_schema_to_struct struct {
}

func init() {
	registerTestCase(&allof_schema_to_struct{})
}

func (me *allof_schema_to_struct) Name() string {
	return myfilename()
}

func (*allof_schema_to_struct) JSONSchema() string {
	return /*jsonc*/ `{
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
	}`
}

func (*allof_schema_to_struct) RawSchema() *jsonschema.Schema {
	return &jsonschema.Schema{
		Schema: ptr("http://json-schema.org/draft-07/schema#"),
		Title:  ptr("AllOfExample"),
		Type:   typ("object"),
		AllOf: &[]*jsonschema.Schema{
			{
				Type: typ("object"),
				Properties: &[]*jsonschema.NamedSchema{
					{Name: "name", Value: &jsonschema.Schema{Type: typ("string")}},
				},
			},
			{
				Type: typ("object"),
				Properties: &[]*jsonschema.NamedSchema{
					{Name: "age", Value: &jsonschema.Schema{Type: typ("integer")}},
				},
			},
		},
	}
}

func (*allof_schema_to_struct) GoCode() string {
	return /*go*/ `
package example

`
}
