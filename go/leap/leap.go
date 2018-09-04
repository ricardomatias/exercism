package leap

// IsLeapYear should have a comment documenting it.
func IsLeapYear(year int) bool {
	evenBy4 := year%4 == 0

	if evenBy4 && year%100 == 0 {
		return year%400 == 0
	}

	return evenBy4
}
