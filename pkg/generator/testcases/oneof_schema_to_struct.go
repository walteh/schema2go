package testcases

import (
	"github.com/google/gnostic/jsonschema"
	"github.com/walteh/schema2go/pkg/generator"
)

type oneof_schema_to_struct struct{}

func init() {
	registerTestCase(&oneof_schema_to_struct{})
}

func (t *oneof_schema_to_struct) Name() string {
	return myfilename()
}

func (t *oneof_schema_to_struct) JSONSchema() string {
	return `{
		"$schema": "http://json-schema.org/draft-07/schema#",
		"title": "OneOfExample",
		"type": "object",
		"properties": {
			"identifier": {
				"oneOf": [{ "type": "string" }, { "type": "integer" }]
			},
			"status": {
				"oneOf": [
					{ "type": "boolean" },
					{ "type": "string", "enum": ["pending", "failed"] }
				]
			}
		}
	}`
}

func (t *oneof_schema_to_struct) RawSchema() *jsonschema.Schema {
	return &jsonschema.Schema{
		Schema: strPtr("http://json-schema.org/draft-07/schema#"),
		Title:  strPtr("OneOfExample"),
		Type:   typePtr("object"),
		Properties: &[]*jsonschema.NamedSchema{
			{
				Name: "identifier",
				Value: &jsonschema.Schema{
					OneOf: &[]*jsonschema.Schema{
						{Type: typePtr("string")},
						{Type: typePtr("integer")},
					},
				},
			},
			{
				Name: "status",
				Value: &jsonschema.Schema{
					OneOf: &[]*jsonschema.Schema{
						{Type: typePtr("boolean")},
						{
							Type: typePtr("string"),
							Enumeration: &[]jsonschema.SchemaEnumValue{
								{String: strPtr("pending")},
								{String: strPtr("failed")},
							},
						},
					},
				},
			},
		},
	}
}

func (t *oneof_schema_to_struct) StaticSchema() *generator.StaticSchema {
	return &generator.StaticSchema{
		Package_: "models",
		Structs_: []generator.Struct{
			&generator.StaticStruct{
				Name_:        "OneOfExample",
				Description_: "",
				Fields_: []generator.Field{
					&generator.StaticField{
						Name_:                "Identifier",
						JSONName_:            "identifier",
						Description_:         "",
						IsRequired_:          false,
						Type_:                "OneOfExample_Identifier_OneOf",
						IsEnum_:              false,
						EnumTypeName_:        "",
						EnumValues_:          nil,
						DefaultValue_:        nil,
						DefaultValueComment_: nil,
						ValidationRules_:     nil,
					},
					&generator.StaticField{
						Name_:                "Status",
						JSONName_:            "status",
						Description_:         "",
						IsRequired_:          false,
						Type_:                "OneOfExample_Status_OneOf",
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
				Name_:        "OneOfExample_Identifier_OneOf",
				Description_: "",
				Fields_: []generator.Field{
					&generator.StaticField{
						Name_:                "StringValue",
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
						Name_:                "IntegerValue",
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
				Name_:        "OneOfExample_Status_OneOf",
				Description_: "",
				Fields_: []generator.Field{
					&generator.StaticField{
						Name_:                "BooleanValue",
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
					&generator.StaticField{
						Name_:                "EnumValue",
						JSONName_:            "enum_value",
						Description_:         "",
						IsRequired_:          false,
						Type_:                "*OneOfExample_Status_OneOf_Enum",
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
				Name_:        "OneOfExample_Status_OneOf_Enum",
				Description_: "",
				Fields_: []generator.Field{
					&generator.StaticField{
						Name_:         "Value",
						JSONName_:     "value",
						Description_:  "",
						IsRequired_:   false,
						Type_:         "string",
						IsEnum_:       true,
						EnumTypeName_: "OneOfExample_Status_OneOf_Enum",
						EnumValues_: []generator.EnumValue{
							{
								Name:  "OneOfExample_Status_OneOf_Enum_Pending",
								Value: "pending",
							},
							{
								Name:  "OneOfExample_Status_OneOf_Enum_Failed",
								Value: "failed",
							},
						},
					},
				},
				HasAllOf_:            false,
				HasCustomMarshaling_: false,
				HasDefaults_:         false,
				HasValidation_:       false,
			},
		},
	}
}

func (t *oneof_schema_to_struct) GoCode() string {
	return `// Code generated by schema2go. DO NOT EDIT.
// 🏗️ Generated from JSON Schema

package models

import (
	"encoding/json"
	"gitlab.com/tozd/go/errors"
)

type OneOfExample struct {
	Identifier OneOfExample_Identifier_OneOf $$$json:"identifier,omitempty"$$$
	Status OneOfExample_Status_OneOf $$$json:"status,omitempty"$$$
}

type OneOfExample_Identifier_OneOf struct {
	StringValue *string $$$json:"string_value,omitempty"$$$
	IntegerValue *int $$$json:"integer_value,omitempty"$$$
}

type OneOfExample_Status_OneOf struct {
	BooleanValue *bool $$$json:"boolean_value,omitempty"$$$
	EnumValue *OneOfExample_Status_OneOf_Enum $$$json:"enum_value,omitempty"$$$
}

type OneOfExample_Status_OneOf_Enum string

const (
	OneOfExample_Status_OneOf_Enum_Pending OneOfExample_Status_OneOf_Enum = "pending"
	OneOfExample_Status_OneOf_Enum_Failed  OneOfExample_Status_OneOf_Enum = "failed"
)

func (o *OneOfExample_Identifier_OneOf) UnmarshalJSON(data []byte) error {
	return nil // TODO: Implement
}

func (o OneOfExample_Identifier_OneOf) MarshalJSON() ([]byte, error) {
	return nil, nil // TODO: Implement
}

func (o *OneOfExample_Status_OneOf) UnmarshalJSON(data []byte) error {
	return nil // TODO: Implement
}

func (o OneOfExample_Status_OneOf) MarshalJSON() ([]byte, error) {
	return nil, nil // TODO: Implement
}`
}
