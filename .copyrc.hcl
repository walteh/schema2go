
copy {
	source {
		repo     = "github.com/yassinebenaid/godump"
		ref      = "master"
		ref_type = "branch"
		path     = "."
	}
	destination {
		path = "pkg/godump"
	}
	options {
		file_patterns = ["**/*.go"]
		replacements = [
			{
				old = "github.com/yassinebenaid/godump"
				new = "github.com/walteh/schema2go/pkg/godump"
			}
		]
	}
}