package prime

func Factors(n int64) []int64 {
	ret := []int64{}

	for n%2 == 0 && n > 0 {
		ret = append(ret, 2)
		n = n / 2
	}

	for i := int64(3); i <= n; i += 2 {
		for n%i == 0 && n > 0 {
			ret = append(ret, i)
			n = n / i
		}
	}

	return ret
}
