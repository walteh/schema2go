// 🔄 Combination Types
// This file handles the generation of anyOf, allOf, and oneOf schema combinations

package generator

import (
	"errors"
	"fmt"
	"strings"

	"github.com/atombender/go-jsonschema/pkg/schemas"
	"github.com/walteh/schema2go/pkg/reformat/codegen"
)

// generateCombinationType handles anyOf, allOf, and oneOf schema combinations
func (g *schemaGenerator) generateCombinationType(t *schemas.Type, scope nameScope) (codegen.Type, error) {
	// Handle allOf - merge all properties into a single type
	if len(t.AllOf) > 0 {
		structType, err := g.generateAllOfType(t.AllOf, scope)
		if err != nil {
			return nil, err
		}

		// Add validation
		g.addValidationMethods(&codegen.TypeDecl{
			Name: scope.string(),
			Type: structType,
		}, []validator{
			&allOfValidator{fields: structType.(*codegen.StructType).Fields},
		})

		return structType, nil
	}

	// Handle anyOf - create a type that can be any of the given types
	if len(t.AnyOf) > 0 {
		structType, err := g.generateAnyOfType(t.AnyOf, scope)
		if err != nil {
			return nil, err
		}

		// Add validation
		g.addValidationMethods(&codegen.TypeDecl{
			Name: scope.string(),
			Type: structType,
		}, []validator{
			&anyOfValidator{fields: structType.(*codegen.StructType).Fields},
		})

		return structType, nil
	}

	// Handle oneOf - create a type that must be exactly one of the given types
	if len(t.OneOf) > 0 {
		structType, err := g.generateOneOfType(t.OneOf, scope)
		if err != nil {
			return nil, err
		}

		// Add validation
		g.addValidationMethods(&codegen.TypeDecl{
			Name: scope.string(),
			Type: structType,
		}, []validator{
			&oneOfValidator{fields: structType.(*codegen.StructType).Fields},
		})

		return structType, nil
	}

	return nil, errors.New("no combination type found")
}

// generateAllOfType merges all properties from the given types into a single type
func (g *schemaGenerator) generateAllOfType(types []*schemas.Type, scope nameScope) (codegen.Type, error) {
	// Create a struct type to hold all properties
	structType := &codegen.StructType{
		Fields: make([]codegen.StructField, 0),
	}

	// Track all required fields
	requiredFields := make(map[string]bool)

	// Add fields from each type
	for _, t := range types {
		// Add required fields to tracking
		for _, req := range t.Required {
			requiredFields[req] = true
		}

		// Add properties
		for name, prop := range t.Properties {
			fieldName := g.caser.Identifierize(name)

			// If the property is an enum, generate it as a separate type
			if len(prop.Enum) > 0 {
				enumType, err := g.generateEnumType(prop, scope.add(fieldName))
				if err != nil {
					return nil, fmt.Errorf("generating enum type for field %s: %w", name, err)
				}
				fieldType := enumType
				if !requiredFields[name] {
					fieldType = codegen.WrapTypeInPointer(fieldType)
				}
				field := codegen.StructField{
					Name:     fieldName,
					Type:     fieldType,
					Tags:     fmt.Sprintf(`json:"%s%s"`, name, g.getOmitEmptyTag(!requiredFields[name])),
					JSONName: name,
				}
				structType.Fields = append(structType.Fields, field)
				continue
			}

			// Handle regular properties
			fieldType, err := g.generateType(prop, scope.add(fieldName))
			if err != nil {
				return nil, fmt.Errorf("generating allOf field %s: %w", name, err)
			}

			// Make the field a pointer if not required
			if !requiredFields[name] {
				fieldType = codegen.WrapTypeInPointer(fieldType)
			}

			field := codegen.StructField{
				Name:     fieldName,
				Type:     fieldType,
				Tags:     fmt.Sprintf(`json:"%s%s"`, name, g.getOmitEmptyTag(!requiredFields[name])),
				JSONName: name,
			}
			structType.Fields = append(structType.Fields, field)
		}
	}

	return structType, nil
}

// Helper function to get the omitempty tag
func (g *schemaGenerator) getOmitEmptyTag(shouldOmit bool) string {
	if shouldOmit {
		return ",omitempty"
	}
	return ""
}

// generateAnyOfType creates a type that can be any of the given types
func (g *schemaGenerator) generateAnyOfType(types []*schemas.Type, scope nameScope) (codegen.Type, error) {
	// Create a struct type with fields for each possible type
	structType := &codegen.StructType{
		Fields: make([]codegen.StructField, 0),
	}

	// Add a field for each possible type
	for i, t := range types {
		fieldName := fmt.Sprintf("Type%d", i+1)
		fieldType, err := g.generateType(t, scope.add(fieldName))
		if err != nil {
			return nil, fmt.Errorf("generating anyOf field %s: %w", fieldName, err)
		}

		// Make the field a pointer since it's optional
		field := codegen.StructField{
			Name:     fieldName,
			Type:     codegen.WrapTypeInPointer(fieldType),
			Tags:     fmt.Sprintf(`json:"%s,omitempty"`, strings.ToLower(fieldName)),
			JSONName: strings.ToLower(fieldName),
		}
		structType.Fields = append(structType.Fields, field)
	}

	return structType, nil
}

// generateOneOfType creates a type that must be exactly one of the given types
func (g *schemaGenerator) generateOneOfType(types []*schemas.Type, scope nameScope) (codegen.Type, error) {
	// Create a struct type with fields for each possible type
	structType := &codegen.StructType{
		Fields: make([]codegen.StructField, 0),
	}

	// Add a field for each possible type
	for i, t := range types {
		fieldName := fmt.Sprintf("Type%d", i+1)
		fieldType, err := g.generateType(t, scope.add(fieldName))
		if err != nil {
			return nil, fmt.Errorf("generating oneOf field %s: %w", fieldName, err)
		}

		// Make the field a pointer since it's optional
		field := codegen.StructField{
			Name:     fieldName,
			Type:     codegen.WrapTypeInPointer(fieldType),
			Tags:     fmt.Sprintf(`json:"%s,omitempty"`, strings.ToLower(fieldName)),
			JSONName: strings.ToLower(fieldName),
		}
		structType.Fields = append(structType.Fields, field)
	}

	return structType, nil
}

// addValidationMethods adds validation methods to the type declaration
func (g *schemaGenerator) addValidationMethods(decl *codegen.TypeDecl, validators []validator) {
	if len(validators) > 0 {
		for _, v := range validators {
			if v.desc().hasError {
				g.output.file.Package.AddImport("fmt", "")
				break
			}
		}

		g.output.file.Package.AddImport("encoding/json", "")

		// Add UnmarshalJSON method
		g.output.file.Package.AddDecl(&codegen.Method{
			Impl: func(out *codegen.Emitter) {
				out.Comment("UnmarshalJSON implements json.Unmarshaler")
				out.Printlnf("func (x *%s) UnmarshalJSON(data []byte) error {", decl.Name)
				out.Indent(1)
				for _, v := range validators {
					v.generate(out)
				}
				out.Indent(-1)
				out.Printlnf("}")
			},
			Name: decl.GetName() + "_unmarshal",
		})

		// Add MarshalJSON method
		g.output.file.Package.AddDecl(&codegen.Method{
			Impl: func(out *codegen.Emitter) {
				out.Comment("MarshalJSON implements json.Marshaler")
				out.Printlnf("func (x %s) MarshalJSON() ([]byte, error) {", decl.Name)
				out.Indent(1)
				for _, v := range validators {
					v.generate(out)
				}
				out.Indent(-1)
				out.Printlnf("}")
			},
			Name: decl.GetName() + "_marshal",
		})
	}
}

// allOfValidator validates allOf combinations
type allOfValidator struct {
	fields []codegen.StructField
}

func (v *allOfValidator) desc() *validatorDesc {
	return &validatorDesc{
		hasError:            true,
		beforeJSONUnmarshal: false,
		requiresRawAfter:    false,
	}
}

func (v *allOfValidator) generate(out *codegen.Emitter) {
	out.Printlnf("var aux struct {")
	for _, f := range v.fields {
		out.Printlnf("\t%s %v `json:\"%s\"`", f.Name, f.Type, f.JSONName)
	}
	out.Printlnf("}")
	out.Printlnf("if err := json.Unmarshal(data, &aux); err != nil {")
	out.Printlnf("\treturn fmt.Errorf(\"unmarshaling allOf fields: %%w\", err)")
	out.Printlnf("}")

	// Copy values and validate required fields
	for _, f := range v.fields {
		out.Printlnf("x.%s = aux.%s", f.Name, f.Name)
		if !strings.Contains(f.Tags, "omitempty") {
			out.Printlnf("if x.%s == nil {", f.Name)
			out.Printlnf("\treturn fmt.Errorf(\"field %s is required\")", f.JSONName)
			out.Printlnf("}")
		}
	}
}

// anyOfValidator validates anyOf combinations
type anyOfValidator struct {
	fields []codegen.StructField
}

func (v *anyOfValidator) desc() *validatorDesc {
	return &validatorDesc{
		hasError:            true,
		beforeJSONUnmarshal: false,
		requiresRawAfter:    true,
	}
}

func (v *anyOfValidator) generate(out *codegen.Emitter) {
	out.Printlnf("var raw json.RawMessage")
	out.Printlnf("if err := json.Unmarshal(data, &raw); err != nil {")
	out.Printlnf("\treturn fmt.Errorf(\"unmarshaling raw message: %%w\", err)")
	out.Printlnf("}")

	for _, f := range v.fields {
		out.Printlnf("var %s %v", strings.ToLower(f.Name), f.Type)
		out.Printlnf("if err := json.Unmarshal(raw, &%s); err == nil {", strings.ToLower(f.Name))
		out.Printlnf("\tx.%s = &%s", f.Name, strings.ToLower(f.Name))
		out.Printlnf("\treturn nil")
		out.Printlnf("}")
	}

	out.Printlnf("return fmt.Errorf(\"value does not match any of the expected types\")")
}

// oneOfValidator validates oneOf combinations
type oneOfValidator struct {
	fields []codegen.StructField
}

func (v *oneOfValidator) desc() *validatorDesc {
	return &validatorDesc{
		hasError:            true,
		beforeJSONUnmarshal: false,
		requiresRawAfter:    true,
	}
}

func (v *oneOfValidator) generate(out *codegen.Emitter) {
	out.Printlnf("var raw json.RawMessage")
	out.Printlnf("if err := json.Unmarshal(data, &raw); err != nil {")
	out.Printlnf("\treturn fmt.Errorf(\"unmarshaling raw message: %%w\", err)")
	out.Printlnf("}")

	out.Printlnf("var validCount int")
	for _, f := range v.fields {
		out.Printlnf("var %s %v", strings.ToLower(f.Name), f.Type)
		out.Printlnf("if err := json.Unmarshal(raw, &%s); err == nil {", strings.ToLower(f.Name))
		out.Printlnf("\tx.%s = &%s", f.Name, strings.ToLower(f.Name))
		out.Printlnf("\tvalidCount++")
		out.Printlnf("}")
	}

	out.Printlnf("if validCount != 1 {")
	out.Printlnf("\treturn fmt.Errorf(\"value must match exactly one of the expected types\")")
	out.Printlnf("}")
	out.Printlnf("return nil")
}

// Helper function to generate allOf alias fields
func (g *schemaGenerator) generateAllOfAliasFields(structType *codegen.StructType) string {
	var fields []string
	for _, field := range structType.Fields {
		fields = append(fields, fmt.Sprintf("%s %v `json:\"%s,omitempty\"`",
			field.Name, field.Type, field.JSONName))
	}
	return strings.Join(fields, "\n\t\t")
}

// Helper function to generate allOf field assignments
func (g *schemaGenerator) generateAllOfFieldAssignments(structType *codegen.StructType) string {
	var assignments []string
	for _, field := range structType.Fields {
		assignments = append(assignments, fmt.Sprintf("x.%s = aux.%s", field.Name, field.Name))
	}
	return strings.Join(assignments, "\n\t")
}

// Helper function to generate anyOf try blocks
func (g *schemaGenerator) generateAnyOfTryBlocks(structType *codegen.StructType) string {
	var blocks []string
	for _, field := range structType.Fields {
		block := fmt.Sprintf(`
	// Try %s
	var %s %v
	if err := json.Unmarshal(raw, &%s); err == nil {
		x.%s = &%s
		return nil
	}`, field.Name, strings.ToLower(field.Name), field.Type,
			strings.ToLower(field.Name), field.Name, strings.ToLower(field.Name))
		blocks = append(blocks, block)
	}
	return strings.Join(blocks, "\n")
}

// Helper function to generate oneOf try blocks
func (g *schemaGenerator) generateOneOfTryBlocks(structType *codegen.StructType) string {
	var blocks []string
	for _, field := range structType.Fields {
		block := fmt.Sprintf(`
	// Try %s
	var %s %v
	if err := json.Unmarshal(raw, &%s); err == nil {
		x.%s = &%s
		validCount++
	}`, field.Name, strings.ToLower(field.Name), field.Type,
			strings.ToLower(field.Name), field.Name, strings.ToLower(field.Name))
		blocks = append(blocks, block)
	}
	return strings.Join(blocks, "\n")
}

// Helper function to generate oneOf marshal body
func (g *schemaGenerator) generateOneOfMarshalBody(structType *codegen.StructType) string {
	var blocks []string
	blocks = append(blocks, `
	var count int
	var value interface{}`)

	for _, field := range structType.Fields {
		block := fmt.Sprintf(`
	if x.%s != nil {
		count++
		value = x.%s
	}`, field.Name, field.Name)
		blocks = append(blocks, block)
	}

	blocks = append(blocks, `
	if count != 1 {
		return nil, errors.Errorf("oneOf requires exactly one value to be set")
	}

	return json.Marshal(value)`)

	return strings.Join(blocks, "\n")
}

// Helper function to generate anyOf marshal body
func (g *schemaGenerator) generateAnyOfMarshalBody(structType *codegen.StructType) string {
	var blocks []string
	for _, field := range structType.Fields {
		block := fmt.Sprintf(`
	if x.%s != nil {
		return json.Marshal(x.%s)
	}`, field.Name, field.Name)
		blocks = append(blocks, block)
	}
	return strings.Join(blocks, "\n")
}

// Helper function to check if a field is required
func isRequired(t *schemas.Type, fieldName string) bool {
	for _, required := range t.Required {
		if required == fieldName {
			return true
		}
	}
	return false
}

// Helper function to check if a type contains a specific type name
func hasType(t *schemas.Type, typeName string) bool {
	for _, typ := range t.Type {
		if string(typ) == typeName {
			return true
		}
	}
	return false
}

// Helper function to generate allOf validation body
func (g *schemaGenerator) generateAllOfValidationBody(structType *codegen.StructType) string {
	var requiredChecks []string
	for _, field := range structType.Fields {
		if !strings.Contains(field.Tags, "omitempty") {
			check := fmt.Sprintf(`
	if x.%s == nil {
		return errors.Errorf("field %s is required")
	}`, field.Name, field.JSONName)
			requiredChecks = append(requiredChecks, check)
		}
	}

	if len(requiredChecks) == 0 {
		return "return nil"
	}

	return strings.Join(requiredChecks, "\n") + "\n\treturn nil"
}
