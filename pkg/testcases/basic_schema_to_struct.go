package testcases

import (
	"github.com/google/gnostic/jsonschema"
	"gopkg.in/yaml.v3"
)

type basic_schema_to_struct struct {
}

func init() {
	registerTestCase(&basic_schema_to_struct{})
}

func (me *basic_schema_to_struct) Name() string {
	return myfilename()
}

func (*basic_schema_to_struct) JSONSchema() string {
	return `{
		"$schema": "http://json-schema.org/draft-07/schema#",
		"title": "BasicExample",
		"type": "object",
		"required": ["id"],
		"properties": {
			"id": {
				"type": "string"
			},
			"count": {
				"type": "integer",
				"default": 0
			},
			"enabled": {
				"type": "boolean",
				"default": false
			},
			"ratio": {
				"type": "number"
			}
		}
	}`
}

func (*basic_schema_to_struct) RawSchema() *jsonschema.Schema {
	return &jsonschema.Schema{
		Schema: ptr("http://json-schema.org/draft-07/schema#"),
		Title:  ptr("BasicExample"),
		Type:   typ("object"),
		Properties: &[]*jsonschema.NamedSchema{
			{
				Name: "id",
				Value: &jsonschema.Schema{
					Type: typ("string"),
				},
			},
			{
				Name: "count",
				Value: &jsonschema.Schema{
					Type: typ("integer"),
					Default: &yaml.Node{
						Kind:  yaml.ScalarNode,
						Value: "0",
						Tag:   "!!int",
					},
				},
			},
			{
				Name: "enabled",
				Value: &jsonschema.Schema{
					Type: typ("boolean"),
					Default: &yaml.Node{
						Kind:  yaml.ScalarNode,
						Value: "false",
						Tag:   "!!bool",
					},
				},
			},
			{
				Name: "ratio",
				Value: &jsonschema.Schema{
					Type: typ("number"),
				},
			},
		},
		Required: &[]string{"id"},
	}
}

func (*basic_schema_to_struct) GoCode() string {
	return ``
}
