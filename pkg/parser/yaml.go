package parser

import (
	"encoding/json"
	"regexp"

	"github.com/google/gnostic/jsonschema"
)

func RemoveYamlLineNumbers(schema *jsonschema.Schema) (*jsonschema.Schema, error) {
	var err error
	schema, err = removeYamlLineNumbers(schema)
	if err != nil {
		return nil, err
	}
	return schema, nil
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
