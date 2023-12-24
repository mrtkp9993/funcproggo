package lib

type Any interface {
	Equal(Any) Bool
}
type Bool bool

func (b Bool) Equal(a Any) Bool {
	return Bool(b == a.(Bool))
}

type Char rune

func (c Char) Equal(a Any) Bool {
	return Bool(c == a.(Char))
}

type String []Char

func (s String) Equal(a Any) Bool {
	if len(s) != len(a.(String)) {
		return false
	}
	for i, v := range s {
		if !v.Equal(a.(String)[i]) {
			return false
		}
	}
	return true
}

type Int int64

func (i Int) Equal(a Any) Bool {
	return Bool(i == a.(Int))
}

type Frac float64

func (f Frac) Equal(a Any) Bool {
	return Bool(f == a.(Frac))
}

type List[T Any] []T

func Length[T Any](lst List[T]) Int {
	return Int(len(lst))
}

func Equal[T Any](lst1, lst2 List[T]) Bool {
	if Length(lst1) != Length(lst2) {
		return false
	}
	for i, v := range lst1 {
		if !v.Equal(lst2[i]) {
			return false
		}
	}
	return true
}

func ListFromArray[T Any](array []T) List[T] {
	list := make(List[T], len(array))
	copy(list, array)
	return list
}
