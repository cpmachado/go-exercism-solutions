package pangram

import "strings"

func IsPangram(input string) bool {
	if len(input) < 26 {
		return false
	}
	lowered := strings.ToLower(input)

	for ch := 'a'; ch <= 'z'; ch++ {
		if !strings.ContainsRune(lowered, ch) {
			return false
		}
	}

	return true
}
