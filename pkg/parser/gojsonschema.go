package parser

import (
	"strings"

	"github.com/walteh/schema2go/pkg/schemas"
)

func ParseGoJsonSchema(input string) (*schemas.Schema, error) {
	return schemas.FromJSONReader(strings.NewReader(input))
}
