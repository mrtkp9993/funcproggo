package lib

import (
	"testing"
)

func isEven[N Int](x N) bool {
	return x%2 == 0
}

func TestFilterIsEven(t *testing.T) {
	lst := ListFromArray[Int]([]Int{1, 2, 3, 4, 5, 6, 7, 8, 9})
	expected := ListFromArray[Int]([]Int{2, 4, 6, 8})
	actual := filter[Int](lst, isEven[Int])
	if !Equal[Int](expected, actual) {
		t.Errorf("Expected %v but got %v", expected, actual)
	}
	t.Logf("TestFilterIsEven: Expected %v, got  %v", expected, actual)
}
