package lib

func Pipe2[A, B, C Any](a A, f func(A) B, g func(B) C) C {
	return g(f(a))
}

func Pipe3[A, B, C, D Any](a A, f func(A) B, g func(B) C, h func(C) D) D {
	return h(g(f(a)))
}
