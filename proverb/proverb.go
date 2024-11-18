package proverb

import "fmt"

// Proverb should have a comment documenting it.
func Proverb(rhyme []string) []string {
	n := len(rhyme)
	ret := make([]string, n)
	if n < 1 {
		return ret
	}
	prev := rhyme[0]
	first := prev

	for i, curr := range rhyme[1:] {
		ret[i] = fmt.Sprintf("For want of a %s the %s was lost.", prev, curr)
		prev = curr
	}
	ret[n-1] = fmt.Sprintf("And all for the want of a %s.", first)
	return ret

}
