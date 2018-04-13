package scrabble

import (
	"strings"
	"unicode"
)

func Keys(m map[string]int) (keys []string) {
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}

func Score(word string) int {
	letterScore := map[string]int{
		"AEIOULNRST": 1,
		"DG":         2,
		"BCMP":       3,
		"FHVWY":      4,
		"K":          5,
		"JX":         8,
		"QZ":         10,
	}
	score := 0
	for _, c := range word {
		for _, key := range Keys(letterScore) {
			if strings.ContainsAny(key, string(unicode.ToUpper(c))) {
				score += letterScore[key]
			}
		}
	}
	return score
}
