package isogram

import (
	"strings"
)

// IsIsogram checks if a word is an Isogram
func IsIsogram(word string) bool {
	runes := make(map[rune]struct{})
	isIsogram := true

	for _, letter := range strings.ToLower(word) {
		if _, ok := runes[letter]; ok {
			// 45 = dash and 32 = space
			isIsogram = (letter == 45 || letter == 32)
			break
		}

		runes[letter] = struct{}{}
	}

	return isIsogram
}
