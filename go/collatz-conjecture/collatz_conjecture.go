package collatzconjecture

import (
	"errors"
)

var count int

// CollatzConjecture Takes any positive integer n. If n is even, divide n by 2 to get n / 2. If n is
// odd, multiply n by 3 and add 1 to get 3n + 1. Repeat the process indefinitely.
// The conjecture states that no matter which number you start with, you will
// always reach 1 eventually.
func CollatzConjecture(number int) (int, error) {
	if number <= 0 {
		return -1, errors.New("Number cannot be lower than 1")
	}

	if number > 1 {
		count++

		if number%2 == 0 {
			return CollatzConjecture(number / 2)
		}

		return CollatzConjecture(number*3 + 1)
	}

	result := count

	count = 0

	return result, nil
}
