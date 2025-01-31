package generator_test

import (
	"github.com/google/gnostic/jsonschema"
	"github.com/walteh/schema2go/pkg/diff"
	"github.com/walteh/schema2go/pkg/generator"
	"github.com/walteh/schema2go/pkg/generator/testcases"
	"testing"
)

func mustLoadSchemaModel(t *testing.T, input string) *generator.SchemaModel {
	schema, err := generator.NewSchemaModel(input)
	if err != nil {
		t.Fatalf("Failed to parse schema: %v", err)
	}
	return schema
}

func TestAllOfSchemaToStruct(t *testing.T) {

	t.Run("raw-schema", func(t *testing.T) {

		want := &generator.SchemaModel{
			SourceSchema: &jsonschema.Schema{
				Title: ptr("AllOfExample"),
				Type:  typ("object"),
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

		tc := testcases.LoadAndParseTestCase("TestAllOfSchemaToStruct")

		got := mustLoadSchemaModel(t, tc.JSONSchema())

		diff.RequireKnownValueEqual(t, want.SourceSchema, got.SourceSchema)
	})

	t.Run("static-schema", func(t *testing.T) {

		staticWant := &generator.StaticSchema{
			Package_: "models",
			Structs_: []generator.Struct{
				&generator.StaticStruct{
					Name_: "AllOfExample",
					Fields_: []generator.Field{
						&generator.StaticField{Name_: "Name_AllOf", JSONName_: "name", Type_: "string"},
						&generator.StaticField{Name_: "Age_AllOf", JSONName_: "age", Type_: "int"},
					},
				},
			},
		}

		tc := testcases.LoadAndParseTestCase("TestAllOfSchemaToStruct")

		got := mustLoadSchemaModel(t, tc.JSONSchema())

		staticGot := generator.NewStaticSchema(got)

		diff.RequireKnownValueEqual(t, staticWant, staticGot)
	})

	t.Run("go-code", func(t *testing.T) {
		tc := testcases.LoadAndParseTestCase("TestAllOfSchemaToStruct")

		got := mustLoadSchemaModel(t, tc.JSONSchema())

		checkGoCode(t, got, tc.GoCode())
	})

}

func TestAllOfWithRefsSchemaToStruct(t *testing.T) {

	t.Run("raw-schema", func(t *testing.T) {
		t.Fatalf("no raw-schema test case defined for testcases/TestAllOfWithRefsSchemaToStruct.md")
	})

	t.Run("static-schema", func(t *testing.T) {
		t.Fatalf("no static-schema test case defined for testcases/TestAllOfWithRefsSchemaToStruct.md")
	})

	t.Run("go-code", func(t *testing.T) {
		tc := testcases.LoadAndParseTestCase("TestAllOfWithRefsSchemaToStruct")

		got := mustLoadSchemaModel(t, tc.JSONSchema())

		checkGoCode(t, got, tc.GoCode())
	})

}

func TestAnyOfSchemaToStruct(t *testing.T) {

	t.Run("raw-schema", func(t *testing.T) {
		t.Fatalf("no raw-schema test case defined for testcases/TestAnyOfSchemaToStruct.md")
	})

	t.Run("static-schema", func(t *testing.T) {
		t.Fatalf("no static-schema test case defined for testcases/TestAnyOfSchemaToStruct.md")
	})

	t.Run("go-code", func(t *testing.T) {
		tc := testcases.LoadAndParseTestCase("TestAnyOfSchemaToStruct")

		got := mustLoadSchemaModel(t, tc.JSONSchema())

		checkGoCode(t, got, tc.GoCode())
	})

}

func TestArrayOfReferencesSchemaToStruct(t *testing.T) {

	t.Run("raw-schema", func(t *testing.T) {
		t.Fatalf("no raw-schema test case defined for testcases/TestArrayOfReferencesSchemaToStruct.md")
	})

	t.Run("static-schema", func(t *testing.T) {
		t.Fatalf("no static-schema test case defined for testcases/TestArrayOfReferencesSchemaToStruct.md")
	})

	t.Run("go-code", func(t *testing.T) {
		tc := testcases.LoadAndParseTestCase("TestArrayOfReferencesSchemaToStruct")

		got := mustLoadSchemaModel(t, tc.JSONSchema())

		checkGoCode(t, got, tc.GoCode())
	})

}

func TestBasicRefSchemaToStruct(t *testing.T) {

	t.Run("raw-schema", func(t *testing.T) {
		t.Fatalf("no raw-schema test case defined for testcases/TestBasicRefSchemaToStruct.md")
	})

	t.Run("static-schema", func(t *testing.T) {
		t.Fatalf("no static-schema test case defined for testcases/TestBasicRefSchemaToStruct.md")
	})

	t.Run("go-code", func(t *testing.T) {
		tc := testcases.LoadAndParseTestCase("TestBasicRefSchemaToStruct")

		got := mustLoadSchemaModel(t, tc.JSONSchema())

		checkGoCode(t, got, tc.GoCode())
	})

}

func TestBasicSchemaToStruct(t *testing.T) {

	t.Run("raw-schema", func(t *testing.T) {
		t.Fatalf("no raw-schema test case defined for testcases/TestBasicSchemaToStruct.md")
	})

	t.Run("static-schema", func(t *testing.T) {
		t.Fatalf("no static-schema test case defined for testcases/TestBasicSchemaToStruct.md")
	})

	t.Run("go-code", func(t *testing.T) {
		tc := testcases.LoadAndParseTestCase("TestBasicSchemaToStruct")

		got := mustLoadSchemaModel(t, tc.JSONSchema())

		checkGoCode(t, got, tc.GoCode())
	})

}

func TestIntegerEnumSchemaToStruct(t *testing.T) {

	t.Run("raw-schema", func(t *testing.T) {
		t.Fatalf("no raw-schema test case defined for testcases/TestIntegerEnumSchemaToStruct.md")
	})

	t.Run("static-schema", func(t *testing.T) {
		t.Fatalf("no static-schema test case defined for testcases/TestIntegerEnumSchemaToStruct.md")
	})

	t.Run("go-code", func(t *testing.T) {
		tc := testcases.LoadAndParseTestCase("TestIntegerEnumSchemaToStruct")

		got := mustLoadSchemaModel(t, tc.JSONSchema())

		checkGoCode(t, got, tc.GoCode())
	})

}

func TestNestedObjectDeep(t *testing.T) {

	t.Run("raw-schema", func(t *testing.T) {
		t.Fatalf("no raw-schema test case defined for testcases/TestNestedObjectDeep.md")
	})

	t.Run("static-schema", func(t *testing.T) {
		t.Fatalf("no static-schema test case defined for testcases/TestNestedObjectDeep.md")
	})

	t.Run("go-code", func(t *testing.T) {
		tc := testcases.LoadAndParseTestCase("TestNestedObjectDeep")

		got := mustLoadSchemaModel(t, tc.JSONSchema())

		checkGoCode(t, got, tc.GoCode())
	})

}

func TestNestedObjectSimple(t *testing.T) {

	t.Run("raw-schema", func(t *testing.T) {
		t.Fatalf("no raw-schema test case defined for testcases/TestNestedObjectSimple.md")
	})

	t.Run("static-schema", func(t *testing.T) {
		t.Fatalf("no static-schema test case defined for testcases/TestNestedObjectSimple.md")
	})

	t.Run("go-code", func(t *testing.T) {
		tc := testcases.LoadAndParseTestCase("TestNestedObjectSimple")

		got := mustLoadSchemaModel(t, tc.JSONSchema())

		checkGoCode(t, got, tc.GoCode())
	})

}

func TestNestedObjectWithOptional(t *testing.T) {

	t.Run("raw-schema", func(t *testing.T) {
		t.Fatalf("no raw-schema test case defined for testcases/TestNestedObjectWithOptional.md")
	})

	t.Run("static-schema", func(t *testing.T) {
		t.Fatalf("no static-schema test case defined for testcases/TestNestedObjectWithOptional.md")
	})

	t.Run("go-code", func(t *testing.T) {
		tc := testcases.LoadAndParseTestCase("TestNestedObjectWithOptional")

		got := mustLoadSchemaModel(t, tc.JSONSchema())

		checkGoCode(t, got, tc.GoCode())
	})

}

func TestOneOfSchemaToStruct(t *testing.T) {

	t.Run("raw-schema", func(t *testing.T) {
		t.Fatalf("no raw-schema test case defined for testcases/TestOneOfSchemaToStruct.md")
	})

	t.Run("static-schema", func(t *testing.T) {
		t.Fatalf("no static-schema test case defined for testcases/TestOneOfSchemaToStruct.md")
	})

	t.Run("go-code", func(t *testing.T) {
		tc := testcases.LoadAndParseTestCase("TestOneOfSchemaToStruct")

		got := mustLoadSchemaModel(t, tc.JSONSchema())

		checkGoCode(t, got, tc.GoCode())
	})

}

func TestPatternPropertiesSchemaToStruct(t *testing.T) {

	t.Run("raw-schema", func(t *testing.T) {
		t.Fatalf("no raw-schema test case defined for testcases/TestPatternPropertiesSchemaToStruct.md")
	})

	t.Run("static-schema", func(t *testing.T) {
		t.Fatalf("no static-schema test case defined for testcases/TestPatternPropertiesSchemaToStruct.md")
	})

	t.Run("go-code", func(t *testing.T) {
		tc := testcases.LoadAndParseTestCase("TestPatternPropertiesSchemaToStruct")

		got := mustLoadSchemaModel(t, tc.JSONSchema())

		checkGoCode(t, got, tc.GoCode())
	})

}

func TestRequiredFieldsSchemaToStruct(t *testing.T) {

	t.Run("raw-schema", func(t *testing.T) {
		t.Fatalf("no raw-schema test case defined for testcases/TestRequiredFieldsSchemaToStruct.md")
	})

	t.Run("static-schema", func(t *testing.T) {
		t.Fatalf("no static-schema test case defined for testcases/TestRequiredFieldsSchemaToStruct.md")
	})

	t.Run("go-code", func(t *testing.T) {
		tc := testcases.LoadAndParseTestCase("TestRequiredFieldsSchemaToStruct")

		got := mustLoadSchemaModel(t, tc.JSONSchema())

		checkGoCode(t, got, tc.GoCode())
	})

}

func TestSchemaDocumentation(t *testing.T) {

	t.Run("raw-schema", func(t *testing.T) {
		t.Fatalf("no raw-schema test case defined for testcases/TestSchemaDocumentation.md")
	})

	t.Run("static-schema", func(t *testing.T) {
		t.Fatalf("no static-schema test case defined for testcases/TestSchemaDocumentation.md")
	})

	t.Run("go-code", func(t *testing.T) {
		tc := testcases.LoadAndParseTestCase("TestSchemaDocumentation")

		got := mustLoadSchemaModel(t, tc.JSONSchema())

		checkGoCode(t, got, tc.GoCode())
	})

}

func TestStringEnumSchemaToStruct(t *testing.T) {

	t.Run("raw-schema", func(t *testing.T) {
		t.Fatalf("no raw-schema test case defined for testcases/TestStringEnumSchemaToStruct.md")
	})

	t.Run("static-schema", func(t *testing.T) {
		t.Fatalf("no static-schema test case defined for testcases/TestStringEnumSchemaToStruct.md")
	})

	t.Run("go-code", func(t *testing.T) {
		tc := testcases.LoadAndParseTestCase("TestStringEnumSchemaToStruct")

		got := mustLoadSchemaModel(t, tc.JSONSchema())

		checkGoCode(t, got, tc.GoCode())
	})

}

func TestTypeNamingConventions(t *testing.T) {

	t.Run("raw-schema", func(t *testing.T) {
		t.Fatalf("no raw-schema test case defined for testcases/TestTypeNamingConventions.md")
	})

	t.Run("static-schema", func(t *testing.T) {
		t.Fatalf("no static-schema test case defined for testcases/TestTypeNamingConventions.md")
	})

	t.Run("go-code", func(t *testing.T) {
		tc := testcases.LoadAndParseTestCase("TestTypeNamingConventions")

		got := mustLoadSchemaModel(t, tc.JSONSchema())

		checkGoCode(t, got, tc.GoCode())
	})

}
