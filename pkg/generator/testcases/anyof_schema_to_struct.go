package testcases

import (
	"github.com/google/gnostic/jsonschema"
	"github.com/walteh/schema2go/pkg/generator"
)

type anyof_schema_to_struct struct{}

func (t *anyof_schema_to_struct) Name() string {
	return "anyof_schema_to_struct"
}

func (t *anyof_schema_to_struct) JSONSchema() string {
	return `{
		"$schema": "http://json-schema.org/draft-07/schema#",
		"title": "AnyOfExample",
		"type": "object",
		"properties": {
			"value": {
				"anyOf": [
					{ "type": "string" },
					{ "type": "number" },
					{ "type": "boolean" }
				]
			},
			"nested": {
				"type": "object",
				"properties": {
					"field": {
						"anyOf": [
							{ "$ref": "#/definitions/RefType" },
							{ "type": "string" }
						]
					}
				}
			}
		},
		"definitions": {
			"RefType": {
				"type": "object",
				"properties": {
					"name": { "type": "string" }
				}
			}
		}
	}`
}

func (t *anyof_schema_to_struct) RawSchema() *jsonschema.Schema {
	return &jsonschema.Schema{
		Schema: strPtr("http://json-schema.org/draft-07/schema#"),
		Title:  strPtr("AnyOfExample"),
		Type:   typePtr("object"),
		Properties: &[]*jsonschema.NamedSchema{
			{
				Name: "value",
				Value: &jsonschema.Schema{
					AnyOf: &[]*jsonschema.Schema{
						{Type: typePtr("string")},
						{Type: typePtr("number")},
						{Type: typePtr("boolean")},
					},
				},
			},
			{
				Name: "nested",
				Value: &jsonschema.Schema{
					Type: typePtr("object"),
					Properties: &[]*jsonschema.NamedSchema{
						{
							Name: "field",
							Value: &jsonschema.Schema{
								AnyOf: &[]*jsonschema.Schema{
									{Ref: strPtr("#/definitions/RefType")},
									{Type: typePtr("string")},
								},
							},
						},
					},
				},
			},
		},
		Definitions: &[]*jsonschema.NamedSchema{
			{
				Name: "RefType",
				Value: &jsonschema.Schema{
					Type: typePtr("object"),
					Properties: &[]*jsonschema.NamedSchema{
						{
							Name: "name",
							Value: &jsonschema.Schema{
								Type: typePtr("string"),
							},
						},
					},
				},
			},
		},
	}
}

func strPtr(s string) *string {
	return &s
}

func typePtr(s string) *jsonschema.StringOrStringArray {
	return &jsonschema.StringOrStringArray{
		String: &s,
	}
}

func (t *anyof_schema_to_struct) StaticSchema() *generator.StaticSchema {
	return &generator.StaticSchema{
		Package_: "models",
		Structs_: []generator.Struct{
			&generator.StaticStruct{
				Name_:        "RefType",
				Description_: "",
				Fields_: []generator.Field{
					&generator.StaticField{
						Name_:                "Name",
						JSONName_:            "name",
						Description_:         "",
						IsRequired_:          false,
						Type_:                "*string",
						IsEnum_:              false,
						EnumTypeName_:        "",
						EnumValues_:          nil,
						DefaultValue_:        nil,
						DefaultValueComment_: nil,
						ValidationRules_:     nil,
					},
				},
				HasAllOf_:            false,
				HasCustomMarshaling_: false,
				HasDefaults_:         false,
				HasValidation_:       false,
			},
			&generator.StaticStruct{
				Name_:        "AnyOfExample",
				Description_: "",
				Fields_: []generator.Field{
					&generator.StaticField{
						Name_:                "Value",
						JSONName_:            "value",
						Description_:         "",
						IsRequired_:          false,
						Type_:                "AnyOfExample_Value",
						IsEnum_:              false,
						EnumTypeName_:        "",
						EnumValues_:          nil,
						DefaultValue_:        nil,
						DefaultValueComment_: nil,
						ValidationRules_:     nil,
					},
					&generator.StaticField{
						Name_:                "Nested",
						JSONName_:            "nested",
						Description_:         "",
						IsRequired_:          false,
						Type_:                "*AnyOfExample_Nested",
						IsEnum_:              false,
						EnumTypeName_:        "",
						EnumValues_:          nil,
						DefaultValue_:        nil,
						DefaultValueComment_: nil,
						ValidationRules_:     nil,
					},
				},
				HasAllOf_:            false,
				HasCustomMarshaling_: false,
				HasDefaults_:         false,
				HasValidation_:       false,
			},
			&generator.StaticStruct{
				Name_:        "AnyOfExample_Value",
				Description_: "",
				Fields_: []generator.Field{
					&generator.StaticField{
						Name_:                "StringValue_AnyOf",
						JSONName_:            "string_value",
						Description_:         "",
						IsRequired_:          false,
						Type_:                "*string",
						IsEnum_:              false,
						EnumTypeName_:        "",
						EnumValues_:          nil,
						DefaultValue_:        nil,
						DefaultValueComment_: nil,
						ValidationRules_:     nil,
					},
					&generator.StaticField{
						Name_:                "NumberValue_AnyOf",
						JSONName_:            "number_value",
						Description_:         "",
						IsRequired_:          false,
						Type_:                "*float64",
						IsEnum_:              false,
						EnumTypeName_:        "",
						EnumValues_:          nil,
						DefaultValue_:        nil,
						DefaultValueComment_: nil,
						ValidationRules_:     nil,
					},
					&generator.StaticField{
						Name_:                "BooleanValue_AnyOf",
						JSONName_:            "boolean_value",
						Description_:         "",
						IsRequired_:          false,
						Type_:                "*bool",
						IsEnum_:              false,
						EnumTypeName_:        "",
						EnumValues_:          nil,
						DefaultValue_:        nil,
						DefaultValueComment_: nil,
						ValidationRules_:     nil,
					},
				},
				HasAllOf_:            false,
				HasCustomMarshaling_: true,
				HasDefaults_:         false,
				HasValidation_:       false,
			},
			&generator.StaticStruct{
				Name_:        "AnyOfExample_Nested",
				Description_: "",
				Fields_: []generator.Field{
					&generator.StaticField{
						Name_:                "Field",
						JSONName_:            "field",
						Description_:         "",
						IsRequired_:          false,
						Type_:                "AnyOfExample_Nested_Field",
						IsEnum_:              false,
						EnumTypeName_:        "",
						EnumValues_:          nil,
						DefaultValue_:        nil,
						DefaultValueComment_: nil,
						ValidationRules_:     nil,
					},
				},
				HasAllOf_:            false,
				HasCustomMarshaling_: false,
				HasDefaults_:         false,
				HasValidation_:       false,
			},
			&generator.StaticStruct{
				Name_:        "AnyOfExample_Nested_Field",
				Description_: "",
				Fields_: []generator.Field{
					&generator.StaticField{
						Name_:                "RefValue_AnyOf",
						JSONName_:            "ref_value",
						Description_:         "",
						IsRequired_:          false,
						Type_:                "*RefType",
						IsEnum_:              false,
						EnumTypeName_:        "",
						EnumValues_:          nil,
						DefaultValue_:        nil,
						DefaultValueComment_: nil,
						ValidationRules_:     nil,
					},
					&generator.StaticField{
						Name_:                "StringValue_AnyOf",
						JSONName_:            "string_value",
						Description_:         "",
						IsRequired_:          false,
						Type_:                "*string",
						IsEnum_:              false,
						EnumTypeName_:        "",
						EnumValues_:          nil,
						DefaultValue_:        nil,
						DefaultValueComment_: nil,
						ValidationRules_:     nil,
					},
				},
				HasAllOf_:            false,
				HasCustomMarshaling_: true,
				HasDefaults_:         false,
				HasValidation_:       false,
			},
		},
	}
}

func (t *anyof_schema_to_struct) GoCode() string {
	return `// Code generated by schema2go. DO NOT EDIT.
// 🏗️ Generated from JSON Schema

package models

import (
	"encoding/json"
	"gitlab.com/tozd/go/errors"
)

type RefType struct {
	Name *string $$$json:\"name,omitempty\"$$$
}

type AnyOfExample struct {
	Value AnyOfExample_Value $$$json:\"value,omitempty\"$$$
	Nested *AnyOfExample_Nested $$$json:\"nested,omitempty\"$$$
}

type AnyOfExample_Value struct {
	StringValue_AnyOf  *string  $$$json:\"string_value,omitempty\"$$$
	NumberValue_AnyOf  *float64 $$$json:\"number_value,omitempty\"$$$
	BooleanValue_AnyOf *bool    $$$json:\"boolean_value,omitempty\"$$$
}

type AnyOfExample_Nested struct {
	Field AnyOfExample_Nested_Field $$$json:\"field,omitempty\"$$$
}

type AnyOfExample_Nested_Field struct {
	RefValue_AnyOf   *RefType $$$json:\"ref_value,omitempty\"$$$
	StringValue_AnyOf *string  $$$json:\"string_value,omitempty\"$$$
}

func (a *AnyOfExample_Value) UnmarshalJSON(data []byte) error {
	return nil // TODO: Implement
}

func (a AnyOfExample_Value) MarshalJSON() ([]byte, error) {
	return nil, nil // TODO: Implement
}

func (a *AnyOfExample_Nested_Field) UnmarshalJSON(data []byte) error {
	return nil // TODO: Implement
}

func (a AnyOfExample_Nested_Field) MarshalJSON() ([]byte, error) {
	return nil, nil // TODO: Implement
}`
}
