package scrabble

import (
	"strings"
	"unicode"
)

// Score computes the Scrabble score for given word
func Score(s string) int {
	var scores = map[string]int{
		"AEIOULNRST": 1,
		"DG":         2,
		"BCMP":       3,
		"FHVWY":      4,
		"K":          5,
		"JX":         8,
		"QZ":         10,
	}
	var result int

	for _, char := range s {
		for key, value := range scores {
			if strings.ContainsRune(key, unicode.ToUpper(char)) {
				result += value
			}
		}
	}
	return result
}
