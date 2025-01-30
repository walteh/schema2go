package generator

// SchemaModel represents a parsed JSON Schema ready for code generation
type SchemaModel struct {
	// Package is the Go package name
	Package string

	// Structs are the Go structs to generate
	Structs []*StructModel

	// Enums are the Go enums to generate
	Enums []*EnumModel

	// Interfaces are the Go interfaces to generate
	Interfaces []*InterfaceModel

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

// EnumModel represents a Go enum to generate
type EnumModel struct {
	// Name is the Go enum type name
	Name string

	// Description is the enum's documentation
	Description string

	// BaseType is the underlying type (string, int, etc)
	BaseType string

	// Values are the enum values
	Values []*EnumValue
}

// EnumValue represents a single enum value
type EnumValue struct {
	// Name is the Go constant name
	Name string

	// Value is the actual value
	Value interface{}

	// Description is the value's documentation
	Description string
}

// InterfaceModel represents a Go interface to generate
type InterfaceModel struct {
	// Name is the Go interface name
	Name string

	// Description is the interface's documentation
	Description string

	// Methods are the interface methods
	Methods []*MethodModel
}

// MethodModel represents an interface method
type MethodModel struct {
	// Name is the method name
	Name string

	// Parameters are the method parameters
	Parameters []*ParameterModel

	// ReturnTypes are the method return types
	ReturnTypes []string

	// Description is the method's documentation
	Description string
}

// ParameterModel represents a method parameter
type ParameterModel struct {
	// Name is the parameter name
	Name string

	// Type is the parameter type
	Type string
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

// ValidationRuleType represents the type of validation rule
type ValidationRuleType string

const (
	ValidationRequired ValidationRuleType = "required"
	ValidationEnum     ValidationRuleType = "enum"
	ValidationNested   ValidationRuleType = "nested"
)

// ValidationRule represents a validation rule for a field
type ValidationRule struct {
	// Type is the type of validation rule
	Type ValidationRuleType

	// Message is the error message to display
	Message string

	// Field is the field being validated
	Field *FieldModel

	// Values is a comma-separated list of valid values for enum validation
	Values string
}

// ValidationTypes
const (
	ValidationMin      = "min"
	ValidationMax      = "max"
	ValidationPattern  = "pattern"
	ValidationMinLen   = "minLength"
	ValidationMaxLen   = "maxLength"
	ValidationMultiple = "multipleOf"
)
