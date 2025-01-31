package main

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"
	"os"
	"text/template"
)

// TODO: Move this to a template file later for better maintainability
const staticTypesTemplate = `// Code generated by statictypes.go; DO NOT EDIT.

package generator

{{ range $iface := .Interfaces }}
var _ {{ $iface.Name }} = &Static{{ $iface.Name }}{}

type Static{{ $iface.Name }} struct {
	{{- range $method := $iface.Methods }}
	{{ $method.Name }}_ {{ $method.ReturnType }}
	{{- end }}
}

{{ range $method := $iface.Methods }}
// {{ $method.Name }} implements {{ $iface.Name }}.
func (b *Static{{ $iface.Name }}) {{ $method.Name }}() {{ $method.ReturnType }} {
	return b.{{ $method.Name }}_
}
{{ end }}
{{ end }}

{{ range $iface := .Interfaces }}
func NewStatic{{ $iface.Name }}(impl {{ $iface.Name }}) *Static{{ $iface.Name }} {
	return &Static{{ $iface.Name }}{
		{{- range $method := $iface.Methods }}
		{{ $method.Name }}_: impl.{{ $method.Name }}(),
		{{- end }}
	}
}
{{ end }}
`

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

	// Generate code using template
	tmpl, err := template.New("staticTypes").Parse(staticTypesTemplate)
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
