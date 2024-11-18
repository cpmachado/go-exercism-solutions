package rotationalcipher

import "unicode"

func RotationalCipher(plain string, shiftKey int) string {
	ret := ""

	for _, ch := range plain {
		switch {
		case unicode.IsUpper(ch):
			ch = rune((int(ch-'A')+shiftKey)%26) + 'A'
		case unicode.IsLower(ch):
			ch = rune((int(ch-'a')+shiftKey)%26) + 'a'
		}
		ret += string(ch)
	}
	return ret
}
