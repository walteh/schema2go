// Code generated by github.com/atombender/go-jsonschema, DO NOT EDIT.

package test

import "encoding/json"
import "fmt"

type RequiredNullable struct {
	// MyNullableObject corresponds to the JSON schema field "myNullableObject".
	MyNullableObject RequiredNullableMyNullableObject `json:"myNullableObject" yaml:"myNullableObject" mapstructure:"myNullableObject"`

	// MyNullableString corresponds to the JSON schema field "myNullableString".
	MyNullableString *string `json:"myNullableString" yaml:"myNullableString" mapstructure:"myNullableString"`

	// MyNullableStringArray corresponds to the JSON schema field
	// "myNullableStringArray".
	MyNullableStringArray []string `json:"myNullableStringArray" yaml:"myNullableStringArray" mapstructure:"myNullableStringArray"`
}

type RequiredNullableMyNullableObject struct {
	// MyNestedProp corresponds to the JSON schema field "myNestedProp".
	MyNestedProp string `json:"myNestedProp" yaml:"myNestedProp" mapstructure:"myNestedProp"`
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *RequiredNullableMyNullableObject) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if _, ok := raw["myNestedProp"]; raw != nil && !ok {
		return fmt.Errorf("field myNestedProp in RequiredNullableMyNullableObject: required")
	}
	type Plain RequiredNullableMyNullableObject
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = RequiredNullableMyNullableObject(plain)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *RequiredNullable) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if _, ok := raw["myNullableObject"]; raw != nil && !ok {
		return fmt.Errorf("field myNullableObject in RequiredNullable: required")
	}
	if _, ok := raw["myNullableString"]; raw != nil && !ok {
		return fmt.Errorf("field myNullableString in RequiredNullable: required")
	}
	if _, ok := raw["myNullableStringArray"]; raw != nil && !ok {
		return fmt.Errorf("field myNullableStringArray in RequiredNullable: required")
	}
	type Plain RequiredNullable
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = RequiredNullable(plain)
	return nil
}
