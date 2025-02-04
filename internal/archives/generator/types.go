package generator

import (
	"encoding/json"
	"fmt"
	"regexp"
	"slices"
	"sort"
	"strconv"
	"strings"
	"unicode"

	"github.com/google/gnostic/jsonschema"
	"github.com/k0kubun/pp/v3"
	"github.com/walteh/schema2go/pkg/parser"
)

type Schema interface {
	Package() string
	Structs() []Struct
	Enums() []EnumModel
	Imports() []string
}

// SchemaModel represents a parsed JSON Schema ready for code generation
type SchemaModel struct {
	SourceSchema *jsonschema.Schema
}

func (s *SchemaModel) RemoveYamlLineNumbers() {
	var err error
	s.SourceSchema, err = removeYamlLineNumbers(s.SourceSchema)
	if err != nil {
		panic(err)
	}
}

var lregex = regexp.MustCompile(`"Line":\d+`)
var cregex = regexp.MustCompile(`"Column":\d+`)

func removeYamlLineNumbers(schema *jsonschema.Schema) (*jsonschema.Schema, error) {
	// marshal it to json , set all the line and column numbers to zero and unmarshal it back
	jsonData, err := json.Marshal(schema)
	if err != nil {
		return nil, err
	}

	jsonData = lregex.ReplaceAll(jsonData, []byte("\"Line\":0"))
	jsonData = cregex.ReplaceAll(jsonData, []byte("\"Column\":0"))

	var unmarshalled jsonschema.Schema
	if err := json.Unmarshal(jsonData, &unmarshalled); err != nil {
		return nil, err
	}

	return &unmarshalled, nil
}

type Struct interface {
	Name() string
	Description() string
	Fields() []Field
	HasAllOf() bool
	HasCustomMarshaling() bool
	HasDefaults() bool
	HasValidation() bool
}

// StructModel represents a Go struct to generate
type StructModel struct {
	SourceSchema *jsonschema.Schema
	ParentSchema *jsonschema.Schema // Parent schema that contains this schema
}

type Field interface {
	Name() string
	JSONName() string
	Description() string
	IsRequired() bool
	Type() string
	IsEnum() bool
	EnumTypeName() string
	EnumValues() []EnumValue
	DefaultValue() *string
	DefaultValueComment() *string
	ValidationRules() []ValidationRule
	BaseType() string
}

var _ Field = &FieldModel{}

// FieldModel represents a Go struct field
// JSON tags are not added to the main struct fields since we use custom marshaling.
// Instead, JSON tags are added to an internal alias struct in the UnmarshalJSON method.
// This approach gives us more control over the marshaling process and allows us to:
// 1. Handle validation before/after marshaling
// 2. Apply default values during unmarshaling
// 3. Keep the main struct clean and focused on the Go representation
type FieldModel struct {
	SourceSchema   *jsonschema.Schema
	ParentSchema   *jsonschema.Schema
	customJSONTag  string // Used for internal JSON tag generation in marshal/unmarshal methods
	customGoName   string // Optional override for the Go field name (e.g., Name_AllOf)
	Required       bool
	Parent         *StructModel
	CustomJSONName string
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
	// Parent  Field
	Values string
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
	Values      []EnumValue
}

func (e *EnumModel) AddValue(name, value, description string) {
	e.Values = append(e.Values, EnumValue{
		Name:        name,
		Value:       value,
		Description: description,
		Parent:      e,
	})
}

// FieldModel Methods

func (f *FieldModel) Name() string {
	if f.customGoName != "" {
		return f.customGoName
	}
	return toGoFieldName(f.JSONName())
}

func (f *FieldModel) JSONName() string {
	// We don't need JSON tags on the main struct since we have custom marshaling
	return ""
}

// This is the internal JSON name used for marshaling/unmarshaling
func (f *FieldModel) internalJSONName() string {
	if f.Parent == nil {
		return ""
	}

	// For regular fields
	name := f.customJSONTag
	if name == "" {
		name = toJSONFieldName(f.Name())
	}

	// Only add omitempty for non-required fields if it's not already there
	if !f.IsRequired() && !strings.Contains(name, "omitempty") {
		name += ",omitempty"
	}

	return name
}

func (f *FieldModel) Description() string {
	return parser.GetDescription(f.SourceSchema)
}

func (f *FieldModel) IsRequired() bool {
	if f.Parent == nil {
		return false
	}

	if f.Parent.HasAllOf() && f.IsEmbedded() {
		return false
	}

	return f.Required
}

func (f *FieldModel) Type() string {
	if f.Parent == nil {
		return ""
	}

	baseType := f.BaseType()
	if !f.IsRequired() && !strings.HasPrefix(baseType, "*") {
		return "*" + baseType
	}
	return baseType
}

func (f *FieldModel) IsEnum() bool {
	return len(parser.GetEnum(f.SourceSchema)) > 0
}

func (f *FieldModel) EnumTypeName() string {
	if !f.IsEnum() {
		return ""
	}
	enumName := toGoFieldName(f.JSONName())
	if strings.HasSuffix(enumName, "Color") {
		enumName = "Color"
	}
	return enumName + "Type"
}

func (f *FieldModel) EnumValues() []EnumValue {
	if !f.IsEnum() {
		return nil
	}

	enum := parser.GetEnum(f.SourceSchema)
	values := make([]EnumValue, 0, len(enum))

	for _, val := range enum {
		if val.String == nil {
			continue
		}
		values = append(values, EnumValue{
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

func (f *FieldModel) ValidationRules() []ValidationRule {
	var rules []ValidationRule

	// Skip validation for allOf fields
	if strings.HasSuffix(f.customGoName, "_AllOf") {
		return nil
	}

	// Required validation
	if f.IsRequired() && parser.GetTypeOrEmpty(f.SourceSchema) != "object" {
		rules = append(rules, ValidationRule{
			Type:    ValidationRequired,
			Message: fmt.Sprintf("%s is required", f.JSONName()),
		})
	}

	// Enum validation
	if f.IsEnum() {
		values := make([]string, 0, len(f.EnumValues()))
		for _, v := range f.EnumValues() {
			values = append(values, v.Value)
		}
		rules = append(rules, ValidationRule{
			Type:    ValidationEnum,
			Message: fmt.Sprintf("invalid %s", f.JSONName()),
			Values:  strings.Join(values, ", "),
		})
	}

	// Nested validation
	if parser.GetTypeOrEmpty(f.SourceSchema) == "object" {
		rules = append(rules, ValidationRule{
			Type:    ValidationNested,
			Message: fmt.Sprintf("validating %s", f.JSONName()),
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

	// For nested objects, use parent title + field name
	if s.ParentSchema != nil {
		if parentTitle := parser.GetTitle(s.ParentSchema); parentTitle != "" {
			props := parser.GetProperties(s.ParentSchema)
			for name, prop := range props {
				if prop == s.SourceSchema {
					return parentTitle + toGoFieldName(name)
				}
			}
		}
	}

	// Fallback to "Object" if no name can be determined
	return "Object"
}

func (s *StructModel) Description() string {
	return parser.GetDescription(s.SourceSchema)
}

func (s *StructModel) Fields() []Field {
	var fields []Field

	// Handle allOf fields
	if s.HasAllOf() {
		for _, allOf := range parser.GetAllOf(s.SourceSchema) {
			if allOf.Ref != nil {
				// For referenced types
				typeName := strings.Split(*allOf.Ref, "/")[len(strings.Split(*allOf.Ref, "/"))-1]
				field := &FieldModel{
					SourceSchema:  allOf,
					ParentSchema:  s.SourceSchema,
					customJSONTag: "", // Empty for embedded fields
					customGoName:  typeName + "_AllOf",
					Required:      false, // Embedded fields are never required
					Parent:        s,
				}
				fields = append(fields, field)
			} else {
				// For inline types
				props := parser.GetProperties(allOf)
				for name, schema := range props {
					field := &FieldModel{
						SourceSchema:  schema,
						ParentSchema:  allOf,
						customJSONTag: "", // Empty for embedded fields
						customGoName:  toGoFieldName(name) + "_AllOf",
						Required:      false, // Embedded fields are never required
						Parent:        s,
					}
					fields = append(fields, field)
				}
			}
		}
	}

	// Handle regular fields
	props := parser.GetProperties(s.SourceSchema)
	for name, schema := range props {
		field := &FieldModel{
			SourceSchema:  schema,
			ParentSchema:  s.SourceSchema,
			customJSONTag: name,
			customGoName:  toGoFieldName(name),
			Required:      parser.IsRequired(s.SourceSchema, name),
			Parent:        s,
		}
		fields = append(fields, field)
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

func (s *StructModel) HasAllOf() bool {
	return parser.HasAllOf(s.SourceSchema)
}

func (s *StructModel) HasCustomMarshaling() bool {
	// Return true if:
	// 1. The struct has allOf fields
	// 2. The struct has validation rules
	// 3. The struct has default values
	return s.HasAllOf() || s.HasValidation() || s.HasDefaults()
}

// SchemaModel Methods

func (s *SchemaModel) Package() string {
	return "models" // TODO: Make configurable
}

func (s *SchemaModel) Structs() []Struct {
	var structs []Struct
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
						// For nested objects, use parent title + property name
						name = parentTitle + toGoFieldName(propName)
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

		// Handle allOf schemas
		if parser.HasAllOf(schema) {
			// Add this struct
			structs = append(structs, &StructModel{
				SourceSchema: schema,
				ParentSchema: parentSchema,
			})

			// Process allOf references to generate their structs
			allOf := *schema.AllOf
			for _, ref := range allOf {
				if ref.Ref != nil {
					// Get the referenced schema
					if refSchema := parser.GetDefinition(s.SourceSchema, *ref.Ref); refSchema != nil {
						// Extract the type name from the reference
						refParts := strings.Split(*ref.Ref, "/")
						typeName := refParts[len(refParts)-1]

						// Create a new struct for the referenced type with the correct name
						refStruct := &StructModel{
							SourceSchema: refSchema,
							ParentSchema: schema,
						}
						// Set the title to ensure correct naming
						refSchema.Title = parser.Ptr(typeName)
						structs = append(structs, refStruct)
					}
				}
			}
			return
		}

		// Add this struct
		structs = append(structs, &StructModel{
			SourceSchema: schema,
			ParentSchema: parentSchema,
		})

		// Process properties for nested objects and references
		props := parser.GetProperties(schema)
		for propName, prop := range props {
			// Handle referenced types
			if prop.Ref != nil {
				if refSchema := parser.GetDefinition(s.SourceSchema, *prop.Ref); refSchema != nil {
					// Extract the type name from the reference
					refParts := strings.Split(*prop.Ref, "/")
					typeName := refParts[len(refParts)-1]

					// Create a new struct for the referenced type
					refStruct := &StructModel{
						SourceSchema: refSchema,
						ParentSchema: schema,
					}
					// Set the title to ensure correct naming
					refSchema.Title = parser.Ptr(typeName)
					structs = append(structs, refStruct)
				}
				continue
			}

			// Handle nested objects
			if parser.GetTypeOrEmpty(prop) == "object" {
				// Set the title for the nested object to ensure correct naming
				if title := parser.GetTitle(schema); title != "" {
					prop.Title = parser.Ptr(title + toGoFieldName(propName))
				}
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
	slices.SortFunc(structs, func(a, b Struct) int {
		return strings.Compare(a.Name(), b.Name())
	})

	return structs
}

func (s *SchemaModel) Enums() []EnumModel {
	var enums []EnumModel
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
				enumModel := EnumModel{
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
					enumModel.Values = append(enumModel.Values, EnumValue{
						Name:   valueName,
						Value:  value,
						Parent: &enumModel,
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
	slices.SortFunc(enums, func(a, b EnumModel) int {
		return strings.Compare(a.Name, b.Name)
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

func toJSONFieldName(name string) string {
	// Special case for ID fields
	if name == "ID" {
		return "id"
	}

	// If the name is already snake_case, return it as is
	if strings.Contains(name, "_") {
		return name
	}

	// Convert camelCase to snake_case
	var result []rune
	for i, r := range name {
		if i > 0 && unicode.IsUpper(r) {
			result = append(result, '_')
		}
		result = append(result, unicode.ToLower(r))
	}
	return string(result)
}

func (f *FieldModel) IsEmbedded() bool {
	return strings.HasSuffix(f.customGoName, "_AllOf")
}

func (f *FieldModel) BaseType() string {
	// For referenced types
	if f.SourceSchema.Ref != nil {
		// Extract the type name from the reference
		refParts := strings.Split(*f.SourceSchema.Ref, "/")
		return refParts[len(refParts)-1]
	}

	// For allOf fields, use the primitive type
	if f.IsEmbedded() {
		baseType := parser.GetTypeOrEmpty(f.SourceSchema)
		switch baseType {
		case "string":
			return "string"
		case "integer":
			return "int"
		case "number":
			return "float64"
		case "boolean":
			return "bool"
		default:
			return strings.TrimSuffix(f.customGoName, "_AllOf")
		}
	}

	if f.IsEnum() {
		return f.EnumTypeName()
	}

	// Get base type
	baseType := parser.GetTypeOrEmpty(f.SourceSchema)

	// Map JSON Schema types to Go types
	switch baseType {
	case "string":
		return "string"
	case "integer":
		return "int"
	case "number":
		return "float64"
	case "boolean":
		return "bool"
	case "array":
		items := parser.GetArrayItems(f.SourceSchema)
		if items == nil || items.Schema == nil {
			return "[]interface{}" // Fallback for unknown array types
		}
		itemField := &FieldModel{
			SourceSchema: items.Schema,
			ParentSchema: f.SourceSchema,
		}
		return "[]" + itemField.Type()
	case "object":
		// For nested objects, use the parent struct name + field name
		if parentTitle := parser.GetTitle(f.ParentSchema); parentTitle != "" {
			return parentTitle + toGoFieldName(f.JSONName())
		}
		return toGoFieldName(f.JSONName())
	default:
		return "interface{}" // Fallback for unknown types
	}
}
