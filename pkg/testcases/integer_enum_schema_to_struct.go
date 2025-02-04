package testcases

import (
	"github.com/google/gnostic/jsonschema"
)

type integer_enum_schema_to_struct struct{}

func init() {
	registerTestCase(&integer_enum_schema_to_struct{})
}

func (t *integer_enum_schema_to_struct) Name() string {
	return myfilename()
}

func (t *integer_enum_schema_to_struct) JSONSchema() string {
	return /*jsonc*/ `{
		"$schema": "http://json-schema.org/draft-07/schema#",
		"title": "HttpConfig",
		"type": "object",
		"properties": {
			"port": {
				"type": "integer",
				"enum": [80, 443, 8080, 8443]
			},
			"status": {
				"type": "integer",
				"enum": [200, 404, 500]
			}
		}
	}`
}

func (t *integer_enum_schema_to_struct) RawSchema() *jsonschema.Schema {
	return &jsonschema.Schema{
		Schema: strPtr("http://json-schema.org/draft-07/schema#"),
		Title:  strPtr("HttpConfig"),
		Type:   typePtr("object"),
		Properties: &[]*jsonschema.NamedSchema{
			{
				Name: "port",
				Value: &jsonschema.Schema{
					Type: typePtr("integer"),
					Enumeration: &[]jsonschema.SchemaEnumValue{
						{String: strPtr("80")},
						{String: strPtr("443")},
						{String: strPtr("8080")},
						{String: strPtr("8443")},
					},
				},
			},
			{
				Name: "status",
				Value: &jsonschema.Schema{
					Type: typePtr("integer"),
					Enumeration: &[]jsonschema.SchemaEnumValue{
						{String: strPtr("200")},
						{String: strPtr("404")},
						{String: strPtr("500")},
					},
				},
			},
		},
	}
}

func (t *integer_enum_schema_to_struct) GoCode() string {
	return /*go*/ `
package example

`
}
