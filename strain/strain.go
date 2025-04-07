package strain

func Keep[T any](s []T, p func(T) bool) []T {
	ret := []T{}

	for _, v := range s {
		if p(v) {
			ret = append(ret, v)
		}
	}

	return ret
}

func Discard[T any](s []T, p func(T) bool) []T {
	ret := []T{}

	for _, v := range s {
		if !p(v) {
			ret = append(ret, v)
		}
	}

	return ret
}
