// Code generated by github.com/walteh/schema2go, DO NOT EDIT.

package test

// A simple schema.
type Description struct {
	// MyDescriptionlessField corresponds to the JSON schema field
	// "myDescriptionlessField".
	MyDescriptionlessField *string `json:"myDescriptionlessField,omitempty" yaml:"myDescriptionlessField,omitempty" mapstructure:"myDescriptionlessField,omitempty"`

	// A string field.
	MyField *string `json:"myField,omitempty" yaml:"myField,omitempty" mapstructure:"myField,omitempty"`
}
