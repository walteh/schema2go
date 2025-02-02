// Code generated by github.com/walteh/schema2go, DO NOT EDIT.

package test

import "encoding/json"
import "fmt"
import "reflect"

type TypedDefaultEnums struct {
	// Some corresponds to the JSON schema field "some".
	Some TypedDefaultEnumsSome `json:"some,omitempty" yaml:"some,omitempty" mapstructure:"some,omitempty"`
}

type TypedDefaultEnumsSome string

const TypedDefaultEnumsSomeOther TypedDefaultEnumsSome = "other"
const TypedDefaultEnumsSomeRandom TypedDefaultEnumsSome = "random"

var enumValues_TypedDefaultEnumsSome = []interface{}{
	"random",
	"other",
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *TypedDefaultEnumsSome) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_TypedDefaultEnumsSome {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_TypedDefaultEnumsSome, v)
	}
	*j = TypedDefaultEnumsSome(v)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *TypedDefaultEnums) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	type Plain TypedDefaultEnums
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	if v, ok := raw["some"]; !ok || v == nil {
		plain.Some = "random"
	}
	*j = TypedDefaultEnums(plain)
	return nil
}
