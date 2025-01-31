package testcases

import (
	"embed"
	"path/filepath"
	"strings"
)

//go:embed *.md
var testCases embed.FS

func LoadTestCases() ([]TestCase, error) {
	files, err := testCases.ReadDir(".")
	if err != nil {
		return nil, err
	}
	cases := []TestCase{}
	for _, file := range files {
		if strings.HasSuffix(file.Name(), ".md") {
			cases = append(cases, LoadAndParseTestCase(strings.TrimSuffix(file.Name(), ".md")))
		}
	}
	return cases, nil
}

// TestCase represents a single schema to struct conversion test
type TestCase struct {
	name                      string
	input                     string
	wantGoCode                string
	wantSchemaModelMockGoCode string
	schema                    string
}

func (tc *TestCase) GenerateInput() string {
	return tc.input
}

func (tc *TestCase) GenerateExpectedGoCode() string {
	return tc.wantGoCode
}

func (tc *TestCase) GenerateExpectedSchemaModelMockGoCode() string {
	return tc.wantSchemaModelMockGoCode
}

func (tc *TestCase) GenerateExpectedSchemaModel() string {
	return tc.schema
}

func (tc *TestCase) Name() string {
	return tc.name
}

func LoadAndParseTestCase(name string) TestCase {
	content, err := testCases.ReadFile(filepath.Join(name + ".md"))
	if err != nil {
		panic(err)
	}
	return parseTestCase(name, string(content))
}

func parseTestCase(name string, text string) TestCase {
	// Split the markdown into sections
	// We expect:
	// ```json
	// <input schema>
	// ```
	//
	// ```go
	// <expected output>
	// ```
	var input, wantGoCode, wantSchemaModelMockGoCode, schema string

	parseJson := func(text string) string {
		jsonStart := strings.Index(text, "```json\n")
		jsonEnd := strings.Index(text[jsonStart+7:], "\n```")
		if jsonStart == -1 || jsonEnd == -1 {
			panic("Could not find JSON section in test case markdown: " + text)
		}
		return strings.TrimSpace(text[jsonStart+7 : jsonStart+7+jsonEnd])
	}

	parseGo := func(text string) string {
		goStart := strings.Index(text, "```go\n")
		goEnd := strings.Index(text[goStart+5:], "\n```")
		if goStart == -1 || goEnd == -1 {
			panic("Could not find Go section in test case markdown: " + text)
		}
		what := strings.TrimSpace(text[goStart+5 : goStart+5+goEnd])
		return what
	}

	splits := strings.Split(text, "\n---\n")
	for _, split := range splits {
		// grab the name of the block (will be in a # title)
		blockName := strings.Split(split[strings.Index(split, "# ")+2:], "\n")[0]
		blockContent := strings.TrimSpace(split[strings.Index(split, blockName)+len(blockName):])
		switch blockName {
		case "input":
			input = parseJson(blockContent)
		case "expected-output":
			wantGoCode = parseGo(blockContent)
		case "expected-schema-model-mock":
			wantSchemaModelMockGoCode = parseGo(blockContent)
		case "expected-raw-schema":
			schema = parseGo(blockContent)
		}
	}

	return TestCase{
		name:                      name,
		input:                     input,
		wantGoCode:                wantGoCode,
		wantSchemaModelMockGoCode: wantSchemaModelMockGoCode,
		schema:                    schema,
	}
}
