package romannumerals

var ConversionError error

var RomanNumerals = map[int]string{
	1:    "I",
	5:    "V",
	10:   "X",
	50:   "L",
	100:  "C",
	500:  "D",
	1000: "M",
}

func ToRomanNumeral(number int) (string, error) {
	var romanNumeral string

	if number < 1 || number > 3000 {
		return "", ConversionError
	}

	return romanNumeral, ConversionError
}
