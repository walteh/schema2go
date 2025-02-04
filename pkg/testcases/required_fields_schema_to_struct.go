package testcases

import (
	"github.com/google/gnostic/jsonschema"
)

type required_fields_schema_to_struct struct {
}

func init() {
	registerTestCase(&required_fields_schema_to_struct{})
}

func (me *required_fields_schema_to_struct) Name() string {
	return myfilename()
}

func (*required_fields_schema_to_struct) JSONSchema() string {
	return /*jsonc*/ `{
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
	}`
}

func (*required_fields_schema_to_struct) RawSchema() *jsonschema.Schema {
	return &jsonschema.Schema{
		Schema: ptr("http://json-schema.org/draft-07/schema#"),
		Title:  ptr("RequiredExample"),
		Type:   typ("object"),
		Required: &[]string{
			"id",
			"name",
		},
		Properties: &[]*jsonschema.NamedSchema{
			{
				Name: "id",
				Value: &jsonschema.Schema{
					Type: typ("string"),
				},
			},
			{
				Name: "name",
				Value: &jsonschema.Schema{
					Type: typ("string"),
				},
			},
			{
				Name: "description",
				Value: &jsonschema.Schema{
					Type: typ("string"),
				},
			},
		},
	}
}

func (*required_fields_schema_to_struct) GoCode() string {
	return /*go*/ `
package example

`
}
