package collatzconjecture

import "errors"

func CollatzConjecture(n int) (int, error) {
	if n < 1 {
		return 0, errors.New("Needs to be positive integer")
	}
	count := 0

	for n > 1 {
		if n%2 == 1 {
			n = 3*n + 1
		} else {
			n /= 2
		}
		count++
	}
	return count, nil
}
