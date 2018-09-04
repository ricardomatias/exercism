package proverb

import "fmt"

// Proverb returns a proverb
func Proverb(rhyme []string) (proverb []string) {
	if len(rhyme) == 0 {
		return proverb
	}

	for index := 0; index < len(rhyme)-1; index++ {
		a, b := rhyme[index], rhyme[index+1]

		proverb = append(proverb, fmt.Sprintf("For want of a %s the %s was lost.", a, b))
	}

	proverb = append(proverb, fmt.Sprintf("And all for the want of a %s.", rhyme[0]))

	return proverb
}
