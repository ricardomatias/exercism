package etl

import (
	"strings"
)

// Transform a scrabble score
func Transform(oldScore map[int][]string) map[string]int {
	transformedScore := make(map[string]int)

	for score, letters := range oldScore {
		for _, char := range letters {
			transformedScore[strings.ToLower(char)] = score
		}

	}

	return transformedScore
}
