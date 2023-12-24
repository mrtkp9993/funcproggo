package lib

func reduce[A Any, B Any](lst List[A], f func(B, A) B, init B) B {
	res := init
	for _, v := range lst {
		res = f(res, v)
	}
	return res
}
