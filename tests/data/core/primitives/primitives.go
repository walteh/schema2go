// Code generated by github.com/walteh/schema2go, DO NOT EDIT.

package test

import "encoding/json"
import "fmt"
import yaml "gopkg.in/yaml.v3"

type Primitives struct {
	// MyBoolean corresponds to the JSON schema field "myBoolean".
	MyBoolean *bool `json:"myBoolean,omitempty" yaml:"myBoolean,omitempty" mapstructure:"myBoolean,omitempty"`

	// MyInteger corresponds to the JSON schema field "myInteger".
	MyInteger *int `json:"myInteger,omitempty" yaml:"myInteger,omitempty" mapstructure:"myInteger,omitempty"`

	// MyNull corresponds to the JSON schema field "myNull".
	MyNull interface{} `json:"myNull,omitempty" yaml:"myNull,omitempty" mapstructure:"myNull,omitempty"`

	// MyNumber corresponds to the JSON schema field "myNumber".
	MyNumber *float64 `json:"myNumber,omitempty" yaml:"myNumber,omitempty" mapstructure:"myNumber,omitempty"`

	// MyString corresponds to the JSON schema field "myString".
	MyString *string `json:"myString,omitempty" yaml:"myString,omitempty" mapstructure:"myString,omitempty"`
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *Primitives) UnmarshalJSON(value []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(value, &raw); err != nil {
		return err
	}
	type Plain Primitives
	var plain Plain
	if err := json.Unmarshal(value, &plain); err != nil {
		return err
	}
	if plain.MyNull != nil {
		return fmt.Errorf("field %s: must be null", "myNull")
	}
	*j = Primitives(plain)
	return nil
}

// UnmarshalYAML implements yaml.Unmarshaler.
func (j *Primitives) UnmarshalYAML(value *yaml.Node) error {
	var raw map[string]interface{}
	if err := value.Decode(&raw); err != nil {
		return err
	}
	type Plain Primitives
	var plain Plain
	if err := value.Decode(&plain); err != nil {
		return err
	}
	if plain.MyNull != nil {
		return fmt.Errorf("field %s: must be null", "myNull")
	}
	*j = Primitives(plain)
	return nil
}
