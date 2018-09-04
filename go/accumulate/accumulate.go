package accumulate

type fn func(string) string

// Accumulate accumulates
func Accumulate(collection []string, accum fn) []string {
	accumulator := make([]string, len(collection))

	for index, element := range collection {
		accumulator[index] = accum(element)
	}

	return accumulator
}
