package testcases

import (
	"github.com/google/gnostic/jsonschema"
)

type schema_documentation struct{}

func init() {
	registerTestCase(&schema_documentation{})
}

func (t *schema_documentation) Name() string {
	return myfilename()
}

func (t *schema_documentation) JSONSchema() string {
	return `{
		"$schema": "http://json-schema.org/draft-07/schema#",
		"title": "DocumentedExample",
		"description": "A thoroughly documented example schema",
		"type": "object",
		"properties": {
			"field1": {
				"type": "string",
				"description": "A well documented string field",
				"examples": ["example1", "example2"]
			},
			"nested": {
				"type": "object",
				"description": "A nested object with its own documentation",
				"properties": {
					"subField": {
						"type": "integer",
						"description": "A documented integer field",
						"minimum": 0,
						"maximum": 100
					}
				}
			}
		}
	}`
}

func (t *schema_documentation) RawSchema() *jsonschema.Schema {
	return &jsonschema.Schema{
		Schema:      strPtr("http://json-schema.org/draft-07/schema#"),
		Title:       strPtr("DocumentedExample"),
		Description: strPtr("A thoroughly documented example schema"),
		Type:        typePtr("object"),
		Properties: &[]*jsonschema.NamedSchema{
			{
				Name: "field1",
				Value: &jsonschema.Schema{
					Type:        typePtr("string"),
					Description: strPtr("A well documented string field"),
				},
			},
			{
				Name: "nested",
				Value: &jsonschema.Schema{
					Type:        typePtr("object"),
					Description: strPtr("A nested object with its own documentation"),
					Properties: &[]*jsonschema.NamedSchema{
						{
							Name: "subField",
							Value: &jsonschema.Schema{
								Type:        typePtr("integer"),
								Description: strPtr("A documented integer field"),
								Minimum:     jsonschema.NewSchemaNumberWithFloat(0),
								Maximum:     jsonschema.NewSchemaNumberWithFloat(100),
							},
						},
					},
				},
			},
		},
	}
}

func float64Ptr(v float64) *float64 {
	return &v
}

func (t *schema_documentation) GoCode() string {
	return ``
}
