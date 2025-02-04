package testcases

import (
	"github.com/google/gnostic/jsonschema"
	"github.com/walteh/schema2go/internal/archives/generator"
	"gopkg.in/yaml.v3"
)

type basic_schema_to_struct struct {
}

func init() {
	registerTestCase(&basic_schema_to_struct{})
}

func (me *basic_schema_to_struct) Name() string {
	return myfilename()
}

func (*basic_schema_to_struct) JSONSchema() string {
	return `{
		"$schema": "http://json-schema.org/draft-07/schema#",
		"title": "BasicExample",
		"type": "object",
		"required": ["id"],
		"properties": {
			"id": {
				"type": "string"
			},
			"count": {
				"type": "integer",
				"default": 0
			},
			"enabled": {
				"type": "boolean",
				"default": false
			},
			"ratio": {
				"type": "number"
			}
		}
	}`
}

func (*basic_schema_to_struct) RawSchema() *jsonschema.Schema {
	return &jsonschema.Schema{
		Schema: ptr("http://json-schema.org/draft-07/schema#"),
		Title:  ptr("BasicExample"),
		Type:   typ("object"),
		Properties: &[]*jsonschema.NamedSchema{
			{
				Name: "id",
				Value: &jsonschema.Schema{
					Type: typ("string"),
				},
			},
			{
				Name: "count",
				Value: &jsonschema.Schema{
					Type: typ("integer"),
					Default: &yaml.Node{
						Kind:  yaml.ScalarNode,
						Value: "0",
						Tag:   "!!int",
					},
				},
			},
			{
				Name: "enabled",
				Value: &jsonschema.Schema{
					Type: typ("boolean"),
					Default: &yaml.Node{
						Kind:  yaml.ScalarNode,
						Value: "false",
						Tag:   "!!bool",
					},
				},
			},
			{
				Name: "ratio",
				Value: &jsonschema.Schema{
					Type: typ("number"),
				},
			},
		},
		Required: &[]string{"id"},
	}
}

func (*basic_schema_to_struct) StaticSchema() *generator.StaticSchema {
	f1 := &generator.StaticField{
		Name_:                "ID",
		JSONName_:            "id",
		Description_:         "",
		IsRequired_:          true,
		Type_:                "string",
		IsEnum_:              false,
		EnumValues_:          nil,
		DefaultValue_:        nil,
		DefaultValueComment_: nil,
		ValidationRules_: []generator.ValidationRule{
			{
				Type:    generator.ValidationRequired,
				Message: "id is required",
				Values:  "",
			},
		},
	}
	f2 := &generator.StaticField{
		Name_:                "Count",
		JSONName_:            "count",
		Description_:         "",
		IsRequired_:          false,
		Type_:                "*int",
		IsEnum_:              false,
		EnumValues_:          nil,
		DefaultValue_:        ptr("0"),
		DefaultValueComment_: ptr("0"),
		ValidationRules_:     nil,
	}
	f3 := &generator.StaticField{
		Name_:                "Enabled",
		JSONName_:            "enabled",
		Description_:         "",
		IsRequired_:          false,
		Type_:                "*bool",
		IsEnum_:              false,
		EnumValues_:          nil,
		DefaultValue_:        ptr("false"),
		DefaultValueComment_: ptr("false"),
		ValidationRules_:     nil,
	}
	f4 := &generator.StaticField{
		Name_:                "Ratio",
		JSONName_:            "ratio",
		Description_:         "",
		IsRequired_:          false,
		Type_:                "*float64",
		IsEnum_:              false,
		EnumValues_:          nil,
		DefaultValue_:        nil,
		DefaultValueComment_: nil,
		ValidationRules_:     nil,
	}

	s1 := &generator.StaticStruct{
		Name_:                "BasicExample",
		Fields_:              []generator.Field{f1, f2, f3, f4},
		HasDefaults_:         true,
		HasValidation_:       true,
		HasCustomMarshaling_: true,
	}

	return &generator.StaticSchema{
		Package_: "models",
		Structs_: []generator.Struct{
			s1,
		},
		Enums_: nil,
		Imports_: []string{
			"encoding/json",
			"gitlab.com/tozd/go/errors",
		},
	}
}

func (*basic_schema_to_struct) GoCode() string {
	return `package models

import (
	"encoding/json"
	"gitlab.com/tozd/go/errors"
)

type BasicExample struct {
	ID      string
	Count   *int
	Enabled *bool
	Ratio   *float64
}

func (x *BasicExample) UnmarshalJSON(data []byte) error {
	type Alias BasicExample
	aux := &struct {
		*Alias
		ID      string   $$$json:"id"$$$
		Count   *int     $$$json:"count,omitempty"$$$
		Enabled *bool    $$$json:"enabled,omitempty"$$$
		Ratio   *float64 $$$json:"ratio,omitempty"$$$
	}{
		Alias: (*Alias)(x),
	}

	if err := json.Unmarshal(data, &aux); err != nil {
		return errors.Errorf("unmarshaling json: %w", err)
	}

	if aux.Count == nil {
		defaultValue := 0
		aux.Count = &defaultValue
	}

	if aux.Enabled == nil {
		defaultValue := false
		aux.Enabled = &defaultValue
	}

	if err := x.Validate(); err != nil {
		return errors.Errorf("validating after unmarshal: %w", err)
	}

	return nil
}

func (x *BasicExample) Validate() error {
	if x.ID == "" {
		return errors.Errorf("id is required")
	}
	return nil
}

func (x BasicExample) MarshalJSON() ([]byte, error) {
	if err := x.Validate(); err != nil {
		return nil, errors.Errorf("validating before marshal: %w", err)
	}

	type Alias BasicExample
	return json.Marshal((*Alias)(&x))
}`
}
