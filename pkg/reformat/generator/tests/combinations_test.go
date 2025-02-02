package tests_test

import (
	"testing"
)

// 🧪 Test Plan:
// 1. Basic anyOf test
// 2. Basic allOf test
// 3. Basic oneOf test
// 4. Complex nested combinations
// 5. References within combinations

func TestAnyOf(t *testing.T) {
	t.Parallel()

	cfg := basicConfig
	testExampleFile(t, cfg, "../testdata/combinations/anyOf/anyOf.json")
}

func TestAllOf(t *testing.T) {
	t.Parallel()

	cfg := basicConfig
	testExampleFile(t, cfg, "../testdata/combinations/allOf/allOf.json")
}

func TestOneOf(t *testing.T) {
	t.Parallel()

	cfg := basicConfig
	testExampleFile(t, cfg, "../testdata/combinations/oneOf/oneOf.json")
}

func TestNestedCombinations(t *testing.T) {
	t.Parallel()

	cfg := basicConfig
	testExampleFile(t, cfg, "../testdata/combinations/nested/nested.json")
}

func TestReferencesInCombinations(t *testing.T) {
	t.Parallel()

	cfg := basicConfig
	testExampleFile(t, cfg, "../testdata/combinations/references/references.json")
}
