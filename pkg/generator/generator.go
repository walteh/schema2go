package generator

import (
	"context"
)

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

// New creates a new Generator with the given options
func New(opts Options) Generator {
	// TODO: Implement the actual generator
	return nil
}
