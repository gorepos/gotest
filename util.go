package main

import (
	"strings"
	"unicode"
)

// convertCamelCase converts a camelCase string to a human-readable string.
func convertCamelCase(input string) string {
	var result strings.Builder
	var lastIsUpper bool
	for i, r := range input {
		nextIsLower := i+1 < len(input) && unicode.IsLower(rune(input[i+1]))
		//nextIsUpper := i+1 < len(input) && unicode.IsUpper(rune(input[i+1]))

		if unicode.IsUpper(r) {
			// Add space if it's not the first letter
			// and the previous letter is not upper case, or it's the start of an abbreviation.
			if i > 0 && (!lastIsUpper || nextIsLower) {
				result.WriteRune(' ')
			}
			lastIsUpper = true
		} else {
			lastIsUpper = false
		}

		if i == 0 {
			// Capitalize the first letter
			r = unicode.ToUpper(r)
		} else {
			// Lowercase the rest
			r = unicode.ToLower(r)
		}

		result.WriteRune(r)
	}

	// Replace underscores with spaces
	str := strings.Replace(result.String(), "_", " ", -1)

	// Remove double spaces
	for strings.Contains(str, "  ") {
		str = strings.Replace(str, "  ", " ", -1)
	}

	return str
}
