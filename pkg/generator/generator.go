package generator

import (
	"context"
	_ "embed"
	"go/format"
	"strings"
	"text/template"

	pp "github.com/k0kubun/pp/v3"
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
}

// generator implements the Generator interface
type generator struct {
	opts     Options
	template *template.Template
}

// New creates a new Generator with the given options
func New(opts Options) Generator {
	pp.Println("New generator")
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
func GenerateWithFormatting(ctx context.Context, g Generator, input string) (string, error) {
	code, err := g.Generate(ctx, input)
	if err != nil {
		return "", errors.Errorf("generating code: %w", err)
	}

	// format the go code
	formatted, err := format.Source([]byte(code))
	if err != nil {
		return "", errors.Errorf("formatting code: %w", err)
	}

	return string(formatted), nil
}

func (g *generator) Generate(ctx context.Context, input string) (string, error) {
	// 🔍 Parse the schema
	schema, err := parser.Parse(input)
	if err != nil {
		return "", errors.Errorf("parsing schema: %w", err)
	}

	// 🏗️ Create our model
	model := &SchemaModel{
		SourceSchema: schema,
	}

	// 📝 Execute template
	var b strings.Builder
	if err := g.template.Execute(&b, model); err != nil {
		return "", errors.Errorf("executing template: %w", err)
	}

	return b.String(), nil
}
