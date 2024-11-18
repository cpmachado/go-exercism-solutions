package clock

import (
	"fmt"
)

// Define the Clock type here.
type Clock struct {
	Hours   int
	Minutes int
}

func New(h, m int) Clock {
	if m < 0 {
		h -= (1 - (m / 60))
		m = 60 + m%60
	}
	if h < 0 {
		h = 24 + h%24
	}
	h = (h + m/60) % 24
	m %= 60
	return Clock{h, m}
}

func (c Clock) Add(m int) Clock {
	return New(c.Hours, c.Minutes+m)
}

func (c Clock) Subtract(m int) Clock {
	return New(c.Hours, c.Minutes-m)
}

func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.Hours, c.Minutes)
}
