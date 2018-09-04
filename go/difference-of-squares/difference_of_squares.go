package diffsquares

func SquareOfSums(n int) (square int) {
	for index := 0; index <= n; index++ {
		square += index
	}

	return square * square
}

func SumOfSquares(n int) (sum int) {
	for index := 0; index <= n; index++ {
		sum += index * index
	}

	return sum
}

func Difference(n int) int {
	square := SquareOfSums(n)
	sum := SumOfSquares(n)

	return square - sum
}
