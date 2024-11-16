package reverse

func Reverse(input string) string {
	ret := ""

	for _, v := range input {
		ret = string(v) + ret
	}

	return ret
}
