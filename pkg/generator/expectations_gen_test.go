package generator_test

import (
	"github.com/google/gnostic/jsonschema"
	"github.com/walteh/schema2go/gen/mockery"
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

func TestAllOfSchemaToStruct_RawSchemaModel(t *testing.T) {

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

	got := mustLoadSchemaModel(t, tc.GenerateInput())

	checkRawSchema(t, want.SourceSchema, got.SourceSchema)
}

func TestAllOfSchemaToStruct_SchemaModelMock(t *testing.T) {

	var mockField1 = mockery.NewMockField_generator(t)
	mockField1.EXPECT().Description().Return("name")
	mockField1.EXPECT().Type().Return("string")
	mockField1.EXPECT().IsRequired().Return(false)
	mockField1.EXPECT().IsEnum().Return(false)
	mockField1.EXPECT().EnumTypeName().Return("")
	mockField1.EXPECT().EnumValues().Return([]*generator.EnumValue{})
	mockField1.EXPECT().DefaultValue().Return(nil)
	mockField1.EXPECT().DefaultValueComment().Return(nil)
	mockField1.EXPECT().ValidationRules().Return([]*generator.ValidationRule{})

	var mockField2 = mockery.NewMockField_generator(t)
	mockField2.EXPECT().Description().Return("age")
	mockField2.EXPECT().Type().Return("integer")
	mockField2.EXPECT().IsRequired().Return(false)
	mockField2.EXPECT().IsEnum().Return(false)
	mockField2.EXPECT().EnumTypeName().Return("")
	mockField2.EXPECT().EnumValues().Return([]*generator.EnumValue{})
	mockField2.EXPECT().DefaultValue().Return(nil)
	mockField2.EXPECT().DefaultValueComment().Return(nil)
	mockField2.EXPECT().ValidationRules().Return([]*generator.ValidationRule{})

	var mockStruct = mockery.NewMockStruct_generator(t)
	mockStruct.EXPECT().Description().Return("AllOfExample")
	mockStruct.EXPECT().Fields().Return([]generator.Field{mockField1, mockField2})
	mockStruct.EXPECT().HasAllOf().Return(true)
	mockStruct.EXPECT().HasCustomMarshaling().Return(false)
	mockStruct.EXPECT().HasDefaults().Return(false)
	mockStruct.EXPECT().HasValidation().Return(false)

	var mockSchema = mockery.NewMockSchema_generator(t)
	mockSchema.EXPECT().Enums().Return([]*generator.EnumModel{})
	mockSchema.EXPECT().Imports().Return([]string{})
	mockSchema.EXPECT().Package().Return("models")
	mockSchema.EXPECT().Structs().Return([]generator.Struct{mockStruct})

	tc := testcases.LoadAndParseTestCase("TestAllOfSchemaToStruct")

	got := mustLoadSchemaModel(t, tc.GenerateInput())

	checkSchemaMock(t, mockSchema, got)
}

func TestAllOfSchemaToStruct_GoCode(t *testing.T) {
	tc := testcases.LoadAndParseTestCase("TestAllOfSchemaToStruct")

	checkGoCode(t, tc.GenerateInput(), tc.GenerateExpectedGoCode())
}

func TestAllOfWithRefsSchemaToStruct_GoCode(t *testing.T) {
	tc := testcases.LoadAndParseTestCase("TestAllOfWithRefsSchemaToStruct")

	checkGoCode(t, tc.GenerateInput(), tc.GenerateExpectedGoCode())
}

func TestAnyOfSchemaToStruct_GoCode(t *testing.T) {
	tc := testcases.LoadAndParseTestCase("TestAnyOfSchemaToStruct")

	checkGoCode(t, tc.GenerateInput(), tc.GenerateExpectedGoCode())
}

func TestArrayOfReferencesSchemaToStruct_GoCode(t *testing.T) {
	tc := testcases.LoadAndParseTestCase("TestArrayOfReferencesSchemaToStruct")

	checkGoCode(t, tc.GenerateInput(), tc.GenerateExpectedGoCode())
}

func TestBasicRefSchemaToStruct_GoCode(t *testing.T) {
	tc := testcases.LoadAndParseTestCase("TestBasicRefSchemaToStruct")

	checkGoCode(t, tc.GenerateInput(), tc.GenerateExpectedGoCode())
}

func TestBasicSchemaToStruct_GoCode(t *testing.T) {
	tc := testcases.LoadAndParseTestCase("TestBasicSchemaToStruct")

	checkGoCode(t, tc.GenerateInput(), tc.GenerateExpectedGoCode())
}

func TestIntegerEnumSchemaToStruct_GoCode(t *testing.T) {
	tc := testcases.LoadAndParseTestCase("TestIntegerEnumSchemaToStruct")

	checkGoCode(t, tc.GenerateInput(), tc.GenerateExpectedGoCode())
}

func TestNestedObjectDeep_GoCode(t *testing.T) {
	tc := testcases.LoadAndParseTestCase("TestNestedObjectDeep")

	checkGoCode(t, tc.GenerateInput(), tc.GenerateExpectedGoCode())
}

func TestNestedObjectSimple_GoCode(t *testing.T) {
	tc := testcases.LoadAndParseTestCase("TestNestedObjectSimple")

	checkGoCode(t, tc.GenerateInput(), tc.GenerateExpectedGoCode())
}

func TestNestedObjectWithOptional_GoCode(t *testing.T) {
	tc := testcases.LoadAndParseTestCase("TestNestedObjectWithOptional")

	checkGoCode(t, tc.GenerateInput(), tc.GenerateExpectedGoCode())
}

func TestOneOfSchemaToStruct_GoCode(t *testing.T) {
	tc := testcases.LoadAndParseTestCase("TestOneOfSchemaToStruct")

	checkGoCode(t, tc.GenerateInput(), tc.GenerateExpectedGoCode())
}

func TestPatternPropertiesSchemaToStruct_GoCode(t *testing.T) {
	tc := testcases.LoadAndParseTestCase("TestPatternPropertiesSchemaToStruct")

	checkGoCode(t, tc.GenerateInput(), tc.GenerateExpectedGoCode())
}

func TestRequiredFieldsSchemaToStruct_GoCode(t *testing.T) {
	tc := testcases.LoadAndParseTestCase("TestRequiredFieldsSchemaToStruct")

	checkGoCode(t, tc.GenerateInput(), tc.GenerateExpectedGoCode())
}

func TestSchemaDocumentation_GoCode(t *testing.T) {
	tc := testcases.LoadAndParseTestCase("TestSchemaDocumentation")

	checkGoCode(t, tc.GenerateInput(), tc.GenerateExpectedGoCode())
}

func TestStringEnumSchemaToStruct_GoCode(t *testing.T) {
	tc := testcases.LoadAndParseTestCase("TestStringEnumSchemaToStruct")

	checkGoCode(t, tc.GenerateInput(), tc.GenerateExpectedGoCode())
}

func TestTypeNamingConventions_GoCode(t *testing.T) {
	tc := testcases.LoadAndParseTestCase("TestTypeNamingConventions")

	checkGoCode(t, tc.GenerateInput(), tc.GenerateExpectedGoCode())
}
