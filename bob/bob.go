package bob

import (
	"strings"
	"unicode"
)

// Hey should have a comment documenting it.
func Hey(remark string) string {
	trimmed := strings.TrimSpace(remark)
	if trimmed == "" {
		return "Fine. Be that way!"
	}

	isQuestion := strings.HasSuffix(trimmed, "?")
	isShout := strings.IndexFunc(trimmed, unicode.IsLetter) >= 0 && strings.ToUpper(trimmed) == trimmed

	switch {
	case isShout && isQuestion:
		return "Calm down, I know what I'm doing!"
	case isShout:
		return "Whoa, chill out!"
	case isQuestion:
		return "Sure."
	default:
		return "Whatever."
	}
}
