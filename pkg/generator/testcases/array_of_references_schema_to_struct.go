package testcases

import (
	"github.com/google/gnostic/jsonschema"
	"github.com/walteh/schema2go/pkg/generator"
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

func (*array_of_references_schema_to_struct) StaticSchema() *generator.StaticSchema {
	memberFields := []generator.Field{
		&generator.StaticField{
			Name_:        "ID",
			JSONName_:    "id",
			Description_: "Unique member identifier",
			IsRequired_:  true,
			Type_:        "string",
			ValidationRules_: []generator.ValidationRule{
				{
					Type:    generator.ValidationRequired,
					Message: "id is required",
					Values:  "",
				},
			},
		},
		&generator.StaticField{
			Name_:        "Role",
			JSONName_:    "role",
			Description_: "Member's role in the team",
			IsRequired_:  false,
			Type_:        "*string",
		},
	}

	projectFields := []generator.Field{
		&generator.StaticField{
			Name_:        "Name",
			JSONName_:    "name",
			Description_: "Project name",
			IsRequired_:  true,
			Type_:        "string",
			ValidationRules_: []generator.ValidationRule{
				{
					Type:    generator.ValidationRequired,
					Message: "name is required",
					Values:  "",
				},
			},
		},
		&generator.StaticField{
			Name_:                "Active",
			JSONName_:            "active",
			Description_:         "Whether the project is active",
			IsRequired_:          false,
			Type_:                "*bool",
			DefaultValue_:        ptr("true"),
			DefaultValueComment_: ptr("true"),
		},
	}

	teamConfigFields := []generator.Field{
		&generator.StaticField{
			Name_:        "Members",
			JSONName_:    "members",
			Description_: "List of team members with their roles",
			IsRequired_:  false,
			Type_:        "[]Member",
		},
		&generator.StaticField{
			Name_:        "Projects",
			JSONName_:    "projects",
			Description_: "List of team projects and their status",
			IsRequired_:  false,
			Type_:        "[]Project",
		},
	}

	memberStruct := &generator.StaticStruct{
		Name_:                "Member",
		Fields_:              memberFields,
		HasDefaults_:         false,
		HasValidation_:       true,
		HasCustomMarshaling_: false,
	}

	projectStruct := &generator.StaticStruct{
		Name_:                "Project",
		Fields_:              projectFields,
		HasDefaults_:         true,
		HasValidation_:       true,
		HasCustomMarshaling_: false,
	}

	teamConfigStruct := &generator.StaticStruct{
		Name_:                "TeamConfig",
		Fields_:              teamConfigFields,
		HasDefaults_:         false,
		HasValidation_:       true,
		HasCustomMarshaling_: false,
	}

	return &generator.StaticSchema{
		Package_: "models",
		Structs_: []generator.Struct{memberStruct, projectStruct, teamConfigStruct},
		Enums_:   nil,
		Imports_: []string{
			"encoding/json",
			"gitlab.com/tozd/go/errors",
		},
	}
}

func (*array_of_references_schema_to_struct) GoCode() string {
	return `
	// Code generated by schema2go. DO NOT EDIT.
	// 🏗️ Generated from JSON Schema
	
	package models
	
	import (
		"encoding/json"
		"gitlab.com/tozd/go/errors"
	)
	
	// Member represents a team member
	type Member struct {
		// Unique member identifier
		ID string $$$json:"id"$$$ // Required
	
		// Member's role in the team
		Role *string $$$json:"role,omitempty"$$$
	}
	
	// Validate ensures all required fields are present
	func (m *Member) Validate() error {
		if m.ID == "" {
			return errors.New("id is required")
		}
		return nil
	}
	
	// Project represents a team project
	type Project struct {
		// Project name
		Name string $$$json:"name"$$$ // Required
	
		// Whether the project is active
		Active *bool $$$json:"active,omitempty"$$$ // Default: true
	}
	
	// Validate ensures all required fields are present
	func (p *Project) Validate() error {
		if p.Name == "" {
			return errors.New("name is required")
		}
		return nil
	}
	
	// TeamConfig represents team configuration with members and projects
	type TeamConfig struct {
		// List of team members with their roles
		Members []Member $$$json:"members,omitempty"$$$
	
		// List of team projects and their status
		Projects []Project $$$json:"projects,omitempty"$$$
	}
	
	// Validate ensures all nested objects are valid
	func (t *TeamConfig) Validate() error {
		for i, member := range t.Members {
			if err := member.Validate(); err != nil {
				return errors.Errorf("validating member %d: %w", i, err)
			}
		}
		for i, project := range t.Projects {
			if err := project.Validate(); err != nil {
				return errors.Errorf("validating project %d: %w", i, err)
			}
		}
		return nil
	}
	`
}
