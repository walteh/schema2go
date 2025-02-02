// Code generated by github.com/walteh/schema2go, DO NOT EDIT.

package test

import "encoding/json"
import "fmt"

type Bar struct {
	// RefToFoo corresponds to the JSON schema field "refToFoo".
	RefToFoo *Foo `json:"refToFoo,omitempty" yaml:"refToFoo,omitempty" mapstructure:"refToFoo,omitempty"`
}

type CyclicAndRequired1 struct {
	// A corresponds to the JSON schema field "a".
	A *Foo `json:"a,omitempty" yaml:"a,omitempty" mapstructure:"a,omitempty"`
}

type Foo struct {
	// RefToBar corresponds to the JSON schema field "refToBar".
	RefToBar Bar `json:"refToBar" yaml:"refToBar" mapstructure:"refToBar"`
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *Foo) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if _, ok := raw["refToBar"]; raw != nil && !ok {
		return fmt.Errorf("field refToBar in Foo: required")
	}
	type Plain Foo
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = Foo(plain)
	return nil
}
