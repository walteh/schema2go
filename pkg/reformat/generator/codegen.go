// 🔄 Code Generation
// This file contains custom code generation types that override the default ones

package generator

import (
	"github.com/atombender/go-jsonschema/pkg/codegen"
)

// CustomFile extends codegen.File to use our own template
type CustomFile struct {
	*codegen.File
}

// Generate overrides the default code generation to use our template
func (f *CustomFile) Generate(out *codegen.Emitter) {
	// Write the header
	out.Printlnf("// Code generated by schema2go. DO NOT EDIT.")
	out.Printlnf("")

	// Write the package declaration
	out.Printlnf("package %s", f.Package.Name())
	out.Printlnf("")

	// Write imports if any
	if len(f.Package.Imports) > 0 {
		for _, imp := range f.Package.Imports {
			if imp.Name != "" {
				out.Printlnf("import %s %q", imp.Name, imp.QualifiedName)
			} else {
				out.Printlnf("import %q", imp.QualifiedName)
			}
		}
		out.Printlnf("")
	}

	// Write declarations
	for _, decl := range f.Package.Decls {
		decl.Generate(out)
		out.Printlnf("")
	}
}

// NewCustomFile creates a new CustomFile
func NewCustomFile(fileName string, pkg codegen.Package) *CustomFile {
	return &CustomFile{
		File: &codegen.File{
			FileName: fileName,
			Package:  pkg,
		},
	}
}

// Helper function to convert a codegen.File to a CustomFile
func ToCustomFile(f *codegen.File) *CustomFile {
	return &CustomFile{File: f}
}
