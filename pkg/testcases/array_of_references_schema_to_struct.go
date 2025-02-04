package testcases

import (
	"github.com/google/gnostic/jsonschema"
	"gopkg.in/yaml.v3"
)

type array_of_references_schema_to_struct struct {
}

func init() {
	registerTestCase(&array_of_references_schema_to_struct{})
}

func (me *array_of_references_schema_to_struct) Name() string {
	return myfilename()
}

func (*array_of_references_schema_to_struct) JSONSchema() string {
	return `{
		"$schema": "http://json-schema.org/draft-07/schema#",
		"title": "TeamConfig",
		"type": "object",
		"properties": {
			"members": {
				"type": "array",
				"items": { "$ref": "#/definitions/Member" },
				"description": "List of team members with their roles"
			},
			"projects": {
				"type": "array",
				"items": { "$ref": "#/definitions/Project" },
				"description": "List of team projects and their status"
			}
		},
		"definitions": {
			"Member": {
				"type": "object",
				"required": ["id"],
				"properties": {
					"id": {
						"type": "string",
						"description": "Unique member identifier"
					},
					"role": {
						"type": "string",
						"description": "Member's role in the team"
					}
				}
			},
			"Project": {
				"type": "object",
				"required": ["name"],
				"properties": {
					"name": {
						"type": "string",
						"description": "Project name"
					},
					"active": {
						"type": "boolean",
						"description": "Whether the project is active",
						"default": true
					}
				}
			}
		}
	}`
}

func (*array_of_references_schema_to_struct) RawSchema() *jsonschema.Schema {
	return &jsonschema.Schema{
		Schema: ptr("http://json-schema.org/draft-07/schema#"),
		Title:  ptr("TeamConfig"),
		Type:   typ("object"),
		Properties: &[]*jsonschema.NamedSchema{
			{
				Name: "members",
				Value: &jsonschema.Schema{
					Type:        typ("array"),
					Description: ptr("List of team members with their roles"),
					Items: &jsonschema.SchemaOrSchemaArray{
						Schema: &jsonschema.Schema{
							Ref: ptr("#/definitions/Member"),
						},
					},
				},
			},
			{
				Name: "projects",
				Value: &jsonschema.Schema{
					Type:        typ("array"),
					Description: ptr("List of team projects and their status"),
					Items: &jsonschema.SchemaOrSchemaArray{
						Schema: &jsonschema.Schema{
							Ref: ptr("#/definitions/Project"),
						},
					},
				},
			},
		},
		Definitions: &[]*jsonschema.NamedSchema{
			{
				Name: "Member",
				Value: &jsonschema.Schema{
					Type: typ("object"),
					Properties: &[]*jsonschema.NamedSchema{
						{
							Name: "id",
							Value: &jsonschema.Schema{
								Type:        typ("string"),
								Description: ptr("Unique member identifier"),
							},
						},
						{
							Name: "role",
							Value: &jsonschema.Schema{
								Type:        typ("string"),
								Description: ptr("Member's role in the team"),
							},
						},
					},
					Required: &[]string{"id"},
				},
			},
			{
				Name: "Project",
				Value: &jsonschema.Schema{
					Type: typ("object"),
					Properties: &[]*jsonschema.NamedSchema{
						{
							Name: "name",
							Value: &jsonschema.Schema{
								Type:        typ("string"),
								Description: ptr("Project name"),
							},
						},
						{
							Name: "active",
							Value: &jsonschema.Schema{
								Type:        typ("boolean"),
								Description: ptr("Whether the project is active"),
								Default:     &yaml.Node{Kind: yaml.ScalarNode, Value: "true", Tag: "!!bool"},
							},
						},
					},
					Required: &[]string{"name"},
				},
			},
		},
	}
}

func (*array_of_references_schema_to_struct) GoCode() string {
	return ``
}
