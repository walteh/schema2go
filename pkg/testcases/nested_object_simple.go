package testcases

import (
	"github.com/google/gnostic/jsonschema"
)

type nested_object_simple struct {
}

func init() {
	registerTestCase(&nested_object_simple{})
}

func (me *nested_object_simple) Name() string {
	return myfilename()
}

func (*nested_object_simple) JSONSchema() string {
	return /*jsonc*/ `{
		"$schema": "http://json-schema.org/draft-07/schema#",
		"title": "Parent",
		"type": "object",
		"required": ["child"],
		"properties": {
			"child": {
				"type": "object",
				"required": ["name"],
				"properties": {
					"name": {
						"type": "string"
					}
				}
			}
		}
	}`
}

func (*nested_object_simple) RawSchema() *jsonschema.Schema {
	return &jsonschema.Schema{
		Schema: ptr("http://json-schema.org/draft-07/schema#"),
		Title:  ptr("Parent"),
		Type:   typ("object"),
		Required: &[]string{
			"child",
		},
		Properties: &[]*jsonschema.NamedSchema{
			{
				Name: "child",
				Value: &jsonschema.Schema{
					Type: typ("object"),
					Required: &[]string{
						"name",
					},
					Properties: &[]*jsonschema.NamedSchema{
						{
							Name: "name",
							Value: &jsonschema.Schema{
								Type: typ("string"),
							},
						},
					},
				},
			},
		},
	}
}

func (*nested_object_simple) GoCode() string {
	return /*go*/ `
package example

`
}
