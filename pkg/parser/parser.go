package parser

import (
	"github.com/google/gnostic/jsonschema"
)

func GetAllOf(schema *jsonschema.Schema) []*jsonschema.Schema {
	if schema.AllOf == nil {
		return nil
	}
	return *schema.AllOf
}
