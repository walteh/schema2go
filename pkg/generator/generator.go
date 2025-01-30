package generator

import (
	"context"
	_ "embed"
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

	return b.String(), nil
}

// buildSchemaModel converts a gnostic schema to our template model
func (g *generator) buildSchemaModel(schema *jsonschema.Schema) (*SchemaModel, error) {
	model := &SchemaModel{
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
	model.Structs = append(model.Structs, structModel)

	return model, nil
}

// buildFieldModel converts a gnostic schema property to our field model
func (g *generator) buildFieldModel(name string, schema *jsonschema.Schema, parent *jsonschema.Schema) (*FieldModel, error) {
	// Convert type
	goType, err := g.schemaToGoType(schema)
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

	// Add validation rules
	if field.IsRequired {
		field.ValidationRules = append(field.ValidationRules, &ValidationRule{
			Type:    ValidationRequired,
			Message: name + " is required",
		})
	}

	// Handle optional fields
	if !field.IsRequired {
		field.Type = "*" + field.Type
	}

	// Handle default value
	if schema.Default != nil {
		field.DefaultValue = schema.Default.Value
	}

	return field, nil
}

// schemaToGoType converts a JSON Schema type to a Go type
func (g *generator) schemaToGoType(schema *jsonschema.Schema) (string, error) {
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
			return "", errors.New("array must have items")
		}
		itemType, err := g.schemaToGoType(items.Schema)
		if err != nil {
			return "", errors.Errorf("getting array item type: %w", err)
		}
		return "[]" + itemType, nil
	case "object":
		// TODO: Handle nested objects
		return "", errors.New("nested objects not yet supported")
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
