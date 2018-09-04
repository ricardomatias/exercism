package acronym

import (
	"regexp"
	"strings"
)

func Abbreviate(s string) string {
	acronym := ""

	for _, word := range regexp.MustCompile(`\W`).Split(s, -1) {
		if len(word) > 0 {
			acronym += strings.ToUpper(string(word[0]))
		}
	}

	return acronym
}
