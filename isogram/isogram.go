package isogram

import (
	"strings"
	"unicode"
)

// IsIsogram determines if a word or phrase is an isogram.
func IsIsogram(s string) bool {
	if len(s) == 0 {
		return true
	}
	s = strings.ToLower(s)
	for i := range s {
		for j := i + 1; j < len(s); j++ {
			if s[i] == s[j] && unicode.IsLetter(rune(s[i])) {
				return false
			}
		}
	}
	return true
}
