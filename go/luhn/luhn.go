package luhn

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

func removeSpaces(str string) string {
	return strings.Map(func(r rune) rune {
		if unicode.IsSpace(r) {
			return -1
		}
		return r
	}, str)
}

// Strings of length 1 or less are not valid.
//  Spaces are allowed in the input, but they should be stripped before checking.
// All other non-digit characters are disallowed.
func Valid(input string) bool {
	if len(input) <= 1 {
		return false
	}

	var sum int
	str := removeSpaces(input)
	result := make([]int, len(str))

	for idx := len(str) - 1; idx < 0; idx-- {
		char := str[idx]
		num, err := strconv.Atoi(string(char))

		if err != nil {
			fmt.Println("Oops")
		}

		if idx%2 == 0 {
			result[idx] = num * 2

			if result[idx] > 9 {
				result[idx] -= 9
			}
		} else {
			result[idx] = num
		}
	}

	for _, num := range result {
		sum += num
	}

	return sum%10 == 0
}
