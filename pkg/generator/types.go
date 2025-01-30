package generator

// SchemaModel represents a parsed JSON Schema ready for code generation
type SchemaModel struct {
	// Package is the Go package name
	Package string

	// Structs are the Go structs to generate
	Structs []*StructModel

	// Imports are the required Go imports
	Imports []string
}

// StructModel represents a Go struct to generate
type StructModel struct {
	// Name is the Go struct name
	Name string

	// Description is the struct's documentation
	Description string

	// Fields are the struct's fields
	Fields []*FieldModel

	// HasValidation indicates if the struct needs validation
	HasValidation bool

	// HasDefaults indicates if the struct needs a constructor
	HasDefaults bool

	// HasCustomMarshaling indicates if the struct needs custom JSON marshaling
	HasCustomMarshaling bool
}

// FieldModel represents a Go struct field
type FieldModel struct {
	// Name is the Go field name
	Name string

	// Type is the Go type (e.g., string, int, *CustomType)
	Type string

	// JSONName is the name in the JSON
	JSONName string

	// Description is the field's documentation
	Description string

	// IsRequired indicates if the field is required
	IsRequired bool

	// DefaultValue is the field's default value, if any
	DefaultValue interface{}

	// ValidationRules are the field's validation rules
	ValidationRules []*ValidationRule
}

// ValidationRule represents a validation rule for a field
type ValidationRule struct {
	// Type is the type of validation (e.g., "required", "min", "max", "pattern")
	Type string

	// Value is the validation value
	Value interface{}

	// Message is the error message
	Message string
}

// ValidationTypes
const (
	ValidationRequired = "required"
	ValidationMin      = "min"
	ValidationMax      = "max"
	ValidationPattern  = "pattern"
	ValidationMinLen   = "minLength"
	ValidationMaxLen   = "maxLength"
	ValidationEnum     = "enum"
	ValidationMultiple = "multipleOf"
)
