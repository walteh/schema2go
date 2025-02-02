package parser

import (
	"strings"

	"github.com/atombender/go-jsonschema/pkg/schemas"
)

func ParseGoJsonSchema(input string) (*schemas.Schema, error) {
	return schemas.FromJSONReader(strings.NewReader(input))
}
