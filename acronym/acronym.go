package acronym

import (
	"strings"
	"unicode"
)

// Abbreviate converts a phrase to its acronym
// e.g. Portable Network Graphics to PNG
func Abbreviate(s string) string {
	var words []string
	var result strings.Builder

	f := func(r rune) bool {
		return unicode.IsSpace(r) || r == '-'
	}

	words = strings.FieldsFunc(s, f)
	for _, w := range words {
		for _, l := range w {
			if !unicode.IsLetter(l) {
				continue
			}
			result.WriteString(strings.ToUpper(string(l)))
			break
		}
	}

	return result.String()
}
