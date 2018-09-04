package raindrops

import "strconv"

/**
# Raindrops

Convert a number to a string, the contents of which depend on the number's factors.

- If the number has 3 as a factor, output 'Pling'.
- If the number has 5 as a factor, output 'Plang'.
- If the number has 7 as a factor, output 'Plong'.
- If the number does not have 3, 5, or 7 as a factor,
  just pass the number's digits straight through.
**/
func Convert(number int) string {
	factors := []int{3, 5, 7}
	raindrops := []string{"Pling", "Plang", "Plong"}

	var converted string

	for index, factor := range factors {
		if number%factor == 0 {
			converted += raindrops[index]
		}
	}

	if converted == "" {
		converted = strconv.Itoa(number)
	}

	return converted
}
