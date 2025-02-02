// Code generated by github.com/walteh/schema2go, DO NOT EDIT.

package test

import "encoding/json"
import "fmt"

type ExclusiveMinimum struct {
	// MyInteger corresponds to the JSON schema field "myInteger".
	MyInteger int `json:"myInteger" yaml:"myInteger" mapstructure:"myInteger"`

	// MyNullableInteger corresponds to the JSON schema field "myNullableInteger".
	MyNullableInteger *int `json:"myNullableInteger,omitempty" yaml:"myNullableInteger,omitempty" mapstructure:"myNullableInteger,omitempty"`

	// MyNullableNumber corresponds to the JSON schema field "myNullableNumber".
	MyNullableNumber *float64 `json:"myNullableNumber,omitempty" yaml:"myNullableNumber,omitempty" mapstructure:"myNullableNumber,omitempty"`

	// MyNumber corresponds to the JSON schema field "myNumber".
	MyNumber float64 `json:"myNumber" yaml:"myNumber" mapstructure:"myNumber"`
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *ExclusiveMinimum) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if _, ok := raw["myInteger"]; raw != nil && !ok {
		return fmt.Errorf("field myInteger in ExclusiveMinimum: required")
	}
	if _, ok := raw["myNumber"]; raw != nil && !ok {
		return fmt.Errorf("field myNumber in ExclusiveMinimum: required")
	}
	type Plain ExclusiveMinimum
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	if 2 >= plain.MyInteger {
		return fmt.Errorf("field %s: must be > %v", "myInteger", 2)
	}
	if plain.MyNullableInteger != nil && 2 >= *plain.MyNullableInteger {
		return fmt.Errorf("field %s: must be > %v", "myNullableInteger", 2)
	}
	if plain.MyNullableNumber != nil && 1.2 >= *plain.MyNullableNumber {
		return fmt.Errorf("field %s: must be > %v", "myNullableNumber", 1.2)
	}
	if 1.2 >= plain.MyNumber {
		return fmt.Errorf("field %s: must be > %v", "myNumber", 1.2)
	}
	*j = ExclusiveMinimum(plain)
	return nil
}
