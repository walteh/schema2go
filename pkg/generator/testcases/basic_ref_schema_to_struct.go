package testcases

import (
	"github.com/google/gnostic/jsonschema"
	"github.com/walteh/schema2go/pkg/generator"
)

type basic_ref_schema_to_struct struct {
}

func init() {
	registerTestCase(&basic_ref_schema_to_struct{})
}

func (me *basic_ref_schema_to_struct) Name() string {
	return myfilename()
}

func (*basic_ref_schema_to_struct) JSONSchema() string {
	return `{
		"$schema": "http://json-schema.org/draft-07/schema#",
		"title": "RefExample",
		"type": "object",
		"properties": {
			"person": {
				"$ref": "#/definitions/Person"
			},
			"innerPerson": {
				"type": "object",
				"properties": {
					"name": {
						"type": "string"
					},
					"age": {
						"type": "integer"
					}
				}
			}
		},
		"definitions": {
			"Person": {
				"type": "object",
				"properties": {
					"name": {
						"type": "string"
					},
					"age": {
						"type": "integer"
					}
				}
			}
		}
	}`
}

func (*basic_ref_schema_to_struct) RawSchema() *jsonschema.Schema {
	return &jsonschema.Schema{
		Schema: ptr("http://json-schema.org/draft-07/schema#"),
		Title:  ptr("RefExample"),
		Type:   typ("object"),
		Properties: &[]*jsonschema.NamedSchema{
			{
				Name: "person",
				Value: &jsonschema.Schema{
					Ref: ptr("#/definitions/Person"),
				},
			},
			{
				Name: "innerPerson",
				Value: &jsonschema.Schema{
					Type: typ("object"),
					Properties: &[]*jsonschema.NamedSchema{
						{
							Name: "name",
							Value: &jsonschema.Schema{
								Type: typ("string"),
							},
						},
						{
							Name: "age",
							Value: &jsonschema.Schema{
								Type: typ("integer"),
							},
						},
					},
				},
			},
		},
		Definitions: &[]*jsonschema.NamedSchema{
			{
				Name: "Person",
				Value: &jsonschema.Schema{
					Type: typ("object"),
					Properties: &[]*jsonschema.NamedSchema{
						{
							Name: "name",
							Value: &jsonschema.Schema{
								Type: typ("string"),
							},
						},
						{
							Name: "age",
							Value: &jsonschema.Schema{
								Type: typ("integer"),
							},
						},
					},
				},
			},
		},
	}
}

func (*basic_ref_schema_to_struct) StaticSchema() *generator.StaticSchema {
	// 🏗️ Person struct (from definitions)
	personFields := []generator.Field{
		&generator.StaticField{
			Name_:        "Name",
			JSONName_:    "name",
			Description_: "",
			IsRequired_:  false,
			Type_:        "*string",
		},
		&generator.StaticField{
			Name_:        "Age",
			JSONName_:    "age",
			Description_: "",
			IsRequired_:  false,
			Type_:        "*int",
		},
	}

	// 🏗️ Inner person struct (nested object)
	innerPersonFields := []generator.Field{
		&generator.StaticField{
			Name_:        "Name",
			JSONName_:    "name",
			Description_: "",
			IsRequired_:  false,
			Type_:        "*string",
		},
		&generator.StaticField{
			Name_:        "Age",
			JSONName_:    "age",
			Description_: "",
			IsRequired_:  false,
			Type_:        "*int",
		},
	}

	// 🏗️ Root struct fields
	refExampleFields := []generator.Field{
		&generator.StaticField{
			Name_:        "Person",
			JSONName_:    "person",
			Description_: "",
			IsRequired_:  false,
			Type_:        "*Person",
			ValidationRules_: []generator.ValidationRule{
				{
					Type:    "nested",
					Message: "validating person",
				},
			},
		},
		&generator.StaticField{
			Name_:        "InnerPerson",
			JSONName_:    "innerPerson",
			Description_: "",
			IsRequired_:  false,
			Type_:        "*RefExampleInnerPerson",
			ValidationRules_: []generator.ValidationRule{
				{
					Type:    "nested",
					Message: "validating innerPerson",
				},
			},
		},
	}

	// 📝 Person struct (from definitions)
	// - Basic validation for fields
	// - Custom marshaling for proper JSON handling
	personStruct := &generator.StaticStruct{
		Name_:                "Person",
		Fields_:              personFields,
		HasDefaults_:         false,
		HasValidation_:       false, // Should validate its own fields
		HasCustomMarshaling_: true,
	}

	// 📝 Root struct
	// - Validates nested objects
	// - Custom marshaling for proper JSON handling
	refExampleStruct := &generator.StaticStruct{
		Name_:                "RefExample",
		Fields_:              refExampleFields,
		HasDefaults_:         false,
		HasValidation_:       true, // Should validate nested objects
		HasCustomMarshaling_: true,
	}

	// 📝 Inner person struct (nested object)
	// - Basic validation for fields
	// - Custom marshaling for proper JSON handling
	innerPersonStruct := &generator.StaticStruct{
		Name_:                "RefExampleInnerPerson",
		Fields_:              innerPersonFields,
		HasDefaults_:         false,
		HasValidation_:       false, // Should validate its own fields
		HasCustomMarshaling_: true,
	}

	return &generator.StaticSchema{
		Package_: "models",
		// Order matters:
		// 1. Referenced types (Person)
		// 2. Root type (RefExample)
		// 3. Nested types (RefExampleInnerPerson)
		Structs_: []generator.Struct{personStruct, refExampleStruct, innerPersonStruct},
		Enums_:   nil, // No enums in this schema
		Imports_: []string{
			"encoding/json",
			"gitlab.com/tozd/go/errors",
		},
	}
}

func (*basic_ref_schema_to_struct) GoCode() string {
	return `
	// Code generated by schema2go. DO NOT EDIT.
	// 🏗️ Generated from JSON Schema
	
	package models
	
	import (
		"encoding/json"
		"gitlab.com/tozd/go/errors"
	)
	
	type Person struct {
		Name *string $$$json:"name,omitempty"$$$
		Age  *int    $$$json:"age,omitempty"$$$
	}
	
	// UnmarshalJSON implements json.Unmarshaler
	func (x *Person) UnmarshalJSON(data []byte) error {
		// Create an alias to avoid infinite recursion
		type Alias Person
		aux := &struct {
			*Alias
		}{
			Alias: (*Alias)(x),
		}
	
		// First unmarshal into our alias struct
		if err := json.Unmarshal(data, &aux); err != nil {
			return errors.Errorf("unmarshaling json: %w", err)
		}
	
		// Validate after unmarshaling
		if err := x.Validate(); err != nil {
			return errors.Errorf("validating after unmarshal: %w", err)
		}
	
		return nil
	}
	
	// Validate ensures all required fields are present and valid
	func (x *Person) Validate() error {
		return nil
	}
	
	// MarshalJSON implements json.Marshaler
	func (x Person) MarshalJSON() ([]byte, error) {
		// Validate before marshaling
		if err := x.Validate(); err != nil {
			return nil, errors.Errorf("validating before marshal: %w", err)
		}
	
		// Use alias to avoid infinite recursion
		type Alias Person
		return json.Marshal((*Alias)(&x))
	}
	
	type RefExample struct {
		Person      *Person                $$$json:"person,omitempty"$$$
		InnerPerson *RefExampleInnerPerson $$$json:"innerPerson,omitempty"$$$
	}
	
	// UnmarshalJSON implements json.Unmarshaler
	func (x *RefExample) UnmarshalJSON(data []byte) error {
		// Create an alias to avoid infinite recursion
		type Alias RefExample
		aux := &struct {
			*Alias
		}{
			Alias: (*Alias)(x),
		}
	
		// First unmarshal into our alias struct
		if err := json.Unmarshal(data, &aux); err != nil {
			return errors.Errorf("unmarshaling json: %w", err)
		}
	
		// Validate after unmarshaling
		if err := x.Validate(); err != nil {
			return errors.Errorf("validating after unmarshal: %w", err)
		}
	
		return nil
	}
	
	// Validate ensures all required fields are present and valid
	func (x *RefExample) Validate() error {
		if x.Person != nil {
			if err := x.Person.Validate(); err != nil {
				return errors.Errorf("validating person: %w", err)
			}
		}
		if x.InnerPerson != nil {
			if err := x.InnerPerson.Validate(); err != nil {
				return errors.Errorf("validating innerPerson: %w", err)
			}
		}
		return nil
	}
	
	// MarshalJSON implements json.Marshaler
	func (x RefExample) MarshalJSON() ([]byte, error) {
		// Validate before marshaling
		if err := x.Validate(); err != nil {
			return nil, errors.Errorf("validating before marshal: %w", err)
		}
	
		// Use alias to avoid infinite recursion
		type Alias RefExample
		return json.Marshal((*Alias)(&x))
	}
	
	type RefExampleInnerPerson struct {
		Name *string $$$json:"name,omitempty"$$$
		Age  *int    $$$json:"age,omitempty"$$$
	}
	
	// UnmarshalJSON implements json.Unmarshaler
	func (x *RefExampleInnerPerson) UnmarshalJSON(data []byte) error {
		// Create an alias to avoid infinite recursion
		type Alias RefExampleInnerPerson
		aux := &struct {
			*Alias
		}{
			Alias: (*Alias)(x),
		}
	
		// First unmarshal into our alias struct
		if err := json.Unmarshal(data, &aux); err != nil {
			return errors.Errorf("unmarshaling json: %w", err)
		}
	
		// Validate after unmarshaling
		if err := x.Validate(); err != nil {
			return errors.Errorf("validating after unmarshal: %w", err)
		}
	
		return nil
	}
	
	// Validate ensures all required fields are present and valid
	func (x *RefExampleInnerPerson) Validate() error {
		return nil
	}
	
	// MarshalJSON implements json.Marshaler
	func (x RefExampleInnerPerson) MarshalJSON() ([]byte, error) {
		// Validate before marshaling
		if err := x.Validate(); err != nil {
			return nil, errors.Errorf("validating before marshal: %w", err)
		}
	
		// Use alias to avoid infinite recursion
		type Alias RefExampleInnerPerson
		return json.Marshal((*Alias)(&x))
	}
	`
}
