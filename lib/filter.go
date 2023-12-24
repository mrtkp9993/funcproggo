package lib

func filter[A Any](lst List[A], f func(A) bool) List[A] {
	var ret List[A]
	for _, v := range lst {
		if f(v) {
			ret = append(ret, v)
		}
	}
	return ret
}
