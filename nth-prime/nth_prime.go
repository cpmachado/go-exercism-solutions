package prime

import "errors"

func Nth(n int) (int, error) {
	if n < 1 {
		return 0, errors.New("n must be >= 1")
	}
	primes := make([]int, n)
	primes[0] = 2
	k := 1

	for i := 3; k < n; i += 2 {
		j := 0
		for ; j < k; j++ {
			if i%primes[j] == 0 {
				break
			}
		}
		if j == k {
			primes[k] = i
			k++
		}
	}
	return primes[k-1], nil
}
