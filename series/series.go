package series

func All(n int, s string) []string {
	ret := []string{}
	for i := 0; i < len(s)-n+1; i++ {
		ret = append(ret, s[i:n+i])
	}
	return ret
}

func UnsafeFirst(n int, s string) string {
	return s[:n]
}

func First(n int, s string) (string, bool) {
	if len(s) < n {
		return "", false
	}
	return s[:n], true
}
