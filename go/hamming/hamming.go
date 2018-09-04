package hamming

import (
	"errors"
	"strings"
)

func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return -1, errors.New("-1")
	}

	distance := 0

	for index := range a {
		if strings.Compare(string(a[index]), string(b[index])) != 0 {
			distance++
		}
	}

	return distance, nil
}
