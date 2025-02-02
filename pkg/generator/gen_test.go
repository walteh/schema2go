package generator_test

import (
	"github.com/google/gnostic/jsonschema"
	"github.com/stretchr/testify/require"
	"github.com/walteh/schema2go/pkg/diff"
	"github.com/walteh/schema2go/pkg/generator"
	"github.com/walteh/schema2go/pkg/generator/testcases"
	"gopkg.in/yaml.v3"
	"testing"
)

const testCasesHash = "6fa60fb8fa188ada872d2141f39f84cce11f12546e62be56feda50f028370c79"

func TestAllOfSchemaToStruct(t *testing.T) {

	tc := testcases.LoadAndParseTestCase("TestAllOfSchemaToStruct")
	require.Equal(t, testCasesHash, testcases.GetHash(), "test cases hash mismatch, please run 'go generate ./...' to update the test cases hash")

	schema, err := generator.NewSchemaModel(tc.JSONSchema())
	require.NoError(t, err, "failed to parse schema")

	t.Run("json-schema", func(t *testing.T) {
		// nothing to do here right now
	})

	t.Run("raw-schema", func(t *testing.T) {

		want := &generator.SchemaModel{
			SourceSchema: &jsonschema.Schema{
				Schema: ptr("http://json-schema.org/draft-07/schema#"),
				Title:  ptr("AllOfExample"),
				Type:   typ("object"),
				AllOf: ptr([]*jsonschema.Schema{
					{
						Type: typ("object"),
						Properties: &[]*jsonschema.NamedSchema{
							{Name: "name", Value: &jsonschema.Schema{Type: typ("string")}},
						},
					},
					{
						Type: typ("object"),
						Properties: &[]*jsonschema.NamedSchema{
							{Name: "age", Value: &jsonschema.Schema{Type: typ("integer")}},
						},
					},
				}),
			},
		}

		diff.RequireKnownValueEqual(t, want.SourceSchema, schema.SourceSchema)
	})

	t.Run("static-schema", func(t *testing.T) {

		f1 := &generator.StaticField{Name_: "Name_AllOf", JSONName_: "name", Type_: "string"}
		f2 := &generator.StaticField{Name_: "Age_AllOf", JSONName_: "age", Type_: "int"}
		s1 := &generator.StaticStruct{Name_: "AllOfExample", Fields_: []generator.Field{f1, f2}}

		staticWant := &generator.StaticSchema{
			Package_: "models",
			Structs_: []generator.Struct{
				s1,
			},
		}

		staticGot := generator.NewStaticSchema(schema)

		diff.RequireKnownValueEqual(t, staticWant, staticGot)
	})

	t.Run("go-code", func(t *testing.T) {

		checkGoCode(t, schema, tc.GoCode())
	})

}

func TestAllOfWithRefsSchemaToStruct(t *testing.T) {

	tc := testcases.LoadAndParseTestCase("TestAllOfWithRefsSchemaToStruct")
	require.Equal(t, testCasesHash, testcases.GetHash(), "test cases hash mismatch, please run 'go generate ./...' to update the test cases hash")

	schema, err := generator.NewSchemaModel(tc.JSONSchema())
	require.NoError(t, err, "failed to parse schema")

	t.Run("json-schema", func(t *testing.T) {
		// nothing to do here right now
	})

	t.Run("raw-schema", func(t *testing.T) {
		t.Fatalf("no raw-schema test case defined for testcases/TestAllOfWithRefsSchemaToStruct.md")
	})

	t.Run("static-schema", func(t *testing.T) {
		t.Fatalf("no static-schema test case defined for testcases/TestAllOfWithRefsSchemaToStruct.md")
	})

	t.Run("go-code", func(t *testing.T) {

		checkGoCode(t, schema, tc.GoCode())
	})

}

func TestAnyOfSchemaToStruct(t *testing.T) {

	tc := testcases.LoadAndParseTestCase("TestAnyOfSchemaToStruct")
	require.Equal(t, testCasesHash, testcases.GetHash(), "test cases hash mismatch, please run 'go generate ./...' to update the test cases hash")

	schema, err := generator.NewSchemaModel(tc.JSONSchema())
	require.NoError(t, err, "failed to parse schema")

	t.Run("json-schema", func(t *testing.T) {
		// nothing to do here right now
	})

	t.Run("raw-schema", func(t *testing.T) {
		t.Fatalf("no raw-schema test case defined for testcases/TestAnyOfSchemaToStruct.md")
	})

	t.Run("static-schema", func(t *testing.T) {
		t.Fatalf("no static-schema test case defined for testcases/TestAnyOfSchemaToStruct.md")
	})

	t.Run("go-code", func(t *testing.T) {

		checkGoCode(t, schema, tc.GoCode())
	})

}

func TestArrayOfReferencesSchemaToStruct(t *testing.T) {

	tc := testcases.LoadAndParseTestCase("TestArrayOfReferencesSchemaToStruct")
	require.Equal(t, testCasesHash, testcases.GetHash(), "test cases hash mismatch, please run 'go generate ./...' to update the test cases hash")

	schema, err := generator.NewSchemaModel(tc.JSONSchema())
	require.NoError(t, err, "failed to parse schema")

	t.Run("json-schema", func(t *testing.T) {
		// nothing to do here right now
	})

	t.Run("raw-schema", func(t *testing.T) {
		t.Fatalf("no raw-schema test case defined for testcases/TestArrayOfReferencesSchemaToStruct.md")
	})

	t.Run("static-schema", func(t *testing.T) {
		t.Fatalf("no static-schema test case defined for testcases/TestArrayOfReferencesSchemaToStruct.md")
	})

	t.Run("go-code", func(t *testing.T) {

		checkGoCode(t, schema, tc.GoCode())
	})

}

func TestBasicRefSchemaToStruct(t *testing.T) {

	tc := testcases.LoadAndParseTestCase("TestBasicRefSchemaToStruct")
	require.Equal(t, testCasesHash, testcases.GetHash(), "test cases hash mismatch, please run 'go generate ./...' to update the test cases hash")

	schema, err := generator.NewSchemaModel(tc.JSONSchema())
	require.NoError(t, err, "failed to parse schema")

	t.Run("json-schema", func(t *testing.T) {
		// nothing to do here right now
	})

	t.Run("raw-schema", func(t *testing.T) {
		t.Fatalf("no raw-schema test case defined for testcases/TestBasicRefSchemaToStruct.md")
	})

	t.Run("static-schema", func(t *testing.T) {
		t.Fatalf("no static-schema test case defined for testcases/TestBasicRefSchemaToStruct.md")
	})

	t.Run("go-code", func(t *testing.T) {

		checkGoCode(t, schema, tc.GoCode())
	})

}

func TestBasicSchemaToStruct(t *testing.T) {

	tc := testcases.LoadAndParseTestCase("TestBasicSchemaToStruct")
	require.Equal(t, testCasesHash, testcases.GetHash(), "test cases hash mismatch, please run 'go generate ./...' to update the test cases hash")

	schema, err := generator.NewSchemaModel(tc.JSONSchema())
	require.NoError(t, err, "failed to parse schema")

	t.Run("json-schema", func(t *testing.T) {
		// nothing to do here right now
	})

	t.Run("raw-schema", func(t *testing.T) {

		want := &generator.SchemaModel{
			SourceSchema: &jsonschema.Schema{
				Schema: ptr("http://json-schema.org/draft-07/schema#"),
				Title:  ptr("BasicExample"),
				Type:   typ("object"),
				Properties: &[]*jsonschema.NamedSchema{
					{
						Name: "id",
						Value: &jsonschema.Schema{
							Type: typ("string"),
						},
					},
					{
						Name: "count",
						Value: &jsonschema.Schema{
							Type: typ("integer"),
							Default: &yaml.Node{
								Kind:  yaml.ScalarNode,
								Value: "0",
							},
						},
					},
					{
						Name: "enabled",
						Value: &jsonschema.Schema{
							Type: typ("boolean"),
							Default: &yaml.Node{
								Kind:  yaml.ScalarNode,
								Value: "false",
							},
						},
					},
					{
						Name: "ratio",
						Value: &jsonschema.Schema{
							Type: typ("number"),
						},
					},
				},
				Required: &[]string{"id"},
			},
		}

		diff.RequireKnownValueEqual(t, want.SourceSchema, schema.SourceSchema)
	})

	t.Run("static-schema", func(t *testing.T) {

		f1 := &generator.StaticField{
			Name_:                "ID",
			JSONName_:            "id",
			Description_:         "",
			IsRequired_:          true,
			Type_:                "string",
			IsEnum_:              false,
			EnumTypeName_:        "IDType",
			EnumValues_:          nil,
			DefaultValue_:        nil,
			DefaultValueComment_: nil,
			ValidationRules_: []generator.ValidationRule{
				{
					Type:    generator.ValidationRequired,
					Message: "id is required",
					// Parnet: will be injected by the test case
					Values: "",
				},
			},
		}
		f2 := &generator.StaticField{
			Name_:                "Count",
			JSONName_:            "count",
			Description_:         "",
			IsRequired_:          false,
			Type_:                "*int",
			IsEnum_:              false,
			EnumTypeName_:        "CountType",
			EnumValues_:          nil,
			DefaultValue_:        ptr("0"),
			DefaultValueComment_: ptr("0"),
			ValidationRules_:     nil,
		}
		f3 := &generator.StaticField{
			Name_:                "Enabled",
			JSONName_:            "enabled",
			Description_:         "",
			IsRequired_:          false,
			Type_:                "*bool",
			IsEnum_:              false,
			EnumTypeName_:        "EnabledType",
			EnumValues_:          nil,
			DefaultValue_:        ptr("false"),
			DefaultValueComment_: ptr("false"),
			ValidationRules_:     nil,
		}
		f4 := &generator.StaticField{
			Name_:                "Ratio",
			JSONName_:            "ratio",
			Description_:         "",
			IsRequired_:          false,
			Type_:                "*float64",
			IsEnum_:              false,
			EnumTypeName_:        "RatioType",
			EnumValues_:          nil,
			DefaultValue_:        nil,
			DefaultValueComment_: nil,
			ValidationRules_:     nil,
		}

		s1 := &generator.StaticStruct{Name_: "BasicExample", Fields_: []generator.Field{f1, f2, f3, f4}}

		staticWant := &generator.StaticSchema{
			Package_: "models",
			Structs_: []generator.Struct{
				s1,
			},
			Enums_: nil,
			Imports_: []string{
				"encoding/json",
				"gitlab.com/tozd/go/errors",
			},
		}

		staticGot := generator.NewStaticSchema(schema)

		diff.RequireKnownValueEqual(t, staticWant, staticGot)
	})

	t.Run("go-code", func(t *testing.T) {

		checkGoCode(t, schema, tc.GoCode())
	})

}

func TestIntegerEnumSchemaToStruct(t *testing.T) {

	tc := testcases.LoadAndParseTestCase("TestIntegerEnumSchemaToStruct")
	require.Equal(t, testCasesHash, testcases.GetHash(), "test cases hash mismatch, please run 'go generate ./...' to update the test cases hash")

	schema, err := generator.NewSchemaModel(tc.JSONSchema())
	require.NoError(t, err, "failed to parse schema")

	t.Run("json-schema", func(t *testing.T) {
		// nothing to do here right now
	})

	t.Run("raw-schema", func(t *testing.T) {
		t.Fatalf("no raw-schema test case defined for testcases/TestIntegerEnumSchemaToStruct.md")
	})

	t.Run("static-schema", func(t *testing.T) {
		t.Fatalf("no static-schema test case defined for testcases/TestIntegerEnumSchemaToStruct.md")
	})

	t.Run("go-code", func(t *testing.T) {

		checkGoCode(t, schema, tc.GoCode())
	})

}

func TestNestedObjectDeep(t *testing.T) {

	tc := testcases.LoadAndParseTestCase("TestNestedObjectDeep")
	require.Equal(t, testCasesHash, testcases.GetHash(), "test cases hash mismatch, please run 'go generate ./...' to update the test cases hash")

	schema, err := generator.NewSchemaModel(tc.JSONSchema())
	require.NoError(t, err, "failed to parse schema")

	t.Run("json-schema", func(t *testing.T) {
		// nothing to do here right now
	})

	t.Run("raw-schema", func(t *testing.T) {
		t.Fatalf("no raw-schema test case defined for testcases/TestNestedObjectDeep.md")
	})

	t.Run("static-schema", func(t *testing.T) {
		t.Fatalf("no static-schema test case defined for testcases/TestNestedObjectDeep.md")
	})

	t.Run("go-code", func(t *testing.T) {

		checkGoCode(t, schema, tc.GoCode())
	})

}

func TestNestedObjectSimple(t *testing.T) {

	tc := testcases.LoadAndParseTestCase("TestNestedObjectSimple")
	require.Equal(t, testCasesHash, testcases.GetHash(), "test cases hash mismatch, please run 'go generate ./...' to update the test cases hash")

	schema, err := generator.NewSchemaModel(tc.JSONSchema())
	require.NoError(t, err, "failed to parse schema")

	t.Run("json-schema", func(t *testing.T) {
		// nothing to do here right now
	})

	t.Run("raw-schema", func(t *testing.T) {
		t.Fatalf("no raw-schema test case defined for testcases/TestNestedObjectSimple.md")
	})

	t.Run("static-schema", func(t *testing.T) {
		t.Fatalf("no static-schema test case defined for testcases/TestNestedObjectSimple.md")
	})

	t.Run("go-code", func(t *testing.T) {

		checkGoCode(t, schema, tc.GoCode())
	})

}

func TestNestedObjectWithOptional(t *testing.T) {

	tc := testcases.LoadAndParseTestCase("TestNestedObjectWithOptional")
	require.Equal(t, testCasesHash, testcases.GetHash(), "test cases hash mismatch, please run 'go generate ./...' to update the test cases hash")

	schema, err := generator.NewSchemaModel(tc.JSONSchema())
	require.NoError(t, err, "failed to parse schema")

	t.Run("json-schema", func(t *testing.T) {
		// nothing to do here right now
	})

	t.Run("raw-schema", func(t *testing.T) {
		t.Fatalf("no raw-schema test case defined for testcases/TestNestedObjectWithOptional.md")
	})

	t.Run("static-schema", func(t *testing.T) {
		t.Fatalf("no static-schema test case defined for testcases/TestNestedObjectWithOptional.md")
	})

	t.Run("go-code", func(t *testing.T) {

		checkGoCode(t, schema, tc.GoCode())
	})

}

func TestOneOfSchemaToStruct(t *testing.T) {

	tc := testcases.LoadAndParseTestCase("TestOneOfSchemaToStruct")
	require.Equal(t, testCasesHash, testcases.GetHash(), "test cases hash mismatch, please run 'go generate ./...' to update the test cases hash")

	schema, err := generator.NewSchemaModel(tc.JSONSchema())
	require.NoError(t, err, "failed to parse schema")

	t.Run("json-schema", func(t *testing.T) {
		// nothing to do here right now
	})

	t.Run("raw-schema", func(t *testing.T) {
		t.Fatalf("no raw-schema test case defined for testcases/TestOneOfSchemaToStruct.md")
	})

	t.Run("static-schema", func(t *testing.T) {
		t.Fatalf("no static-schema test case defined for testcases/TestOneOfSchemaToStruct.md")
	})

	t.Run("go-code", func(t *testing.T) {

		checkGoCode(t, schema, tc.GoCode())
	})

}

func TestPatternPropertiesSchemaToStruct(t *testing.T) {

	tc := testcases.LoadAndParseTestCase("TestPatternPropertiesSchemaToStruct")
	require.Equal(t, testCasesHash, testcases.GetHash(), "test cases hash mismatch, please run 'go generate ./...' to update the test cases hash")

	schema, err := generator.NewSchemaModel(tc.JSONSchema())
	require.NoError(t, err, "failed to parse schema")

	t.Run("json-schema", func(t *testing.T) {
		// nothing to do here right now
	})

	t.Run("raw-schema", func(t *testing.T) {
		t.Fatalf("no raw-schema test case defined for testcases/TestPatternPropertiesSchemaToStruct.md")
	})

	t.Run("static-schema", func(t *testing.T) {
		t.Fatalf("no static-schema test case defined for testcases/TestPatternPropertiesSchemaToStruct.md")
	})

	t.Run("go-code", func(t *testing.T) {

		checkGoCode(t, schema, tc.GoCode())
	})

}

func TestRequiredFieldsSchemaToStruct(t *testing.T) {

	tc := testcases.LoadAndParseTestCase("TestRequiredFieldsSchemaToStruct")
	require.Equal(t, testCasesHash, testcases.GetHash(), "test cases hash mismatch, please run 'go generate ./...' to update the test cases hash")

	schema, err := generator.NewSchemaModel(tc.JSONSchema())
	require.NoError(t, err, "failed to parse schema")

	t.Run("json-schema", func(t *testing.T) {
		// nothing to do here right now
	})

	t.Run("raw-schema", func(t *testing.T) {
		t.Fatalf("no raw-schema test case defined for testcases/TestRequiredFieldsSchemaToStruct.md")
	})

	t.Run("static-schema", func(t *testing.T) {
		t.Fatalf("no static-schema test case defined for testcases/TestRequiredFieldsSchemaToStruct.md")
	})

	t.Run("go-code", func(t *testing.T) {

		checkGoCode(t, schema, tc.GoCode())
	})

}

func TestSchemaDocumentation(t *testing.T) {

	tc := testcases.LoadAndParseTestCase("TestSchemaDocumentation")
	require.Equal(t, testCasesHash, testcases.GetHash(), "test cases hash mismatch, please run 'go generate ./...' to update the test cases hash")

	schema, err := generator.NewSchemaModel(tc.JSONSchema())
	require.NoError(t, err, "failed to parse schema")

	t.Run("json-schema", func(t *testing.T) {
		// nothing to do here right now
	})

	t.Run("raw-schema", func(t *testing.T) {
		t.Fatalf("no raw-schema test case defined for testcases/TestSchemaDocumentation.md")
	})

	t.Run("static-schema", func(t *testing.T) {
		t.Fatalf("no static-schema test case defined for testcases/TestSchemaDocumentation.md")
	})

	t.Run("go-code", func(t *testing.T) {

		checkGoCode(t, schema, tc.GoCode())
	})

}

func TestStringEnumSchemaToStruct(t *testing.T) {

	tc := testcases.LoadAndParseTestCase("TestStringEnumSchemaToStruct")
	require.Equal(t, testCasesHash, testcases.GetHash(), "test cases hash mismatch, please run 'go generate ./...' to update the test cases hash")

	schema, err := generator.NewSchemaModel(tc.JSONSchema())
	require.NoError(t, err, "failed to parse schema")

	t.Run("json-schema", func(t *testing.T) {
		// nothing to do here right now
	})

	t.Run("raw-schema", func(t *testing.T) {
		t.Fatalf("no raw-schema test case defined for testcases/TestStringEnumSchemaToStruct.md")
	})

	t.Run("static-schema", func(t *testing.T) {
		t.Fatalf("no static-schema test case defined for testcases/TestStringEnumSchemaToStruct.md")
	})

	t.Run("go-code", func(t *testing.T) {

		checkGoCode(t, schema, tc.GoCode())
	})

}

func TestTypeNamingConventions(t *testing.T) {

	tc := testcases.LoadAndParseTestCase("TestTypeNamingConventions")
	require.Equal(t, testCasesHash, testcases.GetHash(), "test cases hash mismatch, please run 'go generate ./...' to update the test cases hash")

	schema, err := generator.NewSchemaModel(tc.JSONSchema())
	require.NoError(t, err, "failed to parse schema")

	t.Run("json-schema", func(t *testing.T) {
		// nothing to do here right now
	})

	t.Run("raw-schema", func(t *testing.T) {
		t.Fatalf("no raw-schema test case defined for testcases/TestTypeNamingConventions.md")
	})

	t.Run("static-schema", func(t *testing.T) {
		t.Fatalf("no static-schema test case defined for testcases/TestTypeNamingConventions.md")
	})

	t.Run("go-code", func(t *testing.T) {

		checkGoCode(t, schema, tc.GoCode())
	})

}
