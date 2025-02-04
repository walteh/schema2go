package testcases

import (
	"github.com/google/gnostic/jsonschema"
	"gopkg.in/yaml.v3"
)

type string_enum_schema_to_struct struct {
}

func init() {
	registerTestCase(&string_enum_schema_to_struct{})
}

func (me *string_enum_schema_to_struct) Name() string {
	return myfilename()
}

func (*string_enum_schema_to_struct) JSONSchema() string {
	return /*jsonc*/ `{
		"$schema": "http://json-schema.org/draft-07/schema#",
		"title": "ColorConfig",
		"type": "object",
		"properties": {
			"primaryColor": {
				"type": "string",
				"enum": ["red", "green", "blue"]
			},
			"theme": {
				"type": "string",
				"enum": ["light", "dark"],
				"default": "light"
			}
		}
	}`
}

func (*string_enum_schema_to_struct) RawSchema() *jsonschema.Schema {
	return &jsonschema.Schema{
		Schema: ptr("http://json-schema.org/draft-07/schema#"),
		Title:  ptr("ColorConfig"),
		Type:   typ("object"),
		Properties: &[]*jsonschema.NamedSchema{
			{
				Name: "primaryColor",
				Value: &jsonschema.Schema{
					Type: typ("string"),
					Enumeration: &[]jsonschema.SchemaEnumValue{
						{String: ptr("red")},
						{String: ptr("green")},
						{String: ptr("blue")},
					},
				},
			},
			{
				Name: "theme",
				Value: &jsonschema.Schema{
					Type: typ("string"),
					Enumeration: &[]jsonschema.SchemaEnumValue{
						{String: ptr("light")},
						{String: ptr("dark")},
					},
					Default: &yaml.Node{
						Kind:  yaml.ScalarNode,
						Style: 0,
						Tag:   "!!str",
						Value: "light",
					},
				},
			},
		},
	}
}

func (*string_enum_schema_to_struct) GoCode() string {
	return /*go*/ `
package example

`
}
