// Code generated by github.com/walteh/schema2go, DO NOT EDIT.

package test

type RefToPrimitiveString struct {
	// MyThing corresponds to the JSON schema field "myThing".
	MyThing *Thing `json:"myThing,omitempty" yaml:"myThing,omitempty" mapstructure:"myThing,omitempty"`
}

type Thing string
