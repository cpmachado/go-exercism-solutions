package isogram

import (
	"strings"
	"unicode"
)

func IsIsogram(word string) bool {
	lowered := strings.ToLower(word)
	for i, letter := range lowered {
		if unicode.IsLetter(letter) {
			for _, other := range lowered[i+1:] {
				if letter == other {
					return false
				}
			}
		}
	}
	return true
}
