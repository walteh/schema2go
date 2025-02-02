copy {
	source {
		repo     = "github.com/omissis/go-jsonschema"
		ref      = "main"
		ref_type = "branch"
		path     = "pkg/generator"
	}
	destination {
		path = "pkg/reformat/generator"
	}
	options {
		file_patterns = [
			"**/*.go",
		]
		replacements = [
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
		file_patterns = [
			"**/*.go",
		]
		replacements = [
		]
	}
}

copy {
	source {
		repo     = "github.com/omissis/go-jsonschema"
		ref      = "main"
		ref_type = "branch"
		path     = "tests/data"
	}
	destination {
		path = "pkg/reformat/generator/testdata"
	}
	options {
		no_header_comments = true
		recursive          = true
		file_patterns = [
			"**/*.go",
			"**/*.json",
			"**/*.yaml",
		]
		replacements = [

		]
	}
}

copy {
	source {
		repo     = "github.com/omissis/go-jsonschema"
		ref      = "main"
		ref_type = "branch"
		path     = "tests"
	}
	destination {
		path = "pkg/reformat/generator/tests"
	}
	options {
		recursive = false
		file_patterns = [
			"**/generation_test.go",
		]
		replacements = [
			{
				old = "./data"
				new = "../testdata"
			},
			{
				old = "github.com/atombender/go-jsonschema/pkg/generator"
				new = "github.com/walteh/schema2go/pkg/reformat/generator"
			}
		]
	}
}