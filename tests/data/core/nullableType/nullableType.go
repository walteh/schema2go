// Code generated by github.com/walteh/schema2go, DO NOT EDIT.

package test

type BoolThing *bool

type FloatThing *float64

type IntegerThing *int

type NullableType struct {
	// MyInlineStringValue corresponds to the JSON schema field "MyInlineStringValue".
	MyInlineStringValue *string `json:"MyInlineStringValue,omitempty" yaml:"MyInlineStringValue,omitempty" mapstructure:"MyInlineStringValue,omitempty"`

	// MyStringValue corresponds to the JSON schema field "MyStringValue".
	MyStringValue StringThing `json:"MyStringValue,omitempty" yaml:"MyStringValue,omitempty" mapstructure:"MyStringValue,omitempty"`
}

type StringThing *string
