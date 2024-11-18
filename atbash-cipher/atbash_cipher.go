package atbash

import "unicode"

func Atbash(s string) string {
	ret := ""
	k := 0
	chunk_size := 5
	for _, r := range s {
		if unicode.IsSpace(r) || unicode.IsPunct(r) {
			continue
		}
		if k > 0 && k%chunk_size == 0 {
			ret += " "
		}
		switch {
		case unicode.IsLetter(r):
			ret += string('z' - unicode.ToLower(r) + 'a')
		default:
			ret += string(r)
		}
		k++
	}
	return ret
}
