package stringset

import (
	"fmt"
	"strings"
)

// Implement Set as a collection of unique string values.
//
// For Set.String, use '{' and '}', output elements as double-quoted strings
// safely escaped with Go syntax, and use a comma and a single space between
// elements. For example, a set with 2 elements, "a" and "b", should be formatted as {"a", "b"}.
// Format the empty set as {}.

// Define the Set type here.
type Set map[string]bool

func New() Set {
	return make(map[string]bool)
}

func NewFromSlice(l []string) Set {
	s := New()
	for _, k := range l {
		s.Add(k)
	}
	return s
}

func (s Set) String() string {
	var sb strings.Builder
	keys := []string{}

	sb.WriteRune('{')

	for k := range s {
		keys = append(keys, fmt.Sprintf("%q", k))
	}

	sb.WriteString(strings.Join(keys, ", "))

	sb.WriteRune('}')
	return sb.String()
}

func (s Set) IsEmpty() bool {
	for range s {
		return false
	}
	return true
}

func (s Set) Has(elem string) bool {
	_, found := s[elem]
	return found
}

func (s Set) Add(elem string) {
	s[elem] = true
}

func Subset(s1, s2 Set) bool {
	for k := range s1 {
		if !s2.Has(k) {
			return false
		}
	}
	return true
}

func Disjoint(s1, s2 Set) bool {
	for k := range s1 {
		if s2.Has(k) {
			return false
		}
	}
	return true
}

func Equal(s1, s2 Set) bool {
	return Subset(s1, s2) && Subset(s2, s1)
}

func Intersection(s1, s2 Set) Set {
	i := New()

	for k := range s1 {
		if s2.Has(k) {
			i.Add(k)
		}
	}

	return i
}

func Difference(s1, s2 Set) Set {
	d := New()

	for k := range s1 {
		if !s2.Has(k) {
			d.Add(k)
		}
	}

	return d
}

func Union(s1, s2 Set) Set {
	u := New()

	for k := range s1 {
		u.Add(k)
	}
	for k := range s2 {
		u.Add(k)
	}

	return u
}
