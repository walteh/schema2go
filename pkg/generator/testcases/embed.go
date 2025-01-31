package testcases

import (
	"crypto/sha256"
	"embed"
	"encoding/hex"
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
	name         string
	jsonSchema   string
	goCode       string
	staticSchema string
	rawSchema    string
}

func (tc *TestCase) JSONSchema() string {
	return tc.jsonSchema
}

func (tc *TestCase) GoCode() string {
	return tc.goCode
}

func (tc *TestCase) StaticSchema() string {
	return tc.staticSchema
}

func (tc *TestCase) RawSchema() string {
	return tc.rawSchema
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
	var jsonSchema, goCode, staticSchema, rawSchema string

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
		case "json-schema":
			jsonSchema = parseJson(blockContent)
		case "go-code":
			goCode = parseGo(blockContent)
		case "static-schema":
			staticSchema = parseGo(blockContent)
		case "raw-schema":
			rawSchema = parseGo(blockContent)
		}
	}

	return TestCase{
		name:         name,
		jsonSchema:   jsonSchema,
		goCode:       goCode,
		staticSchema: staticSchema,
		rawSchema:    rawSchema,
	}
}

func GetHash() string {
	hash := sha256.New()
	files, err := testCases.ReadDir(".")
	if err != nil {
		panic(err)
	}
	for _, file := range files {
		content, err := testCases.ReadFile(filepath.Join(file.Name()))
		if err != nil {
			panic(err)
		}
		hash.Write(content)
	}
	return hex.EncodeToString(hash.Sum(nil))
}
