// Code generated by github.com/atombender/go-jsonschema, DO NOT EDIT.

package test

type ExtRef struct {
	// MyThing corresponds to the JSON schema field "myThing".
	MyThing *Thing `json:"myThing,omitempty" yaml:"myThing,omitempty" mapstructure:"myThing,omitempty"`

	// MyThing2 corresponds to the JSON schema field "myThing2".
	MyThing2 *Thing `json:"myThing2,omitempty" yaml:"myThing2,omitempty" mapstructure:"myThing2,omitempty"`
}

type RefExternalFile struct {
	// MyExternalThing corresponds to the JSON schema field "myExternalThing".
	MyExternalThing *Thing `json:"myExternalThing,omitempty" yaml:"myExternalThing,omitempty" mapstructure:"myExternalThing,omitempty"`

	// SomeOtherExternalThing corresponds to the JSON schema field
	// "someOtherExternalThing".
	SomeOtherExternalThing *Thing `json:"someOtherExternalThing,omitempty" yaml:"someOtherExternalThing,omitempty" mapstructure:"someOtherExternalThing,omitempty"`
}

type Thing struct {
	// Name corresponds to the JSON schema field "name".
	Name *string `json:"name,omitempty" yaml:"name,omitempty" mapstructure:"name,omitempty"`
}
