package sublist

// Relation type is defined in relations.go file.

func isEqual(l1, l2 []int) bool {
	if len(l1) != len(l2) {
		return false
	}
	for i, v := range l1 {
		if v != l2[i] {
			return false
		}
	}
	return true
}

func isSublist(l1, l2 []int) bool {
	nl1, nl2 := len(l1), len(l2)
	if nl1 > nl2 {
		return false
	}
	for i := 0; nl1+i <= nl2; i++ {
		j := 0
		for ; j < nl1; j++ {
			if l2[i+j] != l1[j] {
				break
			}
		}
		if j == nl1 {
			return true
		}
	}
	return false
}

func Sublist(l1, l2 []int) Relation {
	switch {
	case isEqual(l1, l2):
		return RelationEqual
	case isSublist(l1, l2):
		return RelationSublist
	case isSublist(l2, l1):
		return RelationSuperlist
	}
	return RelationUnequal
}
