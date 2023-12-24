package lib

type Predicate[A any] func(A) bool

func filter[A Any](lst List[A], f Predicate[A]) List[A] {
	var ret List[A]
	for _, v := range lst {
		if f(v) {
			ret = append(ret, v)
		}
	}
	return ret
}

func IsAny[A Any](lst List[A], f Predicate[A]) bool {
	for _, v := range lst {
		if f(v) {
			return true
		}
	}
	return false
}

func IsAll[A Any](lst List[A], f Predicate[A]) bool {
	for _, v := range lst {
		if !f(v) {
			return false
		}
	}
	return true
}
