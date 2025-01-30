package parser

import (
	"testing"

	"github.com/google/gnostic/jsonschema"
	"github.com/stretchr/testify/require"
	"github.com/walteh/schema2go/pkg/diff"
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

// TestGnosticTypes verifies our understanding of how gnostic represents different schema components
func TestGnosticTypes(t *testing.T) {
	// 🔍 Test Case 1: Type Field Variations
	t.Run("type_field_variations", func(t *testing.T) {
		input := `{
			"properties": {
				"single_type": {
					"type": "string"
				},
				"multiple_types": {
					"type": ["string", "null"]
				},
				"no_type": {
					"description": "field with no type"
				}
			}
		}`

		got, err := Parse(input)
		require.NoError(t, err, "parsing schema")

		want := &jsonschema.Schema{
			Properties: &[]*jsonschema.NamedSchema{
				{
					Name: "single_type",
					Value: &jsonschema.Schema{
						Type: &jsonschema.StringOrStringArray{
							String: Ptr("string"),
						},
					},
				},
				{
					Name: "multiple_types",
					Value: &jsonschema.Schema{
						Type: &jsonschema.StringOrStringArray{
							StringArray: &[]string{"string", "null"},
						},
					},
				},
				{
					Name: "no_type",
					Value: &jsonschema.Schema{
						Description: Ptr("field with no type"),
					},
				},
			},
		}

		diff.RequireKnownValueEqual(t, want, got)
	})

	// 🔄 Test Case 2: Array Field Variations
	t.Run("array_field_variations", func(t *testing.T) {
		input := `{
			"type": "object",
			"properties": {
				"simple_array": {
					"type": "array",
					"items": {
						"type": "string"
					}
				},
				"tuple_array": {
					"type": "array",
					"items": [
						{"type": "string"},
						{"type": "integer"}
					]
				},
				"nested_array": {
					"type": "array",
					"items": {
						"type": "array",
						"items": {
							"type": "string"
						}
					}
				}
			}
		}`

		got, err := Parse(input)
		require.NoError(t, err, "parsing schema")

		want := &jsonschema.Schema{
			Type: &jsonschema.StringOrStringArray{
				String: Ptr("object"),
			},
			Properties: &[]*jsonschema.NamedSchema{
				{
					Name: "simple_array",
					Value: &jsonschema.Schema{
						Type: &jsonschema.StringOrStringArray{
							String: Ptr("array"),
						},
						Items: &jsonschema.SchemaOrSchemaArray{
							Schema: &jsonschema.Schema{
								Type: &jsonschema.StringOrStringArray{
									String: Ptr("string"),
								},
							},
						},
					},
				},
				{
					Name: "tuple_array",
					Value: &jsonschema.Schema{
						Type: &jsonschema.StringOrStringArray{
							String: Ptr("array"),
						},
						Items: &jsonschema.SchemaOrSchemaArray{
							SchemaArray: &[]*jsonschema.Schema{
								{
									Type: &jsonschema.StringOrStringArray{
										String: Ptr("string"),
									},
								},
								{
									Type: &jsonschema.StringOrStringArray{
										String: Ptr("integer"),
									},
								},
							},
						},
					},
				},
				{
					Name: "nested_array",
					Value: &jsonschema.Schema{
						Type: &jsonschema.StringOrStringArray{
							String: Ptr("array"),
						},
						Items: &jsonschema.SchemaOrSchemaArray{
							Schema: &jsonschema.Schema{
								Type: &jsonschema.StringOrStringArray{
									String: Ptr("array"),
								},
								Items: &jsonschema.SchemaOrSchemaArray{
									Schema: &jsonschema.Schema{
										Type: &jsonschema.StringOrStringArray{
											String: Ptr("string"),
										},
									},
								},
							},
						},
					},
				},
			},
		}

		diff.RequireKnownValueEqual(t, want, got)
	})

	// 📚 Test Case 3: Validation Rules
	t.Run("validation_rules", func(t *testing.T) {
		input := `{
			"type": "object",
			"properties": {
				"string_field": {
					"type": "string",
					"minLength": 1,
					"maxLength": 100,
					"pattern": "^[a-z]+$"
				},
				"number_field": {
					"type": "number",
					"minimum": 0,
					"maximum": 100,
					"multipleOf": 0.5
				},
				"array_field": {
					"type": "array",
					"minItems": 1,
					"maxItems": 10,
					"uniqueItems": true
				}
			}
		}`

		got, err := Parse(input)
		require.NoError(t, err, "parsing schema")

		want := &jsonschema.Schema{
			Type: &jsonschema.StringOrStringArray{
				String: Ptr("object"),
			},
			Properties: &[]*jsonschema.NamedSchema{
				{
					Name: "string_field",
					Value: &jsonschema.Schema{
						Type: &jsonschema.StringOrStringArray{
							String: Ptr("string"),
						},
						MinLength: Ptr(int64(1)),
						MaxLength: Ptr(int64(100)),
						Pattern:   Ptr("^[a-z]+$"),
					},
				},
				{
					Name: "number_field",
					Value: &jsonschema.Schema{
						Type: &jsonschema.StringOrStringArray{
							String: Ptr("number"),
						},
						Minimum:    jsonschema.NewSchemaNumberWithInteger(0),
						Maximum:    jsonschema.NewSchemaNumberWithInteger(100),
						MultipleOf: jsonschema.NewSchemaNumberWithFloat(0.5),
					},
				},
				{
					Name: "array_field",
					Value: &jsonschema.Schema{
						Type: &jsonschema.StringOrStringArray{
							String: Ptr("array"),
						},
						MinItems:    Ptr(int64(1)),
						MaxItems:    Ptr(int64(10)),
						UniqueItems: Ptr(true),
					},
				},
			},
		}

		diff.RequireKnownValueEqual(t, want, got)
	})

	// 🎯 Test Case 4: References and Definitions
	t.Run("references_and_definitions", func(t *testing.T) {
		input := `{
			"type": "object",
			"definitions": {
				"address": {
					"type": "object",
					"properties": {
						"street": { "type": "string" },
						"city": { "type": "string" }
					}
				}
			},
			"properties": {
				"home": {
					"$ref": "#/definitions/address"
				},
				"work": {
					"$ref": "#/definitions/address"
				}
			}
		}`

		got, err := Parse(input)
		require.NoError(t, err, "parsing schema")

		want := &jsonschema.Schema{
			Type: &jsonschema.StringOrStringArray{String: Ptr("object")},
			Definitions: &[]*jsonschema.NamedSchema{
				{
					Name: "address",
					Value: &jsonschema.Schema{
						Type: &jsonschema.StringOrStringArray{String: Ptr("object")},
						Properties: &[]*jsonschema.NamedSchema{
							{
								Name:  "street",
								Value: NewType("string"),
							},
							{
								Name:  "city",
								Value: NewType("string"),
							},
						},
					},
				},
			},
			Properties: &[]*jsonschema.NamedSchema{
				{
					Name:  "home",
					Value: NewDefinitionRef("address"),
				},
				{
					Name:  "work",
					Value: NewDefinitionRef("address"),
				},
			},
		}

		diff.RequireKnownValueEqual(t, want, got)
	})

	// 🔄 Test Case 5: AllOf Variations
	t.Run("allof_variations", func(t *testing.T) {
		input := `{
			"type": "object",
			"allOf": [
				{
					"type": "object",
					"properties": {
						"name": { "type": "string" }
					},
					"required": ["name"]
				},
				{
					"type": "object",
					"properties": {
						"age": { "type": "integer" }
					},
					"required": ["age"]
				}
			]
		}`

		got, err := Parse(input)
		require.NoError(t, err, "parsing schema")

		want := &jsonschema.Schema{
			Type: &jsonschema.StringOrStringArray{
				String: Ptr("object"),
			},
			AllOf: &[]*jsonschema.Schema{
				{
					Type: &jsonschema.StringOrStringArray{
						String: Ptr("object"),
					},
					Properties: &[]*jsonschema.NamedSchema{
						{
							Name:  "name",
							Value: NewType("string"),
						},
					},
					Required: &[]string{"name"},
				},
				{
					Type: &jsonschema.StringOrStringArray{
						String: Ptr("object"),
					},
					Properties: &[]*jsonschema.NamedSchema{
						{
							Name:  "age",
							Value: NewType("integer"),
						},
					},
					Required: &[]string{"age"},
				},
			},
		}

		diff.RequireKnownValueEqual(t, want, got)
	})

	// 🔄 Test Case 6: AnyOf Variations
	t.Run("anyof_variations", func(t *testing.T) {
		input := `{
			"type": "object",
			"properties": {
				"value": {
					"anyOf": [
						{
							"type": "string",
							"minLength": 1
						},
						{
							"type": "integer",
							"minimum": 0
						},
						{
							"type": "object",
							"properties": {
								"custom": { "type": "string" }
							}
						}
					]
				}
			}
		}`

		got, err := Parse(input)
		require.NoError(t, err, "parsing schema")

		want := &jsonschema.Schema{
			Type: &jsonschema.StringOrStringArray{
				String: Ptr("object"),
			},
			Properties: &[]*jsonschema.NamedSchema{
				{
					Name: "value",
					Value: &jsonschema.Schema{
						AnyOf: &[]*jsonschema.Schema{
							{
								Type: &jsonschema.StringOrStringArray{
									String: Ptr("string"),
								},
								MinLength: Ptr(int64(1)),
							},
							{
								Type: &jsonschema.StringOrStringArray{
									String: Ptr("integer"),
								},
								Minimum: jsonschema.NewSchemaNumberWithInteger(0),
							},
							{
								Type: &jsonschema.StringOrStringArray{
									String: Ptr("object"),
								},
								Properties: &[]*jsonschema.NamedSchema{
									{
										Name:  "custom",
										Value: NewType("string"),
									},
								},
							},
						},
					},
				},
			},
		}

		diff.RequireKnownValueEqual(t, want, got)
	})

	// 🎯 Test Case 5: OneOf with Discriminator
	// Note: In gnostic's JSON Schema implementation, 'const' values are represented using
	// the Enumeration field with a single value. This is because gnostic doesn't have a
	// dedicated 'const' field in its Schema struct. For example:
	//   JSON Schema: { "const": "dog" }
	//   Gnostic: { "enumeration": [{ "string": "dog" }] }
	t.Run("oneof_with_discriminator", func(t *testing.T) {
		input := `{
			"type": "object",
			"oneOf": [
				{
					"type": "object",
					"properties": {
						"type": {
							"type": "string",
							"enum": ["dog"]
						},
						"bark": {
							"type": "string"
						}
					},
					"required": ["type", "bark"]
				},
				{
					"type": "object",
					"properties": {
						"type": {
							"type": "string",
							"enum": ["cat"]
						},
						"meow": {
							"type": "string"
						}
					},
					"required": ["type", "meow"]
				}
			]
		}`

		got, err := Parse(input)
		require.NoError(t, err, "parsing schema")

		want := &jsonschema.Schema{
			Type: &jsonschema.StringOrStringArray{String: Ptr("object")},
			OneOf: &[]*jsonschema.Schema{
				{
					Type: &jsonschema.StringOrStringArray{String: Ptr("object")},
					Properties: &[]*jsonschema.NamedSchema{
						{
							Name: "type",
							Value: &jsonschema.Schema{
								Type: &jsonschema.StringOrStringArray{String: Ptr("string")},
								Enumeration: &[]jsonschema.SchemaEnumValue{
									{String: Ptr("dog")},
								},
							},
						},
						{
							Name:  "bark",
							Value: NewType("string"),
						},
					},
					Required: &[]string{"type", "bark"},
				},
				{
					Type: &jsonschema.StringOrStringArray{String: Ptr("object")},
					Properties: &[]*jsonschema.NamedSchema{
						{
							Name: "type",
							Value: &jsonschema.Schema{
								Type: &jsonschema.StringOrStringArray{String: Ptr("string")},
								Enumeration: &[]jsonschema.SchemaEnumValue{
									{String: Ptr("cat")},
								},
							},
						},
						{
							Name:  "meow",
							Value: NewType("string"),
						},
					},
					Required: &[]string{"type", "meow"},
				},
			},
		}

		diff.RequireKnownValueEqual(t, want, got)
	})

	// 🔄 Test Case 8: Nested Combinations
	t.Run("nested_combinations", func(t *testing.T) {
		input := `{
			"type": "object",
			"properties": {
				"mixed": {
					"anyOf": [
						{
							"type": "object",
							"allOf": [
								{
									"type": "object",
									"properties": {
										"id": { "type": "string" }
									}
								},
								{
									"type": "object",
									"properties": {
										"value": { "type": "integer" }
									}
								}
							]
						},
						{
							"type": "object",
							"oneOf": [
								{
									"type": "object",
									"properties": {
										"name": { "type": "string" }
									}
								},
								{
									"type": "object",
									"properties": {
										"code": { "type": "integer" }
									}
								}
							]
						}
					]
				}
			}
		}`

		got, err := Parse(input)
		require.NoError(t, err, "parsing schema")

		want := &jsonschema.Schema{
			Type: &jsonschema.StringOrStringArray{
				String: Ptr("object"),
			},
			Properties: &[]*jsonschema.NamedSchema{
				{
					Name: "mixed",
					Value: &jsonschema.Schema{
						AnyOf: &[]*jsonschema.Schema{
							{
								Type: &jsonschema.StringOrStringArray{
									String: Ptr("object"),
								},
								AllOf: &[]*jsonschema.Schema{
									{
										Type: &jsonschema.StringOrStringArray{
											String: Ptr("object"),
										},
										Properties: &[]*jsonschema.NamedSchema{
											{
												Name:  "id",
												Value: NewType("string"),
											},
										},
									},
									{
										Type: &jsonschema.StringOrStringArray{
											String: Ptr("object"),
										},
										Properties: &[]*jsonschema.NamedSchema{
											{
												Name:  "value",
												Value: NewType("integer"),
											},
										},
									},
								},
							},
							{
								Type: &jsonschema.StringOrStringArray{
									String: Ptr("object"),
								},
								OneOf: &[]*jsonschema.Schema{
									{
										Type: &jsonschema.StringOrStringArray{
											String: Ptr("object"),
										},
										Properties: &[]*jsonschema.NamedSchema{
											{
												Name:  "name",
												Value: NewType("string"),
											},
										},
									},
									{
										Type: &jsonschema.StringOrStringArray{
											String: Ptr("object"),
										},
										Properties: &[]*jsonschema.NamedSchema{
											{
												Name:  "code",
												Value: NewType("integer"),
											},
										},
									},
								},
							},
						},
					},
				},
			},
		}

		diff.RequireKnownValueEqual(t, want, got)
	})
}
