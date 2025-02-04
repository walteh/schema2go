package testcases

import (
	"github.com/google/gnostic/jsonschema"
	"github.com/walteh/schema2go/pkg/generator"
)

type allof_schema_to_struct struct {
}

func init() {
	registerTestCase(&allof_schema_to_struct{})
}

func (me *allof_schema_to_struct) Name() string {
	return myfilename()
}

func (*allof_schema_to_struct) JSONSchema() string {
	return `{
		"$schema": "http://json-schema.org/draft-07/schema#",
		"title": "AllOfExample",
		"type": "object",
		"allOf": [
			{
				"type": "object",
				"properties": {
					"name": {
						"type": "string"
					}
				}
			},
			{
				"type": "object",
				"properties": {
					"age": {
						"type": "integer"
					}
				}
			}
		]
	}`
}

func (*allof_schema_to_struct) RawSchema() *jsonschema.Schema {
	return &jsonschema.Schema{
		Schema: ptr("http://json-schema.org/draft-07/schema#"),
		Title:  ptr("AllOfExample"),
		Type:   typ("object"),
		AllOf: &[]*jsonschema.Schema{
			{
				Type: typ("object"),
				Properties: &[]*jsonschema.NamedSchema{
					{Name: "name", Value: &jsonschema.Schema{Type: typ("string")}},
				},
			},
			{
				Type: typ("object"),
				Properties: &[]*jsonschema.NamedSchema{
					{Name: "age", Value: &jsonschema.Schema{Type: typ("integer")}},
				},
			},
		},
	}
}

func (*allof_schema_to_struct) StaticSchema() *generator.StaticSchema {
	f1 := &generator.StaticField{
		Name_:                "Name_AllOf",
		JSONName_:            "name",
		Description_:         "",
		IsRequired_:          false,
		Type_:                "*string",
		IsEnum_:              false,
		EnumTypeName_:        "NameType",
		EnumValues_:          nil,
		DefaultValue_:        nil,
		DefaultValueComment_: nil,
		ValidationRules_:     nil,
	}

	f2 := &generator.StaticField{
		Name_:                "Age_AllOf",
		JSONName_:            "age",
		Description_:         "",
		IsRequired_:          false,
		Type_:                "*int",
		IsEnum_:              false,
		EnumTypeName_:        "AgeType",
		EnumValues_:          nil,
		DefaultValue_:        nil,
		DefaultValueComment_: nil,
		ValidationRules_:     nil,
	}

	s1 := &generator.StaticStruct{
		Name_:                "AllOfExample",
		Fields_:              []generator.Field{f1, f2},
		HasDefaults_:         false,
		HasValidation_:       false,
		HasCustomMarshaling_: true,
		HasAllOf_:            true,
	}

	return &generator.StaticSchema{
		Package_: "models",
		Structs_: []generator.Struct{s1},
		Enums_:   nil,
		Imports_: []string{
			"encoding/json",
			"gitlab.com/tozd/go/errors",
		},
	}
}

func (*allof_schema_to_struct) GoCode() string {
	return `package models

import (
	"encoding/json"
	"gitlab.com/tozd/go/errors"
)

type AllOfExample struct {
	Name_AllOf *string
	Age_AllOf  *int
}

func (x *AllOfExample) UnmarshalJSON(data []byte) error {
	type Alias AllOfExample
	aux := &struct {
		*Alias
		Name_AllOf *string $$$json:"name,omitempty"$$$
		Age_AllOf  *int    $$$json:"age,omitempty"$$$
	}{
		Alias: (*Alias)(x),
	}

	if err := json.Unmarshal(data, &aux); err != nil {
		return errors.Errorf("unmarshaling json: %w", err)
	}

	if err := x.Validate(); err != nil {
		return errors.Errorf("validating after unmarshal: %w", err)
	}

	return nil
}

func (x *AllOfExample) Validate() error {
	return nil
}

func (x AllOfExample) MarshalJSON() ([]byte, error) {
	if err := x.Validate(); err != nil {
		return nil, errors.Errorf("validating before marshal: %w", err)
	}

	type Alias AllOfExample
	return json.Marshal((*Alias)(&x))
}`
}
