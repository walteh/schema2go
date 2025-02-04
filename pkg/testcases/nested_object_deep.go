package testcases

import (
	"github.com/google/gnostic/jsonschema"
)

type nested_object_deep struct{}

func init() {
	registerTestCase(&nested_object_deep{})
}

func (t *nested_object_deep) Name() string {
	return myfilename()
}

func (t *nested_object_deep) JSONSchema() string {
	return /*jsonc*/ `{
		"$schema": "http://json-schema.org/draft-07/schema#",
		"title": "Location",
		"type": "object",
		"required": ["address"],
		"properties": {
			"address": {
				"type": "object",
				"required": ["coordinates"],
				"properties": {
					"coordinates": {
						"type": "object",
						"required": ["latitude", "longitude"],
						"properties": {
							"latitude": {
								"type": "number"
							},
							"longitude": {
								"type": "number"
							}
						}
					}
				}
			}
		}
	}`
}

func (t *nested_object_deep) RawSchema() *jsonschema.Schema {
	return &jsonschema.Schema{
		Schema: strPtr("http://json-schema.org/draft-07/schema#"),
		Title:  strPtr("Location"),
		Type:   typePtr("object"),
		Required: &[]string{
			"address",
		},
		Properties: &[]*jsonschema.NamedSchema{
			{
				Name: "address",
				Value: &jsonschema.Schema{
					Type: typePtr("object"),
					Required: &[]string{
						"coordinates",
					},
					Properties: &[]*jsonschema.NamedSchema{
						{
							Name: "coordinates",
							Value: &jsonschema.Schema{
								Type: typePtr("object"),
								Required: &[]string{
									"latitude",
									"longitude",
								},
								Properties: &[]*jsonschema.NamedSchema{
									{
										Name: "latitude",
										Value: &jsonschema.Schema{
											Type: typePtr("number"),
										},
									},
									{
										Name: "longitude",
										Value: &jsonschema.Schema{
											Type: typePtr("number"),
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

func (t *nested_object_deep) GoCode() string {
	return /*go*/ `
package example

`
}
