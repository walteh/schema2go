package testcases

import (
	"github.com/google/gnostic/jsonschema"
	"github.com/walteh/schema2go/internal/archives/generator"
)

type type_naming_conventions struct{}

func init() {
	registerTestCase(&type_naming_conventions{})
}

func (t *type_naming_conventions) Name() string {
	return myfilename()
}

func (t *type_naming_conventions) JSONSchema() string {
	return `{
		"$schema": "http://json-schema.org/draft-07/schema#",
		"title": "Config",
		"type": "object",
		"properties": {
			"inline": {
				"type": "object",
				"properties": {
					"value": {
						"oneOf": [{ "type": "string" }, { "type": "integer" }]
					}
				}
			},
			"referenced": {
				"$ref": "#/definitions/ReferencedType"
			},
			"mixedArray": {
				"type": "array",
				"items": {
					"oneOf": [
						{ "$ref": "#/definitions/ReferencedType" },
						{
							"type": "object",
							"properties": { "inline": { "type": "string" } }
						}
					]
				}
			}
		},
		"definitions": {
			"ReferencedType": {
				"type": "object",
				"properties": {
					"value": {
						"oneOf": [{ "type": "string" }, { "type": "integer" }]
					}
				}
			}
		}
	}`
}

func (t *type_naming_conventions) RawSchema() *jsonschema.Schema {
	return &jsonschema.Schema{
		Schema: strPtr("http://json-schema.org/draft-07/schema#"),
		Title:  strPtr("Config"),
		Type:   typePtr("object"),
		Properties: &[]*jsonschema.NamedSchema{
			{
				Name: "inline",
				Value: &jsonschema.Schema{
					Type: typePtr("object"),
					Properties: &[]*jsonschema.NamedSchema{
						{
							Name: "value",
							Value: &jsonschema.Schema{
								OneOf: &[]*jsonschema.Schema{
									{Type: typePtr("string")},
									{Type: typePtr("integer")},
								},
							},
						},
					},
				},
			},
			{
				Name: "referenced",
				Value: &jsonschema.Schema{
					Ref: strPtr("#/definitions/ReferencedType"),
				},
			},
			{
				Name: "mixedArray",
				Value: &jsonschema.Schema{
					Type: typePtr("array"),
					Items: &jsonschema.SchemaOrSchemaArray{
						Schema: &jsonschema.Schema{
							OneOf: &[]*jsonschema.Schema{
								{Ref: strPtr("#/definitions/ReferencedType")},
								{
									Type: typePtr("object"),
									Properties: &[]*jsonschema.NamedSchema{
										{
											Name: "inline",
											Value: &jsonschema.Schema{
												Type: typePtr("string"),
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
		Definitions: &[]*jsonschema.NamedSchema{
			{
				Name: "ReferencedType",
				Value: &jsonschema.Schema{
					Type: typePtr("object"),
					Properties: &[]*jsonschema.NamedSchema{
						{
							Name: "value",
							Value: &jsonschema.Schema{
								OneOf: &[]*jsonschema.Schema{
									{Type: typePtr("string")},
									{Type: typePtr("integer")},
								},
							},
						},
					},
				},
			},
		},
	}
}

func (t *type_naming_conventions) StaticSchema() *generator.StaticSchema {
	return &generator.StaticSchema{
		Package_: "models",
		Structs_: []generator.Struct{
			&generator.StaticStruct{
				Name_:        "Config",
				Description_: "",
				Fields_: []generator.Field{
					&generator.StaticField{
						Name_:                "Inline",
						JSONName_:            "inline",
						Description_:         "",
						IsRequired_:          false,
						Type_:                "*Config_Inline",
						IsEnum_:              false,
						EnumTypeName_:        "",
						EnumValues_:          nil,
						DefaultValue_:        nil,
						DefaultValueComment_: nil,
						ValidationRules_:     nil,
					},
					&generator.StaticField{
						Name_:                "Referenced",
						JSONName_:            "referenced",
						Description_:         "",
						IsRequired_:          false,
						Type_:                "*ReferencedType",
						IsEnum_:              false,
						EnumTypeName_:        "",
						EnumValues_:          nil,
						DefaultValue_:        nil,
						DefaultValueComment_: nil,
						ValidationRules_:     nil,
					},
					&generator.StaticField{
						Name_:                "MixedArray",
						JSONName_:            "mixedArray",
						Description_:         "",
						IsRequired_:          false,
						Type_:                "[]Config_MixedArray",
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
				Name_:        "Config_Inline",
				Description_: "",
				Fields_: []generator.Field{
					&generator.StaticField{
						Name_:                "Value",
						JSONName_:            "value",
						Description_:         "",
						IsRequired_:          false,
						Type_:                "Config_Inline_Value",
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
				Name_:        "Config_Inline_Value",
				Description_: "",
				Fields_: []generator.Field{
					&generator.StaticField{
						Name_:                "StringValue_OneOf",
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
						Name_:                "IntegerValue_OneOf",
						JSONName_:            "integer_value",
						Description_:         "",
						IsRequired_:          false,
						Type_:                "*int",
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
				Name_:        "ReferencedType",
				Description_: "",
				Fields_: []generator.Field{
					&generator.StaticField{
						Name_:                "Value",
						JSONName_:            "value",
						Description_:         "",
						IsRequired_:          false,
						Type_:                "ReferencedType_Value",
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
				Name_:        "ReferencedType_Value",
				Description_: "",
				Fields_: []generator.Field{
					&generator.StaticField{
						Name_:                "StringValue_OneOf",
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
						Name_:                "IntegerValue_OneOf",
						JSONName_:            "integer_value",
						Description_:         "",
						IsRequired_:          false,
						Type_:                "*int",
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

func (t *type_naming_conventions) GoCode() string {
	return `// Code generated by schema2go. DO NOT EDIT.
// 🏗️ Generated from JSON Schema

package models

import (
	"encoding/json"
	"gitlab.com/tozd/go/errors"
)

type Config struct {
	Inline *Config_Inline $$$json:"inline,omitempty"$$$
	Referenced *ReferencedType $$$json:"referenced,omitempty"$$$
	MixedArray []Config_MixedArray $$$json:"mixedArray,omitempty"$$$
}

type Config_Inline struct {
	Value Config_Inline_Value $$$json:"value,omitempty"$$$
}

type Config_Inline_Value struct {
	StringValue_OneOf *string $$$json:"string_value,omitempty"$$$
	IntegerValue_OneOf *int $$$json:"integer_value,omitempty"$$$
}

type ReferencedType struct {
	Value ReferencedType_Value $$$json:"value,omitempty"$$$
}

type ReferencedType_Value struct {
	StringValue_OneOf *string $$$json:"string_value,omitempty"$$$
	IntegerValue_OneOf *int $$$json:"integer_value,omitempty"$$$
}`
}
