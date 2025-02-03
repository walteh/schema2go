package testcases

import (
	"github.com/google/gnostic/jsonschema"
	"github.com/walteh/schema2go/pkg/generator"
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
	return `{
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

func (*required_fields_schema_to_struct) StaticSchema() *generator.StaticSchema {
	fields := []generator.Field{
		&generator.StaticField{
			Name_:        "ID",
			JSONName_:    "id",
			Description_: "",
			IsRequired_:  true,
			Type_:        "string",
			ValidationRules_: []generator.ValidationRule{
				{
					Type:    generator.ValidationRequired,
					Message: "id is required",
				},
			},
		},
		&generator.StaticField{
			Name_:        "Name",
			JSONName_:    "name",
			Description_: "",
			IsRequired_:  true,
			Type_:        "string",
			ValidationRules_: []generator.ValidationRule{
				{
					Type:    generator.ValidationRequired,
					Message: "name is required",
				},
			},
		},
		&generator.StaticField{
			Name_:        "Description",
			JSONName_:    "description",
			Description_: "",
			IsRequired_:  false,
			Type_:        "*string",
		},
	}

	requiredStruct := &generator.StaticStruct{
		Name_:                "RequiredExample",
		Fields_:              fields,
		HasDefaults_:         false,
		HasValidation_:       true,
		HasCustomMarshaling_: true,
	}

	return &generator.StaticSchema{
		Package_: "models",
		Structs_: []generator.Struct{
			requiredStruct,
		},
		Imports_: []string{
			"encoding/json",
			"gitlab.com/tozd/go/errors",
		},
	}
}

func (*required_fields_schema_to_struct) GoCode() string {
	return `
	// Code generated by schema2go. DO NOT EDIT.
	// 🏗️ Generated from JSON Schema
	
	package models
	
	import (
		"encoding/json"
		"gitlab.com/tozd/go/errors"
	)
	
	type RequiredExample struct {
		ID string $$$json:"id"$$$ // Required
	
		Name string $$$json:"name"$$$ // Required
	
		Description *string $$$json:"description,omitempty"$$$
	}
	
	// Validate ensures all required fields are present
	func (r *RequiredExample) Validate() error {
		if r.ID == "" {
			return errors.New("id is required")
		}
		if r.Name == "" {
			return errors.New("name is required")
		}
		return nil
	}
	`
}
