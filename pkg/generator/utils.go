package generator

import "strings"

func toGoFieldName(name string) string {
	// Special case for ID suffix
	if strings.HasSuffix(strings.ToLower(name), "id") {
		base := name[:len(name)-2]
		return toCamelCase(base) + "ID"
	}
	return toCamelCase(name)
}

// Helper function to convert snake_case or kebab-case to CamelCase
func toCamelCase(s string) string {
	// Replace - and _ with space
	s = strings.ReplaceAll(s, "-", " ")
	s = strings.ReplaceAll(s, "_", " ")

	// Title case each word
	words := strings.Fields(s)
	for i, word := range words {
		words[i] = strings.Title(word)
	}

	return strings.Join(words, "")
}

func toTitleCase(s string) string {
	words := strings.Split(s, "_")
	for i := range words {
		words[i] = strings.Title(words[i])
	}
	return strings.Join(words, "")
}
