package lib

import "testing"

func TestReduceSum(t *testing.T) {
	lst := ListFromArray[Int]([]Int{1, 2, 3, 4, 5, 6, 7, 8, 9})
	var expected Int = 45
	actual := reduce[Int, Int](lst, func(a, b Int) Int { return a + b }, 0)
	if expected != actual {
		t.Errorf("Expected %v but got %v", expected, actual)
	}
	t.Logf("TestReduceSum: Expected %v, got  %v", expected, actual)
}
