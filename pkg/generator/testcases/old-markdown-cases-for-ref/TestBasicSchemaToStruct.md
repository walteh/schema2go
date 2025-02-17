# json-schema

```json
{
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
}
```

---

# go-code

```go
// Code generated by schema2go. DO NOT EDIT.
// 🏗️ Generated from JSON Schema

package models

import (
	"encoding/json"
	"gitlab.com/tozd/go/errors"
)

type BasicExample struct {
	ID string `json:"id"` // Required
	Count *int `json:"count,omitempty"` // Default: 0
	Enabled *bool `json:"enabled,omitempty"` // Default: false
	Ratio *float64 `json:"ratio,omitempty"`
}

// UnmarshalJSON implements json.Unmarshaler
func (x *BasicExample) UnmarshalJSON(data []byte) error {
	// Define an alias to prevent recursive UnmarshalJSON calls
	type Alias BasicExample
	aux := &struct {
		*Alias
	}{
		Alias: (*Alias)(x),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return errors.Errorf("unmarshaling json: %w", err)
	}

	// Apply defaults for missing fields
	if x.Count == nil {
		defaultValue := 0
		x.Count = &defaultValue
	}
	if x.Enabled == nil {
		defaultValue := false
		x.Enabled = &defaultValue
	}

	// Validate after applying defaults
	if err := x.Validate(); err != nil {
		return errors.Errorf("validating after unmarshal: %w", err)
	}

	return nil
}

// Validate ensures all required fields are present and valid
func (x *BasicExample) Validate() error {
	if x.ID == "" {
		return errors.New("id is required")
	}
	return nil
}

// MarshalJSON implements json.Marshaler
func (x BasicExample) MarshalJSON() ([]byte, error) {
	// Validate before marshaling
	if err := x.Validate(); err != nil {
		return nil, errors.Errorf("validating before marshal: %w", err)
	}

	// Use alias to avoid infinite recursion
	type Alias BasicExample
	return json.Marshal((*Alias)(&x))
}
```

---

# raw-schema

```go
&jsonschema.Schema{
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
					Kind: yaml.ScalarNode,
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
					Kind: yaml.ScalarNode,
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
```

---

# static-schema

```go
f1 := &generator.StaticField{
					Name_: "ID",
					JSONName_: "id",
					Description_: "",
					IsRequired_: true,
					Type_: "string",
					IsEnum_: false,
					EnumTypeName_: "IDType",
					EnumValues_: nil,
					DefaultValue_: nil,
					DefaultValueComment_: nil,
					ValidationRules_: []generator.ValidationRule{
						{
							Type: generator.ValidationRequired,
							Message: "id is required",
							// Parnet: will be injected by the test case
							Values: "",
						},
					},
				}
f2 := &generator.StaticField{
	Name_: "Count",
	JSONName_: "count",
	Description_: "",
	IsRequired_: false,
	Type_: "*int",
	IsEnum_: false,
	EnumTypeName_: "CountType",
	EnumValues_: nil,
	DefaultValue_: ptr("0"),
	DefaultValueComment_: ptr("0"),
	ValidationRules_: nil,
}
f3 := &generator.StaticField{
					Name_: "Enabled",
					JSONName_: "enabled",
					Description_: "",
					IsRequired_: false,
					Type_: "*bool",
					IsEnum_: false,
					EnumTypeName_: "EnabledType",
					EnumValues_: nil,
					DefaultValue_: ptr("false"),
					DefaultValueComment_: ptr("false"),
					ValidationRules_: nil,
				}
f4 := &generator.StaticField{
					Name_: "Ratio",
					JSONName_: "ratio",
					Description_: "",
					IsRequired_: false,
					Type_: "*float64",
					IsEnum_: false,
					EnumTypeName_: "RatioType",
					EnumValues_: nil,
					DefaultValue_: nil,
					DefaultValueComment_: nil,
					ValidationRules_: nil,
				}

s1 := &generator.StaticStruct{Name_: "BasicExample", Fields_: []generator.Field{f1, f2, f3, f4},
HasDefaults_: true,
					HasValidation_: true,
					HasCustomMarshaling_: true,}

staticWant := &generator.StaticSchema{
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
```
