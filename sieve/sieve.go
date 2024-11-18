package sieve

func Sieve(limit int) []int {
	if limit < 2 {
		return []int{}
	}
	primes := []int{2}
	k := 1

	for i := 3; i <= limit; i += 2 {
		j := 0
		for ; j < k; j++ {
			if i%primes[j] == 0 {
				break
			}
		}
		if j == k {
			primes = append(primes, i)
			k++
		}
	}
	return primes
}
