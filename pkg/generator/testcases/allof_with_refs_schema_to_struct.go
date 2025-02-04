package testcases

import (
	"github.com/google/gnostic/jsonschema"
	"github.com/walteh/schema2go/pkg/generator"
)

type allof_with_refs_schema_to_struct struct{}

func init() {
	registerTestCase(&allof_with_refs_schema_to_struct{})
}

func (t *allof_with_refs_schema_to_struct) Name() string {
	return myfilename()
}

func (t *allof_with_refs_schema_to_struct) JSONSchema() string {
	return `{
		"$schema": "http://json-schema.org/draft-07/schema#",
		"title": "AllOfWithRefsExample",
		"type": "object",
		"allOf": [
			{ "$ref": "#/definitions/PersonInfo" },
			{ "$ref": "#/definitions/EmployeeInfo" }
		],
		"definitions": {
			"PersonInfo": {
				"type": "object",
				"properties": {
					"name": {
						"type": "string"
					},
					"age": {
						"type": "integer"
					}
				}
			},
			"EmployeeInfo": {
				"type": "object",
				"properties": {
					"employeeId": {
						"type": "string"
					},
					"department": {
						"type": "string"
					}
				}
			}
		}
	}`
}

func (t *allof_with_refs_schema_to_struct) RawSchema() *jsonschema.Schema {
	return &jsonschema.Schema{
		Schema: strPtr("http://json-schema.org/draft-07/schema#"),
		Title:  strPtr("AllOfWithRefsExample"),
		Type:   typePtr("object"),
		AllOf: &[]*jsonschema.Schema{
			{Ref: strPtr("#/definitions/PersonInfo")},
			{Ref: strPtr("#/definitions/EmployeeInfo")},
		},
		Definitions: &[]*jsonschema.NamedSchema{
			{
				Name: "PersonInfo",
				Value: &jsonschema.Schema{
					Type: typePtr("object"),
					Properties: &[]*jsonschema.NamedSchema{
						{
							Name: "name",
							Value: &jsonschema.Schema{
								Type: typePtr("string"),
							},
						},
						{
							Name: "age",
							Value: &jsonschema.Schema{
								Type: typePtr("integer"),
							},
						},
					},
				},
			},
			{
				Name: "EmployeeInfo",
				Value: &jsonschema.Schema{
					Type: typePtr("object"),
					Properties: &[]*jsonschema.NamedSchema{
						{
							Name: "employeeId",
							Value: &jsonschema.Schema{
								Type: typePtr("string"),
							},
						},
						{
							Name: "department",
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

func (t *allof_with_refs_schema_to_struct) StaticSchema() *generator.StaticSchema {
	return &generator.StaticSchema{
		Package_: "models",
		Structs_: []generator.Struct{
			&generator.StaticStruct{
				Name_:        "AllOfWithRefsExample",
				Description_: "",
				Fields_: []generator.Field{
					&generator.StaticField{
						Name_:                "EmployeeInfo_AllOf",
						JSONName_:            "",
						Description_:         "",
						IsRequired_:          false,
						Type_:                "EmployeeInfo",
						IsEnum_:              false,
						EnumTypeName_:        "",
						EnumValues_:          nil,
						DefaultValue_:        nil,
						DefaultValueComment_: nil,
						ValidationRules_:     nil,
					},
					&generator.StaticField{
						Name_:                "PersonInfo_AllOf",
						JSONName_:            "",
						Description_:         "",
						IsRequired_:          false,
						Type_:                "PersonInfo",
						IsEnum_:              false,
						EnumTypeName_:        "",
						EnumValues_:          nil,
						DefaultValue_:        nil,
						DefaultValueComment_: nil,
						ValidationRules_:     nil,
					},
				},
				HasAllOf_:            true,
				HasCustomMarshaling_: true,
				HasDefaults_:         false,
				HasValidation_:       false,
			},
			&generator.StaticStruct{
				Name_:        "PersonInfo",
				Description_: "",
				Fields_: []generator.Field{
					&generator.StaticField{
						Name_:                "Age",
						JSONName_:            "age",
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
				Name_:        "EmployeeInfo",
				Description_: "",
				Fields_: []generator.Field{
					&generator.StaticField{
						Name_:                "Department",
						JSONName_:            "department",
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
						Name_:                "EmployeeID",
						JSONName_:            "employeeId",
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
		},
		Imports_: []string{
			"encoding/json",
			"gitlab.com/tozd/go/errors",
		},
	}
}

func (t *allof_with_refs_schema_to_struct) GoCode() string {
	return `// Code generated by schema2go. DO NOT EDIT.
// 🏗️ Generated from JSON Schema

package models

import (
	"encoding/json"
	"gitlab.com/tozd/go/errors"
)

type AllOfWithRefsExample struct {
	EmployeeInfo_AllOf EmployeeInfo
	PersonInfo_AllOf   PersonInfo
}

func (x *AllOfWithRefsExample) UnmarshalJSON(data []byte) error {
	return nil // TODO: Implement custom unmarshaling for allOf fields
}

func (x AllOfWithRefsExample) MarshalJSON() ([]byte, error) {
	return nil, nil // TODO: Implement custom marshaling for allOf fields
}

type EmployeeInfo struct {
	Department *string $$$json:"department,omitempty"$$$
	EmployeeID *string $$$json:"employeeId,omitempty"$$$
}

type PersonInfo struct {
	Age  *int    $$$json:"age,omitempty"$$$
	Name *string $$$json:"name,omitempty"$$$
}`
}
