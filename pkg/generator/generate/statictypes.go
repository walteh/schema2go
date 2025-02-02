package main

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"
	"os"
	"strings"
	"text/template"

	_ "embed"
)

//go:embed templates/statictypes.go.tmpl
var staticTypesTemplate string

type Method struct {
	Name       string
	ReturnType string
}

type Interface struct {
	Name    string
	Methods []Method
}

type TemplateData struct {
	Interfaces []Interface
}

func generateStaticTypes() {
	// Parse the source file
	fset := token.NewFileSet()
	content, err := os.ReadFile("../types.go")
	if err != nil {
		panic(fmt.Errorf("reading types.go: %w", err))
	}

	node, err := parser.ParseFile(fset, "../types.go", content, 0)
	if err != nil {
		panic(fmt.Errorf("parsing types.go: %w", err))
	}

	// Extract interfaces
	var interfaces []Interface
	for _, decl := range node.Decls {
		if genDecl, ok := decl.(*ast.GenDecl); ok {
			for _, spec := range genDecl.Specs {
				if typeSpec, ok := spec.(*ast.TypeSpec); ok {
					if interfaceType, ok := typeSpec.Type.(*ast.InterfaceType); ok {
						iface := Interface{
							Name: typeSpec.Name.Name,
						}

						// Extract methods
						for _, method := range interfaceType.Methods.List {
							funcType := method.Type.(*ast.FuncType)

							// Get return type as string
							var returnType string
							if len(funcType.Results.List) > 0 {
								returnType = string(content[funcType.Results.List[0].Type.Pos()-1 : funcType.Results.List[0].Type.End()-1])
							}

							iface.Methods = append(iface.Methods, Method{
								Name:       method.Names[0].Name,
								ReturnType: returnType,
							})
						}

						interfaces = append(interfaces, iface)
					}
				}
			}
		}
	}

	// Create a map of interface names for the template
	interfaceNames := make(map[string]bool)
	for _, iface := range interfaces {
		interfaceNames[iface.Name] = true
	}

	// Create template functions
	funcMap := template.FuncMap{
		"isKnownType": func(typeName string) bool {
			return interfaceNames[typeName]
		},
		"getSliceBaseType": func(typeName string) string {
			if strings.HasPrefix(typeName, "[]") {
				return strings.TrimPrefix(typeName, "[]")
			}
			return ""
		},
		"isSliceType": func(typeName string) bool {
			return strings.HasPrefix(typeName, "[]")
		},
		"trimPrefix": func(s, prefix string) string {
			return strings.TrimPrefix(s, prefix)
		},
		"slice": func(s string, i, j int) string {
			if j > len(s) {
				j = len(s)
			}
			return s[i:j]
		},
		"hasParentField": func(returnType, parentType string) bool {
			fmt.Printf("returnType: %s\n", returnType)
			fmt.Printf("parentType: %s\n", parentType)
			// Check if the type has a Parent field by looking for it in the AST
			for _, decl := range node.Decls {
				if genDecl, ok := decl.(*ast.GenDecl); ok {
					for _, spec := range genDecl.Specs {
						if typeSpec, ok := spec.(*ast.TypeSpec); ok {
							fmt.Printf("typeSpec.Name.Name: %s\n", typeSpec.Name.Name)
							if structType, ok := typeSpec.Type.(*ast.StructType); ok {
								// If this is the type we're looking for
								if "[]"+typeSpec.Name.Name == returnType || "[]*"+typeSpec.Name.Name == returnType {
									// Look for a Parent field
									fmt.Printf("structType.Fields.List: %v\n", structType.Fields.List)
									for _, field := range structType.Fields.List {
										fmt.Printf("field.Names: %v\n", field.Names)
										if len(field.Names) > 0 && field.Names[0].Name == "Parent" {
											// Check if the field type matches our parent type
											if ident, ok := field.Type.(*ast.Ident); ok {
												fmt.Printf("ident.Name == parentType: %v\n", ident.Name == parentType)
												fmt.Printf("ident.Name: %s\n", ident.Name)
												return ident.Name == parentType
											}
										}
									}
								}
							}
						}
					}
				}
			}
			return false
		},
	}

	// Generate code using template
	tmpl, err := template.New("staticTypes").Funcs(funcMap).Parse(staticTypesTemplate)
	if err != nil {
		panic(fmt.Errorf("parsing template: %w", err))
	}

	var buf bytes.Buffer
	err = tmpl.Execute(&buf, TemplateData{Interfaces: interfaces})
	if err != nil {
		panic(fmt.Errorf("executing template: %w", err))
	}

	// Format the generated code
	formattedCode, err := format.Source(buf.Bytes())
	if err != nil {
		fmt.Println("Failed to format code:", err)
		fmt.Println("Raw code:", buf.String())
		panic(err)
	}

	// Write the generated code
	if err := os.WriteFile("../statictypes.gen.go", formattedCode, 0644); err != nil {
		panic(fmt.Errorf("writing generated file: %w", err))
	}
}
