# Schema2Go Generator Implementation

```ascii
                    ┌──────────────────────┐
                    │ JSON Schema Input    │
                    └───────────┬──────────┘
                                │
                                ▼
            ┌─────────────────────────────────────┐
            │       Schema2Go Generator Core      │
            │                                     │
            │  ┌───────────────────────────────┐  │
            │  │    Name Scoping System        │  │
            │  └───────────────────────────────┘  │
            │                                     │
            │  ┌───────────────────────────────┐  │
            │  │    Combination Type Handlers  │  │
            │  │                               │  │
            │  │    ┌───────┐┌──────┐┌──────┐  │  │
            │  │    │allOf  ││oneOf ││anyOf │  │  │
            │  │    └───────┘└──────┘└──────┘  │  │
            │  └───────────────────────────────┘  │
            │                                     │
            │  ┌───────────────────────────────┐  │
            │  │   Marshaling & Validation     │  │
            │  └───────────────────────────────┘  │
            └─────────────────────────────────────┘
                                │
                                ▼
                    ┌──────────────────────┐
                    │  Go Structs Output   │
                    └──────────────────────┘
```

## 🎯 Core Objective

Implement a schema generator that directly converts JSON Schema to Go structs with proper handling of combination types (allOf, oneOf, anyOf) - focusing on the same requirements as outlined in the postprocessor.mdc specification.

## 🔍 Key Focus Areas

1. A robust name scoping system to handle nested and complex types
2. Direct implementation of combination types in the schema generator
3. Custom JSON marshaling/unmarshaling with proper validation
4. Consistent naming conventions matching the specified patterns

## 💻 Implementation Strategy

### 1. Name Scoping System Enhancements

The current system needs improvement to properly handle nested types and combinations:

```go
// Improved name scoping mechanism
type NameScope struct {
    parts []string
}

// New scope from parent type and field name
func NewNestedScope(parentTypeName, fieldName string) *NameScope {
    return &NameScope{
        parts: []string{parentTypeName, fieldName},
    }
}

// Get combination type name
func (ns *NameScope) CombinationTypeName(combinationType string) string {
    return strings.Join(ns.parts, "_") + "_" + combinationType
}

// Get field name for combination types
func (ns *NameScope) FieldWithCombinationSuffix(fieldName, combinationType string) string {
    return fieldName + "_" + combinationType
}

// Get value field name for specific types in combinations
func GetValueFieldName(typeName, combinationType string) string {
    return typeName + "Value_" + combinationType
}
```

### 2. AllOf Implementation

AllOf must combine multiple schema properties into a single struct, following naming conventions:

```go
// In schema_generator.go
func (g *Generator) generateAllOfType(schema *jsonschema.Schema, scope *NameScope) (*codegen.StructType, error) {
    // Create the struct
    structType := &codegen.StructType{
        Name: scope.typeName,
    }

    // Process each schema in the allOf array
    allOfSchemas := parser.GetAllOf(schema)
    for _, subSchema := range allOfSchemas {
        // For references, resolve them first
        if subSchema.Ref != nil {
            resolvedSchema, err := g.resolveReference(*subSchema.Ref)
            if err != nil {
                return nil, err
            }
            subSchema = resolvedSchema
        }

        // Get properties from this schema
        props := parser.GetProperties(subSchema)
        for propName, propSchema := range props {
            // Add suffix to avoid field name collisions
            fieldName := scope.FieldWithCombinationSuffix(propName, "AllOf")

            // Generate field type
            fieldType, err := g.generateTypeForProperty(propSchema, propName, scope)
            if err != nil {
                return nil, err
            }

            // Add as struct field
            field := &codegen.Field{
                Name: fieldName,
                Type: fieldType,
                Tags: map[string]string{
                    "json": propName + ",omitempty",
                },
                Required: parser.IsRequired(subSchema, propName),
            }
            structType.Fields = append(structType.Fields, field)
        }
    }

    // Add custom marshaling methods
    g.addAllOfMarshaling(structType)

    return structType, nil
}
```

### 3. OneOf Implementation

OneOf requires creating a wrapper struct with fields for each possible type:

```go
// In schema_generator.go
func (g *Generator) generateOneOfType(schema *jsonschema.Schema, fieldName string, parentScope *NameScope) (*codegen.StructType, error) {
    // Create the wrapper struct
    oneOfTypeName := parentScope.CombinationTypeName(fieldName + "_OneOf")
    structType := &codegen.StructType{
        Name: oneOfTypeName,
    }

    // Process each schema in the oneOf array
    oneOfSchemas := parser.GetOneOf(schema)
    for _, subSchema := range oneOfSchemas {
        var fieldType codegen.Type
        var fieldName string

        if subSchema.Type != nil && subSchema.Type.String != nil {
            // Basic type field
            typeName := *subSchema.Type.String
            fieldName = GetValueFieldName(typeName, "OneOf")
            fieldType = g.generateBasicType(typeName)
        } else if subSchema.Ref != nil {
            // Reference field
            refName := getRefShortName(*subSchema.Ref)
            fieldName = GetValueFieldName("Ref", "OneOf")
            fieldType = &codegen.NamedType{Name: refName}
        } else if parser.HasEnum(subSchema) {
            // Enum field
            enumTypeName := oneOfTypeName + "_Enum"
            fieldName = GetValueFieldName("Enum", "OneOf")
            fieldType = &codegen.NamedType{Name: enumTypeName}

            // Generate the enum type
            g.generateEnumType(subSchema, enumTypeName)
        }

        // All fields are pointers to make optional
        ptrType := &codegen.PointerType{Elem: fieldType}

        // Add field to struct
        structType.Fields = append(structType.Fields, &codegen.Field{
            Name: fieldName,
            Type: ptrType,
            Tags: map[string]string{
                "json": strings.ToLower(strings.TrimSuffix(fieldName, "_OneOf")) + ",omitempty",
            },
        })
    }

    // Add validation and marshaling methods
    g.addOneOfValidation(structType)
    g.addOneOfMarshaling(structType)

    return structType, nil
}
```

### 4. AnyOf Implementation

AnyOf is similar to OneOf but validates that at least one field is set:

```go
// In schema_generator.go
func (g *Generator) generateAnyOfType(schema *jsonschema.Schema, fieldName string, parentScope *NameScope) (*codegen.StructType, error) {
    // Create the wrapper struct
    anyOfTypeName := parentScope.CombinationTypeName(fieldName + "_AnyOf")
    structType := &codegen.StructType{
        Name: anyOfTypeName,
    }

    // Process each schema in the anyOf array
    anyOfSchemas := parser.GetAnyOf(schema)
    for _, subSchema := range anyOfSchemas {
        var fieldType codegen.Type
        var fieldName string

        if subSchema.Type != nil && subSchema.Type.String != nil {
            // Basic type field
            typeName := *subSchema.Type.String
            fieldName = typeName + "Value_AnyOf"
            fieldType = g.generateBasicType(typeName)
        } else if subSchema.Ref != nil {
            // Reference field
            refName := getRefShortName(*subSchema.Ref)
            fieldName = "RefValue_AnyOf"
            fieldType = &codegen.NamedType{Name: refName}
        }

        // All fields are pointers to make optional
        ptrType := &codegen.PointerType{Elem: fieldType}

        // Add field to struct
        structType.Fields = append(structType.Fields, &codegen.Field{
            Name: fieldName,
            Type: ptrType,
            Tags: map[string]string{
                "json": strings.ToLower(strings.TrimSuffix(fieldName, "_AnyOf")) + "_value,omitempty",
            },
        })
    }

    // Add validation and marshaling methods
    g.addAnyOfValidation(structType)
    g.addAnyOfMarshaling(structType)

    return structType, nil
}
```

### 5. Array with Combination Types

Handle arrays containing items with combination types:

```go
// In schema_generator.go
func (g *Generator) generateArrayType(schema *jsonschema.Schema, field string, parentScope *NameScope) (*codegen.ArrayType, error) {
    items := parser.GetArrayItems(schema)
    if items == nil {
        return &codegen.ArrayType{ElementType: &codegen.AnyType{}}, nil
    }

    itemSchema := items.Schema
    if itemSchema == nil {
        return nil, errors.New("array with tuple validation not yet supported")
    }

    var elementType codegen.Type
    var err error

    if parser.HasOneOf(itemSchema) {
        // OneOf items need a special type
        itemTypeName := parentScope.typeName + "_" + field + "_Item"
        elementType, err = g.generateOneOfType(itemSchema, "Item", &NameScope{parts: []string{itemTypeName}})
    } else if parser.HasAnyOf(itemSchema) {
        // AnyOf items need a special type
        itemTypeName := parentScope.typeName + "_" + field + "_Item"
        elementType, err = g.generateAnyOfType(itemSchema, "Item", &NameScope{parts: []string{itemTypeName}})
    } else if parser.HasAllOf(itemSchema) {
        // AllOf items need a special type
        itemTypeName := parentScope.typeName + "_" + field + "_Item"
        elementType, err = g.generateAllOfType(itemSchema, &NameScope{parts: []string{itemTypeName}})
    } else {
        // Regular type
        elementType, err = g.generateType(itemSchema, field, parentScope)
    }

    if err != nil {
        return nil, err
    }

    return &codegen.ArrayType{ElementType: elementType}, nil
}
```

### 6. Pattern Properties Implementation

Handle pattern properties in schema:

```go
// In schema_generator.go
func (g *Generator) generatePatternProperties(schema *jsonschema.Schema, scope *NameScope) error {
    if schema.PatternProperties == nil {
        return nil
    }

    patternProps := *schema.PatternProperties
    for _, patternProp := range patternProps {
        pattern := patternProp.Name
        valueSchema := patternProp.Value

        // Convert pattern to field name
        fieldName := patternToFieldName(pattern) + "_Pattern"

        // Generate type for pattern values
        valueType, err := g.generateType(valueSchema, "Value", scope)
        if err != nil {
            return err
        }

        // Create map field
        mapType := &codegen.MapType{
            KeyType: &codegen.StringType{},
            ValueType: valueType,
        }

        // Add field to struct
        g.currentStruct.Fields = append(g.currentStruct.Fields, &codegen.Field{
            Name: fieldName,
            Type: mapType,
            Tags: map[string]string{"json": "-"},
            Comments: []string{fmt.Sprintf("Pattern: %s", pattern)},
        })
    }

    // Generate custom marshaling for pattern properties
    g.addPatternPropertiesMarshaling(g.currentStruct, schema)

    return nil
}

// Helper to convert regex pattern to field name
func patternToFieldName(pattern string) string {
    // Convert "^S_" to "String", "^N_" to "Number", etc.
    pattern = strings.TrimPrefix(pattern, "^")
    if strings.HasPrefix(pattern, "S_") {
        return "String"
    } else if strings.HasPrefix(pattern, "N_") {
        return "Number"
    }
    // Convert other patterns
    return "Pattern" + strings.NewReplacer(
        "^", "", "$", "", ".", "", "*", "", "+", "",
        "[", "", "]", "", "(", "", ")", "", "\\", "",
    ).Replace(pattern)
}
```

### 7. JSON Marshaling and Validation

Add custom marshaling methods to generated types:

```go
func (g *Generator) addOneOfValidation(structType *codegen.StructType) {
    // Add Validate method
    validateMethod := &codegen.Method{
        Name: "Validate",
        Receiver: codegen.Receiver{
            Name: "x",
            Type: structType,
        },
        Returns: []*codegen.Field{
            {Type: &codegen.ErrorType{}},
        },
        Body: `
    // Check that exactly one field is set
    count := 0
    {{range .Fields}}
    if x.{{.Name}} != nil {
        count++
    }
    {{end}}

    if count == 0 {
        return errors.Errorf("invalid {{.Name}}: exactly one field must be set")
    }
    if count > 1 {
        return errors.Errorf("invalid {{.Name}}: only one field must be set, found %d", count)
    }

    return nil
`,
    }
    structType.Methods = append(structType.Methods, validateMethod)
}

func (g *Generator) addAllOfMarshaling(structType *codegen.StructType) {
    // Add UnmarshalJSON method
    unmarshalMethod := &codegen.Method{
        Name: "UnmarshalJSON",
        Receiver: codegen.Receiver{
            Name: "x",
            Type: &codegen.PointerType{Elem: structType},
        },
        Params: []*codegen.Field{
            {Name: "data", Type: &codegen.ArrayType{ElementType: &codegen.ByteType{}}},
        },
        Returns: []*codegen.Field{
            {Type: &codegen.ErrorType{}},
        },
        Body: `
    type Alias {{.Name}}
    aux := &struct{ *Alias }{Alias: (*Alias)(x)}
    if err := json.Unmarshal(data, &aux); err != nil {
        return errors.Errorf("unmarshaling allOf fields: %w", err)
    }
    return x.Validate()
`,
    }

    // Add MarshalJSON method
    marshalMethod := &codegen.Method{
        Name: "MarshalJSON",
        Receiver: codegen.Receiver{
            Name: "x",
            Type: structType,
        },
        Returns: []*codegen.Field{
            {Type: &codegen.ArrayType{ElementType: &codegen.ByteType{}}},
            {Type: &codegen.ErrorType{}},
        },
        Body: `
    type Alias {{.Name}}
    return json.Marshal((*Alias)(&x))
`,
    }

    structType.Methods = append(structType.Methods, unmarshalMethod, marshalMethod)
}
```

## 🧪 Testing Strategy

Tests must cover all the different combination types with various scenarios:

```go
func TestAllOfGeneration(t *testing.T) {
    input := `{
        "title": "AllOfExample",
        "allOf": [
            { "properties": { "name": { "type": "string" } } },
            { "properties": { "age": { "type": "integer" } } }
        ]
    }`

    schema, err := parser.Parse(input)
    require.NoError(t, err)

    gen := NewGenerator()
    code, err := gen.Generate(schema)
    require.NoError(t, err)

    // Verify output contains expected structs and fields
    assert.Contains(t, code, "type AllOfExample struct {")
    assert.Contains(t, code, "Name_AllOf *string `json:\"name,omitempty\"`")
    assert.Contains(t, code, "Age_AllOf *int `json:\"age,omitempty\"`")

    // Check marshaling methods are included
    assert.Contains(t, code, "func (x *AllOfExample) UnmarshalJSON(data []byte) error {")
    assert.Contains(t, code, "func (x AllOfExample) MarshalJSON() ([]byte, error) {")
}

func TestOneOfGeneration(t *testing.T) {
    input := `{
        "title": "OneOfExample",
        "properties": {
            "identifier": {
                "oneOf": [
                    { "type": "string" },
                    { "type": "integer" }
                ]
            }
        }
    }`

    schema, err := parser.Parse(input)
    require.NoError(t, err)

    gen := NewGenerator()
    code, err := gen.Generate(schema)
    require.NoError(t, err)

    // Verify output contains expected types
    assert.Contains(t, code, "type OneOfExample struct {")
    assert.Contains(t, code, "Identifier OneOfExample_Identifier_OneOf `json:\"identifier,omitempty\"`")
    assert.Contains(t, code, "type OneOfExample_Identifier_OneOf struct {")
    assert.Contains(t, code, "StringValue_OneOf *string `json:\"string_value,omitempty\"`")
    assert.Contains(t, code, "IntegerValue_OneOf *int `json:\"integer_value,omitempty\"`")

    // Check validation is included
    assert.Contains(t, code, "func (x OneOfExample_Identifier_OneOf) Validate() error {")
}

func TestAnyOfGeneration(t *testing.T) {
    input := `{
        "title": "AnyOfExample",
        "properties": {
            "value": {
                "anyOf": [
                    { "type": "string" },
                    { "type": "number" },
                    { "type": "boolean" }
                ]
            }
        }
    }`

    schema, err := parser.Parse(input)
    require.NoError(t, err)

    gen := NewGenerator()
    code, err := gen.Generate(schema)
    require.NoError(t, err)

    // Verify output contains expected types
    assert.Contains(t, code, "type AnyOfExample struct {")
    assert.Contains(t, code, "Value AnyOfExample_Value_AnyOf `json:\"value,omitempty\"`")
    assert.Contains(t, code, "type AnyOfExample_Value_AnyOf struct {")
    assert.Contains(t, code, "StringValue_AnyOf *string `json:\"string_value,omitempty\"`")
    assert.Contains(t, code, "NumberValue_AnyOf *float64 `json:\"number_value,omitempty\"`")
    assert.Contains(t, code, "BooleanValue_AnyOf *bool `json:\"boolean_value,omitempty\"`")

    // Check validation is included
    assert.Contains(t, code, "func (x AnyOfExample_Value_AnyOf) Validate() error {")
}

func TestPatternProperties(t *testing.T) {
    input := `{
        "title": "DynamicConfig",
        "patternProperties": {
            "^S_": { "type": "string" },
            "^N_": { "type": "number" }
        },
        "additionalProperties": false
    }`

    schema, err := parser.Parse(input)
    require.NoError(t, err)

    gen := NewGenerator()
    code, err := gen.Generate(schema)
    require.NoError(t, err)

    // Verify output contains expected pattern fields
    assert.Contains(t, code, "type DynamicConfig struct {")
    assert.Contains(t, code, "StringFields_Pattern map[string]string `json:\"-\"`")
    assert.Contains(t, code, "NumberFields_Pattern map[string]float64 `json:\"-\"`")

    // Check custom marshaling
    assert.Contains(t, code, "func (d *DynamicConfig) UnmarshalJSON(data []byte) error {")
    assert.Contains(t, code, "func (d DynamicConfig) MarshalJSON() ([]byte, error) {")
}
```

## 📋 Implementation Checklist

1. ✅ Name scoping system

    - [ ] Implement NameScope struct and methods
    - [ ] Add support for nested types
    - [ ] Ensure consistent naming conventions

2. ✅ AllOf implementation

    - [ ] Property merging from multiple schemas
    - [ ] Field naming with \_AllOf suffix
    - [ ] Custom marshaling methods

3. ✅ OneOf implementation

    - [ ] Generate wrapper struct for each oneOf field
    - [ ] Add type-specific value fields
    - [ ] Add validation for exactly-one field rule

4. ✅ AnyOf implementation

    - [ ] Generate wrapper struct for each anyOf field
    - [ ] Add type-specific value fields
    - [ ] Add validation for at-least-one field rule

5. ✅ Pattern properties support

    - [ ] Generate map fields for patterns
    - [ ] Add custom marshaling for pattern fields
    - [ ] Add pattern validation

6. ✅ Array combinations support

    - [ ] Handle arrays with oneOf/anyOf/allOf items
    - [ ] Generate appropriate item types

7. ✅ Testing
    - [ ] Unit tests for each combination type
    - [ ] Integration tests with complex schemas
    - [ ] Edge cases and error handling

## 🚀 Success Metrics

1. All test cases pass with expected output
2. Generated code follows naming conventions exactly as specified
3. Marshaling/unmarshaling works correctly for all types
4. Validation enforces the correct constraints
5. Error handling provides clear context about validation failures
