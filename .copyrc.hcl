copy {
	source {
		repo     = "github.com/omissis/go-jsonschema"
		ref      = "main"
		ref_type = "branch"
		path     = "pkg/generator"
	}
	destination {
		path = "pkg/reformat"
	}
	options {
		replacements = [
			{
				old = "package generator"
				new = "package reformat"
			},
			{
				old = "\"github.com/atombender/go-jsonschema/internal/x/text\""
				new = "\"github.com/walteh/schema2go/pkg/reformat/internal/x/text\""
			}
		]
	}
}

copy {
	source {
		repo     = "github.com/omissis/go-jsonschema"
		ref      = "main"
		ref_type = "branch"
		path     = "internal/x/text"
	}
	destination {
		path = "pkg/reformat/internal/x/text"
	}
	options {
		replacements = [

		]
	}
}