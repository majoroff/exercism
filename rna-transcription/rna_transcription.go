package strand

// ToRNA :
// Given a DNA strand, return its RNA complement (per RNA transcription).
func ToRNA(dna string) string {
	var result []rune
	var trans = map[rune]rune{
		'G': 'C',
		'C': 'G',
		'T': 'A',
		'A': 'U',
	}

	if dna == "" {
		return ""
	}

	for _, s := range dna {
		t, ok := trans[s]
		if ok {
			result = append(result, t)
		} else {
			return ""
		}
	}
	return string(result)
}
