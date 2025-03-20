package codegen

import (
	"context"

	"github.com/google/gnostic/jsonschema"
)

func GenerateWithFormatting(ctx context.Context, schema *jsonschema.Schema) (string, error) {

	return "package example\n\n", nil

}

func GenerateNoFormatting(ctx context.Context, schema *jsonschema.Schema) (string, error) {

	return "package example\n\n", nil

}
