package lib

import (
	"testing"
)

func Square[N Int](x N) N {
	return x * x
}

func TestMapSquare(t *testing.T) {
	lst := ListFromArray[Int]([]Int{1, 2, 3, 4, 5, 6, 7, 8, 9})
	expected := ListFromArray[Int]([]Int{1, 4, 9, 16, 25, 36, 49, 64, 81})
	actual := Map[Int, Int](lst, Square[Int])
	if !Equal[Int](expected, actual) {
		t.Errorf("Expected %v but got %v", expected, actual)
	}
	t.Logf("TestMapSquare: Expected %v, got  %v", expected, actual)
}
