package generator

import (
	"context"
	_ "embed"
	"fmt"
	"go/format"
	"sort"
	"strings"
	"text/template"

	"github.com/google/gnostic/jsonschema"
	"github.com/walteh/schema2go/pkg/parser"
	"gitlab.com/tozd/go/errors"
)

//go:embed templates/schema.go.tmpl
var schemaTemplate string

// Generator is responsible for converting JSON Schema into Go code
type Generator interface {
	// Generate takes a JSON schema as input and returns the generated Go code
	Generate(ctx context.Context, input string) (string, error)
}

// Options configures the generator behavior
type Options struct {
	// PackageName is the name of the generated Go package
	PackageName string

	// TODO: Add more options as needed:
	// - Custom type mappings
	// - Import path handling
	// - Validation options
	// - etc.
}

// generator implements the Generator interface
type generator struct {
	opts     Options
	template *template.Template
	model    *SchemaModel // Current model being built
}

// New creates a new Generator with the given options
func New(opts Options) Generator {
	// Create template with helper functions
	tmpl := template.New("schema").Funcs(template.FuncMap{
		"ToLower": strings.ToLower,
	})

	// Parse template
	tmpl = template.Must(tmpl.Parse(schemaTemplate))

	return &generator{
		opts:     opts,
		template: tmpl,
	}
}

// Generate implements the Generator interface
func (g *generator) Generate(ctx context.Context, input string) (string, error) {
	// 🔍 Parse the schema
	schema, err := parser.Parse(input)
	if err != nil {
		return "", errors.Errorf("parsing schema: %w", err)
	}

	// 🏗️ Convert schema to our template model
	model, err := g.buildSchemaModel(schema)
	if err != nil {
		return "", errors.Errorf("building schema model: %w", err)
	}

	// 📝 Execute template
	var b strings.Builder
	if err := g.template.Execute(&b, model); err != nil {
		return "", errors.Errorf("executing template: %w", err)
	}

	// format the go code
	formatted, err := format.Source([]byte(b.String()))
	if err != nil {
		return "", errors.Errorf("formatting code: %w", err)
	}

	return string(formatted), nil
}

// buildSchemaModel converts a gnostic schema to our template model
func (g *generator) buildSchemaModel(schema *jsonschema.Schema) (*SchemaModel, error) {
	g.model = &SchemaModel{
		Package: g.opts.PackageName,
		Imports: []string{
			"encoding/json",
			"gitlab.com/tozd/go/errors",
		},
	}

	// Get schema title (struct name)
	title := parser.GetTitle(schema)
	if title == "" {
		return nil, errors.New("schema must have a title")
	}

	// Create struct model
	structModel := &StructModel{
		Name:        title,
		Description: parser.GetDescription(schema),
	}

	// Add fields
	props := parser.GetProperties(schema)
	var fieldNames []string
	for name := range props {
		fieldNames = append(fieldNames, name)
	}
	sort.Strings(fieldNames)

	for _, name := range fieldNames {
		prop := props[name]
		field, err := g.buildFieldModel(name, prop, schema)
		if err != nil {
			return nil, errors.Errorf("building field model for %q: %w", name, err)
		}
		structModel.Fields = append(structModel.Fields, field)

		// Update struct flags based on field
		if field.IsRequired {
			structModel.HasValidation = true
		}
		if field.DefaultValue != nil {
			structModel.HasDefaults = true
		}
		if len(field.ValidationRules) > 0 {
			structModel.HasValidation = true
		}

		// Check if field has enum values
		if enum := parser.GetEnum(prop); enum != nil {
			enumModel := &EnumModel{
				Name:     toGoFieldName(name) + "Type",
				BaseType: field.Type,
			}
			for _, val := range enum {
				enumModel.Values = append(enumModel.Values, &EnumValue{
					Name:  toGoFieldName(fmt.Sprintf("%s_%s", name, *val.String)),
					Value: *val.String,
				})
			}
			g.model.Enums = append(g.model.Enums, enumModel)
			field.Type = enumModel.Name // Update field type to use enum type
		}
	}

	// Sort fields by required first, then alphabetically
	sort.SliceStable(structModel.Fields, func(i, j int) bool {
		// Required fields come first
		if structModel.Fields[i].IsRequired != structModel.Fields[j].IsRequired {
			return structModel.Fields[i].IsRequired
		}
		// Then sort alphabetically
		return structModel.Fields[i].Name < structModel.Fields[j].Name
	})

	// Add struct to model
	g.model.Structs = append(g.model.Structs, structModel)

	return g.model, nil
}

// buildFieldModel converts a gnostic schema property to our field model
func (g *generator) buildFieldModel(name string, schema *jsonschema.Schema, parent *jsonschema.Schema) (*FieldModel, error) {
	// Convert type
	goType, err := g.schemaToGoType(schema, "", name)
	if err != nil {
		return nil, errors.Errorf("converting type: %w", err)
	}

	// Create field model
	field := &FieldModel{
		Name:        toGoFieldName(name),
		Type:        goType,
		JSONName:    name,
		Description: parser.GetDescription(schema),
		IsRequired:  parser.IsRequired(parent, name),
	}

	// Handle validation rules
	if field.IsRequired {
		// For nested objects, we only want to validate the nested object
		if parser.GetTypeOrEmpty(schema) != "object" {
			field.ValidationRules = append(field.ValidationRules, &ValidationRule{
				Type:    ValidationRequired,
				Message: name + " is required",
				Field:   field,
			})
		}
	}

	// Handle optional fields
	if !field.IsRequired {
		field.Type = "*" + field.Type
	}

	// Handle default value
	if schema.Default != nil {
		field.DefaultValue = schema.Default.Value
	}

	// Handle nested object validation
	if parser.GetTypeOrEmpty(schema) == "object" {
		field.ValidationRules = append(field.ValidationRules, &ValidationRule{
			Type:    ValidationNested,
			Message: fmt.Sprintf("validating %s", name),
			Field:   field,
		})
	}

	return field, nil
}

// schemaToGoType converts a JSON Schema type to a Go type
func (g *generator) schemaToGoType(schema *jsonschema.Schema, parentName, fieldName string) (string, error) {
	typ := parser.GetTypeOrEmpty(schema)
	switch typ {
	case "string":
		return "string", nil
	case "integer":
		return "int", nil
	case "number":
		return "float64", nil
	case "boolean":
		return "bool", nil
	case "array":
		items := parser.GetArrayItems(schema)
		if items == nil || items.Schema == nil {
			return "", errors.New("array items not found")
		}
		itemType, err := g.schemaToGoType(items.Schema, parentName, fieldName)
		if err != nil {
			return "", errors.Errorf("getting array item type: %w", err)
		}
		return "[]" + itemType, nil
	case "object":
		// For nested objects, we'll use the field name as the type name
		if fieldName == "" {
			return "", errors.New("field name is required for nested objects")
		}

		// Create a new struct model for the nested type
		structName := toGoFieldName(fieldName)
		structModel := &StructModel{
			Name:        structName,
			Description: parser.GetDescription(schema),
		}

		// Get properties
		props := parser.GetProperties(schema)
		for name, prop := range props {
			field, err := g.buildFieldModel(name, prop, schema)
			if err != nil {
				return "", errors.Errorf("building field model for %s: %w", name, err)
			}
			structModel.Fields = append(structModel.Fields, field)

			// Update struct flags based on field
			if field.IsRequired {
				structModel.HasValidation = true
			}
			if field.DefaultValue != nil {
				structModel.HasDefaults = true
			}
		}

		// Sort fields by required first, then alphabetically
		sort.SliceStable(structModel.Fields, func(i, j int) bool {
			// Required fields come first
			if structModel.Fields[i].IsRequired != structModel.Fields[j].IsRequired {
				return structModel.Fields[i].IsRequired
			}
			// Then sort alphabetically
			return structModel.Fields[i].Name < structModel.Fields[j].Name
		})

		// Add the struct model to the list
		g.model.Structs = append(g.model.Structs, structModel)

		return structName, nil
	default:
		return "", errors.Errorf("unsupported type: %s", typ)
	}
}

// toGoFieldName converts a JSON field name to a Go field name
func toGoFieldName(name string) string {
	// Special case for "id" -> "ID"
	if strings.ToLower(name) == "id" {
		return "ID"
	}

	// Split on underscores and capitalize each word
	words := strings.Split(name, "_")
	for i, word := range words {
		if word == "" {
			continue
		}
		words[i] = strings.Title(word)
	}
	return strings.Join(words, "")
}
