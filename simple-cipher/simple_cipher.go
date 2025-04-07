package cipher

import (
	"strings"
	"unicode"
)

type (
	shift    int
	vigenere string
)

func NewCaesar() Cipher {
	return NewShift(3)
}

func NewShift(distance int) Cipher {
	if distance <= -26 || distance == 0 || distance >= 26 {
		return nil
	}
	return shift((distance + 26) % 26)
}

func (c shift) Encode(input string) string {
	res := strings.Builder{}

	for _, r := range strings.ToLower(input) {
		if !unicode.IsLower(r) {
			continue
		}
		res.WriteRune((r-'a'+rune(c)+26)%26 + 'a')
	}

	return res.String()
}

func (c shift) Decode(input string) string {
	res := strings.Builder{}

	for _, r := range strings.ToLower(input) {
		res.WriteRune((r-'a'-rune(c)+26)%26 + 'a')
	}

	return res.String()
}

func NewVigenere(key string) Cipher {
	diffThanA := false
	for _, r := range key {
		if r != 'a' {
			diffThanA = true
		}
		if !unicode.IsLower(r) {
			return nil
		}
	}
	if !diffThanA {
		return nil
	}
	return vigenere(key)
}

func (v vigenere) Encode(input string) string {
	res := strings.Builder{}
	n := len(v)
	d := 0

	for i, r := range strings.ToLower(input) {
		k := rune(v[(i-d)%n])
		if !unicode.IsLower(r) {
			d++
			continue
		}
		res.WriteRune((r+k-2*'a'+26)%26 + 'a')
	}

	return res.String()
}

func (v vigenere) Decode(input string) string {
	res := strings.Builder{}
	n := len(v)
	d := 0

	for i, r := range strings.ToLower(input) {
		k := rune(v[(i-d)%n])
		if !unicode.IsLower(r) {
			d++
			continue
		}
		res.WriteRune((r-k+26)%26 + 'a')
	}

	return res.String()
}
