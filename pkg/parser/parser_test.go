package parser

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestParseBasicSchema(t *testing.T) {
	input := `{
		"title": "Person",
		"type": "object",
		"properties": {
			"name": {
				"type": "string"
			},
			"age": {
				"type": "integer"
			}
		},
		"required": ["name"]
	}`

	// Parse schema
	schema, err := Parse(input)
	require.NoError(t, err, "parsing schema")

	// Check basic properties
	require.Equal(t, "Person", GetTitle(schema), "schema title")
	require.Equal(t, "object", GetTypeOrEmpty(schema), "schema type")

	// Check properties
	props := GetProperties(schema)
	require.Len(t, props, 2, "number of properties")

	// Check name property
	nameProp := props["name"]
	require.NotNil(t, nameProp, "name property exists")
	require.Equal(t, "string", GetTypeOrEmpty(nameProp), "name property type")
	require.True(t, IsRequired(schema, "name"), "name property required")

	// Check age property
	ageProp := props["age"]
	require.NotNil(t, ageProp, "age property exists")
	require.Equal(t, "integer", GetTypeOrEmpty(ageProp), "age property type")
	require.False(t, IsRequired(schema, "age"), "age property not required")
}

func TestParseOneOf(t *testing.T) {
	input := `{
		"title": "Status",
		"oneOf": [
			{
				"type": "string",
				"enum": ["active", "inactive"]
			},
			{
				"type": "integer",
				"minimum": 0,
				"maximum": 100
			}
		]
	}`

	// Parse schema
	schema, err := Parse(input)
	require.NoError(t, err, "parsing schema")

	// Check basic properties
	require.Equal(t, "Status", GetTitle(schema), "schema title")
	require.True(t, HasOneOf(schema), "has oneOf")

	// Get oneOf schemas
	oneOf := *schema.OneOf
	require.Len(t, oneOf, 2, "number of oneOf schemas")

	// Check string enum
	strSchema := oneOf[0]
	require.Equal(t, "string", GetTypeOrEmpty(strSchema), "first oneOf type")

	// Check integer range
	intSchema := oneOf[1]
	require.Equal(t, "integer", GetTypeOrEmpty(intSchema), "second oneOf type")
}

func TestParseArray(t *testing.T) {
	input := `{
		"title": "Tags",
		"type": "array",
		"items": {
			"type": "string",
			"minLength": 1
		}
	}`

	// Parse schema
	schema, err := Parse(input)
	require.NoError(t, err, "parsing schema")

	// Check basic properties
	require.Equal(t, "Tags", GetTitle(schema), "schema title")
	require.True(t, IsArray(schema), "is array type")

	// Check items
	items := GetArrayItems(schema)
	require.NotNil(t, items, "items exists")
	require.Nil(t, items.SchemaArray, "items array should not exist")
	require.NotNil(t, items.Schema, "items schema should exist")
}
