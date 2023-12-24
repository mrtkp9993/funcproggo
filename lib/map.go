package lib

func Map[A Any, B Any](lst List[A], f func(A) B) List[B] {
	var ret List[B]
	for _, v := range lst {
		ret = append(ret, f(v))
	}
	return ret
}
