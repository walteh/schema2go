package generator

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/google/gnostic/jsonschema"
	"github.com/k0kubun/pp/v3"
	"github.com/walteh/schema2go/pkg/parser"
)

// SchemaModel represents a parsed JSON Schema ready for code generation
type SchemaModel struct {
	SourceSchema *jsonschema.Schema
}

// StructModel represents a Go struct to generate
type StructModel struct {
	SourceSchema *jsonschema.Schema
	ParentSchema *jsonschema.Schema // Parent schema that contains this schema
}

// FieldModel represents a Go struct field
type FieldModel struct {
	SourceSchema *jsonschema.Schema
	ParentSchema *jsonschema.Schema
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
	Type    ValidationRuleType
	Message string
	Field   *FieldModel
	Values  string
}

// EnumValue represents a single enum value
type EnumValue struct {
	Name        string
	Value       string
	Description string
	Parent      *EnumModel
}

// EnumModel represents a Go enum to generate
type EnumModel struct {
	Name        string
	BaseType    string
	Description string
	Values      []*EnumValue
}

// FieldModel Methods

func (f *FieldModel) Name() string {
	return toGoFieldName(f.JSONName())
}

func (f *FieldModel) JSONName() string {
	props := parser.GetProperties(f.ParentSchema)
	for name, prop := range props {
		if prop == f.SourceSchema {
			return name
		}
	}
	return "" // Should never happen if model is constructed correctly
}

func (f *FieldModel) Description() string {
	return parser.GetDescription(f.SourceSchema)
}

func (f *FieldModel) IsRequired() bool {
	return parser.IsRequired(f.ParentSchema, f.JSONName())
}

func (f *FieldModel) Type() string {
	if f.IsEnum() {
		return "*" + f.EnumTypeName()
	}

	// Get base type
	baseType := parser.GetTypeOrEmpty(f.SourceSchema)
	goType := ""

	// Map JSON Schema types to Go types
	switch baseType {
	case "string":
		goType = "string"
	case "integer":
		goType = "int"
	case "number":
		goType = "float64"
	case "boolean":
		goType = "bool"
	case "array":
		items := parser.GetArrayItems(f.SourceSchema)
		if items == nil || items.Schema == nil {
			goType = "[]interface{}" // Fallback for unknown array types
		} else {
			itemField := &FieldModel{
				SourceSchema: items.Schema,
				ParentSchema: f.SourceSchema,
			}
			goType = "[]" + itemField.Type()
		}
	case "object":
		// For nested objects, use the field name as the type name
		goType = toGoFieldName(f.JSONName())
	default:
		goType = "interface{}" // Fallback for unknown types
	}

	// Add pointer for optional fields
	if !f.IsRequired() && !strings.HasPrefix(goType, "[]") {
		goType = "*" + goType
	}

	return goType
}

func (f *FieldModel) IsEnum() bool {
	return len(parser.GetEnum(f.SourceSchema)) > 0
}

func (f *FieldModel) EnumTypeName() string {
	enumName := toGoFieldName(f.JSONName())
	if strings.HasSuffix(enumName, "Color") {
		enumName = "Color"
	}
	return enumName + "Type"
}

func (f *FieldModel) EnumValues() []*EnumValue {
	if !f.IsEnum() {
		return nil
	}

	enum := parser.GetEnum(f.SourceSchema)
	values := make([]*EnumValue, 0, len(enum))

	for _, val := range enum {
		if val.String == nil {
			continue
		}
		values = append(values, &EnumValue{
			Name:        fmt.Sprintf("%s%s", f.EnumTypeName(), toTitleCase(*val.String)),
			Value:       fmt.Sprintf("%q", *val.String),
			Description: fmt.Sprintf("Enum value %q", *val.String),
		})
	}
	return values
}

// DefaultValue returns the default value to use in code
func (f *FieldModel) DefaultValue() *string {
	if f.SourceSchema.Default == nil {
		return nil
	}

	// For enums, use the enum constant name
	if f.IsEnum() {
		v := fmt.Sprintf("%s%s", f.EnumTypeName(), toTitleCase(f.SourceSchema.Default.Value))
		return &v
	}

	v := f.SourceSchema.Default.Value
	return &v
}

// DefaultValueComment returns the raw default value for comments
func (f *FieldModel) DefaultValueComment() *string {
	if f.SourceSchema.Default == nil {
		return nil
	}

	v := f.SourceSchema.Default.Value
	return &v
}

func (f *FieldModel) ValidationRules() []*ValidationRule {
	var rules []*ValidationRule

	// Required validation
	if f.IsRequired() && parser.GetTypeOrEmpty(f.SourceSchema) != "object" {
		rules = append(rules, &ValidationRule{
			Type:    ValidationRequired,
			Message: fmt.Sprintf("%s is required", f.JSONName()),
			Field:   f,
		})
	}

	// Enum validation
	if f.IsEnum() {
		values := make([]string, 0)
		for _, v := range f.EnumValues() {
			values = append(values, v.Name)
		}
		rules = append(rules, &ValidationRule{
			Type:    ValidationEnum,
			Message: fmt.Sprintf("invalid %s", f.JSONName()),
			Field:   f,
			Values:  strings.Join(values, ", "),
		})
	}

	// Nested validation
	if parser.GetTypeOrEmpty(f.SourceSchema) == "object" {
		rules = append(rules, &ValidationRule{
			Type:    ValidationNested,
			Message: fmt.Sprintf("validating %s", f.JSONName()),
			Field:   f,
		})
	}

	return rules
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

// StructModel Methods

func (s *StructModel) Name() string {
	// For root schema, use title
	if title := parser.GetTitle(s.SourceSchema); title != "" {
		return title
	}

	// For nested objects, use field name directly if no title
	if s.ParentSchema != nil {
		props := parser.GetProperties(s.ParentSchema)
		for name, prop := range props {
			if prop == s.SourceSchema {
				return toGoFieldName(name)
			}
		}
	}

	// Fallback to "Object" if no name can be determined
	return "Object"
}

func (s *StructModel) Description() string {
	return parser.GetDescription(s.SourceSchema)
}

func (s *StructModel) Fields() []*FieldModel {
	props := parser.GetProperties(s.SourceSchema)
	fields := make([]*FieldModel, 0, len(props))

	// First add required fields in the order they appear in the required list
	required := parser.GetRequiredFields(s.SourceSchema)
	for _, name := range required {
		if prop, ok := props[name]; ok {
			fields = append(fields, &FieldModel{
				SourceSchema: prop,
				ParentSchema: s.SourceSchema,
			})
			delete(props, name) // Remove from props so we don't add it twice
		}
	}

	// Then add remaining fields in alphabetical order
	var optionalNames []string
	for name := range props {
		optionalNames = append(optionalNames, name)
	}
	sort.Strings(optionalNames)

	for _, name := range optionalNames {
		fields = append(fields, &FieldModel{
			SourceSchema: props[name],
			ParentSchema: s.SourceSchema,
		})
	}

	return fields
}

func (s *StructModel) HasValidation() bool {
	// Check if any fields are required or have validation rules
	for _, field := range s.Fields() {
		if field.IsRequired() || len(field.ValidationRules()) > 0 {
			return true
		}
	}
	return false
}

func (s *StructModel) HasDefaults() bool {
	// Check if any fields have default values
	for _, field := range s.Fields() {
		if field.DefaultValue() != nil {
			return true
		}
	}
	return false
}

func (s *StructModel) HasCustomMarshaling() bool {
	// Check if we need custom marshaling (e.g., for enums, validation, defaults)
	return s.HasValidation() || s.HasDefaults()
}

// SchemaModel Methods

func (s *SchemaModel) Package() string {
	return "models" // TODO: Make configurable
}

func (s *SchemaModel) Structs() []*StructModel {
	var structs []*StructModel
	seen := make(map[string]bool)

	// Helper function to recursively collect structs
	var collectStructs func(schema *jsonschema.Schema, parentSchema *jsonschema.Schema)
	collectStructs = func(schema *jsonschema.Schema, parentSchema *jsonschema.Schema) {
		// Skip if nil schema
		if schema == nil {
			return
		}

		// Get struct name
		name := ""
		if title := parser.GetTitle(schema); title != "" {
			name = title
		} else if parentSchema != nil {
			if parentTitle := parser.GetTitle(parentSchema); parentTitle != "" {
				props := parser.GetProperties(parentSchema)
				for propName, prop := range props {
					if prop == schema {
						name = parentTitle + "_" + toGoFieldName(propName)
						break
					}
				}
			}
		}

		// Skip if already processed
		if seen[name] {
			return
		}
		seen[name] = true

		// Add this struct
		structs = append(structs, &StructModel{
			SourceSchema: schema,
			ParentSchema: parentSchema,
		})

		// Process properties for nested objects
		props := parser.GetProperties(schema)
		for _, prop := range props {
			if parser.GetTypeOrEmpty(prop) == "object" {
				collectStructs(prop, schema)
			}
		}

		// Process array items for nested objects
		if parser.GetTypeOrEmpty(schema) == "array" {
			items := parser.GetArrayItems(schema)
			if items != nil && items.Schema != nil && parser.GetTypeOrEmpty(items.Schema) == "object" {
				collectStructs(items.Schema, schema)
			}
		}

		// Process oneOf/anyOf/allOf schemas
		if parser.HasOneOf(schema) {
			for _, oneOf := range *schema.OneOf {
				collectStructs(oneOf, schema)
			}
		}
		if parser.HasAnyOf(schema) {
			for _, anyOf := range *schema.AnyOf {
				collectStructs(anyOf, schema)
			}
		}
		if parser.HasAllOf(schema) {
			for _, allOf := range *schema.AllOf {
				collectStructs(allOf, schema)
			}
		}
	}

	// Start collection from root schema
	collectStructs(s.SourceSchema, nil)

	// Sort structs by name for consistent output
	sort.Slice(structs, func(i, j int) bool {
		return structs[i].Name() < structs[j].Name()
	})

	return structs
}

func (s *SchemaModel) Enums() []*EnumModel {
	var enums []*EnumModel
	seen := make(map[string]bool)

	// Helper function to collect enums from a schema
	var collectEnums func(schema *jsonschema.Schema, parentName string)
	collectEnums = func(schema *jsonschema.Schema, parentName string) {
		// Skip if nil schema
		if schema == nil {
			return
		}

		// Process properties for enums
		props := parser.GetProperties(schema)
		for propName, prop := range props {
			if enum := parser.GetEnum(prop); len(enum) > 0 {
				pp.Printf("🔍 Processing enum for property %s\n", propName)
				pp.Printf("📝 Schema type: %s\n", parser.GetTypeOrEmpty(prop))
				pp.Printf("🔢 Enum values: %+v\n", enum)

				// For enums, use property name as type name
				enumName := toGoFieldName(propName)
				if strings.HasSuffix(enumName, "Color") {
					enumName = "Color"
				}
				enumName = enumName + "Type"

				// Skip if already processed
				if seen[enumName] {
					continue
				}
				seen[enumName] = true

				// Determine base type from first enum value
				baseType := "string" // Default to string
				if len(enum) > 0 {
					pp.Printf("🎯 First enum value: %+v\n", enum[0])
					if enum[0].Bool != nil {
						baseType = "bool"
					} else if enum[0].String != nil {
						// Check if the string value represents an integer
						if _, err := strconv.Atoi(*enum[0].String); err == nil {
							baseType = "int"
						}
					} else if parser.GetTypeOrEmpty(prop) == "integer" {
						baseType = "int"
					}
					pp.Printf("📊 Determined base type: %s\n", baseType)
				}

				// Create enum model
				enumModel := &EnumModel{
					Name:        enumName,
					BaseType:    baseType,
					Description: parser.GetDescription(prop),
				}

				// Add enum values in the order they appear in the schema
				for _, val := range enum {
					pp.Printf("🔄 Processing enum value: %+v\n", val)
					var enumValue string
					if val.String != nil {
						enumValue = *val.String
					} else if val.Bool != nil {
						enumValue = fmt.Sprintf("%v", *val.Bool)
					} else if enumModel.BaseType == "int" {
						// For integer enums, try to parse the string value as an integer
						if val.String != nil {
							if intVal, err := strconv.Atoi(*val.String); err == nil {
								enumValue = fmt.Sprintf("%d", intVal)
							} else {
								continue // Skip invalid integer values
							}
						}
					} else {
						continue // Skip invalid values
					}

					valueName := fmt.Sprintf("%s%s", enumModel.Name, toTitleCase(enumValue))
					value := enumValue
					if enumModel.BaseType == "string" {
						value = fmt.Sprintf("%q", enumValue)
					} else if enumModel.BaseType == "int" {
						// For integer enums, ensure we have a valid integer
						if intVal, err := strconv.Atoi(enumValue); err == nil {
							value = fmt.Sprintf("%d", intVal)
						} else {
							continue // Skip invalid integer values
						}
					}
					pp.Printf("✨ Created enum value - Name: %s, Value: %s\n", valueName, value)
					enumModel.Values = append(enumModel.Values, &EnumValue{
						Name:   valueName,
						Value:  value,
						Parent: enumModel,
					})
				}

				enums = append(enums, enumModel)
			}

			// Recursively process nested objects
			if parser.GetTypeOrEmpty(prop) == "object" {
				collectEnums(prop, propName)
			}
		}

		// Process array items for enums
		if parser.GetTypeOrEmpty(schema) == "array" {
			items := parser.GetArrayItems(schema)
			if items != nil && items.Schema != nil {
				collectEnums(items.Schema, parentName+"Item")
			}
		}

		// Process oneOf/anyOf/allOf schemas
		if parser.HasOneOf(schema) {
			for _, oneOf := range *schema.OneOf {
				collectEnums(oneOf, parentName+"_OneOf")
			}
		}
		if parser.HasAnyOf(schema) {
			for _, anyOf := range *schema.AnyOf {
				collectEnums(anyOf, parentName+"_AnyOf")
			}
		}
		if parser.HasAllOf(schema) {
			for _, allOf := range *schema.AllOf {
				collectEnums(allOf, parentName+"_AllOf")
			}
		}
	}

	// Start collection from root schema
	collectEnums(s.SourceSchema, "")

	// Sort enums by name for consistent output
	sort.Slice(enums, func(i, j int) bool {
		return enums[i].Name < enums[j].Name
	})

	return enums
}

func (s *SchemaModel) Imports() []string {
	imports := make(map[string]bool)

	// Always include basic imports
	imports["encoding/json"] = true
	imports["gitlab.com/tozd/go/errors"] = true

	// Add imports based on types used
	for _, st := range s.Structs() {
		if st.HasValidation() {
			imports["gitlab.com/tozd/go/errors"] = true
		}
		// Add more imports based on needs...
	}

	// Convert map to sorted slice
	var result []string
	for imp := range imports {
		result = append(result, imp)
	}
	sort.Strings(result)

	return result
}
