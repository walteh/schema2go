package testcases

import (
	"github.com/google/gnostic/jsonschema"
	"gopkg.in/yaml.v3"
)

type nested_object_with_optional struct {
}

func init() {
	registerTestCase(&nested_object_with_optional{})
}

func (me *nested_object_with_optional) Name() string {
	return myfilename()
}

func (*nested_object_with_optional) JSONSchema() string {
	return /*jsonc*/ `{
		"$schema": "http://json-schema.org/draft-07/schema#",
		"title": "Container",
		"type": "object",
		"properties": {
			"config": {
				"type": "object",
				"properties": {
					"enabled": {
						"type": "boolean",
						"default": true
					},
					"count": {
						"type": "integer"
					}
				}
			}
		}
	}`
}

func (*nested_object_with_optional) RawSchema() *jsonschema.Schema {
	return &jsonschema.Schema{
		Schema: ptr("http://json-schema.org/draft-07/schema#"),
		Title:  ptr("Container"),
		Type:   typ("object"),
		Properties: &[]*jsonschema.NamedSchema{
			{
				Name: "config",
				Value: &jsonschema.Schema{
					Type: typ("object"),
					Properties: &[]*jsonschema.NamedSchema{
						{
							Name: "enabled",
							Value: &jsonschema.Schema{
								Type: typ("boolean"),
								Default: &yaml.Node{
									Kind:  yaml.ScalarNode,
									Style: 0,
									Tag:   "!!bool",
									Value: "true",
								},
							},
						},
						{
							Name: "count",
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

func (*nested_object_with_optional) GoCode() string {
	return /*go*/ `
package example

`
}
