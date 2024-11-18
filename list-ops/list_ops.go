package listops

// IntList is an abstraction of a list of integers which we can define methods on
type IntList []int

func (s IntList) Foldl(fn func(int, int) int, initial int) int {
	acc := initial

	for _, v := range s {
		acc = fn(acc, v)
	}

	return acc
}

func (s IntList) Foldr(fn func(int, int) int, initial int) int {
	acc := initial
	n := len(s) - 1
	for i, _ := range s {
		acc = fn(s[n-i], acc)
	}
	return acc
}

func (s IntList) Filter(fn func(int) bool) IntList {
	ret := IntList{}

	for _, v := range s {
		if fn(v) {
			ret = append(ret, v)
		}
	}

	return ret
}

func (s IntList) Length() int {
	return len(s)
}

func (s IntList) Map(fn func(int) int) IntList {
	ret := IntList{}

	for _, v := range s {
		ret = append(ret, fn(v))
	}

	return ret
}

func (s IntList) Reverse() IntList {
	ret := IntList{}
	n := len(s)
	for i := 0; i < n; i++ {
		ret = append(ret, s[n-i-1])
	}
	return ret
}

func (s IntList) Append(lst IntList) IntList {
	return append(s, lst...)
}

func (s IntList) Concat(lists []IntList) IntList {
	ret := IntList{}.Append(s)
	for _, lst := range lists {
		ret = ret.Append(lst)
	}
	return ret
}
