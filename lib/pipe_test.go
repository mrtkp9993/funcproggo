package lib

import "testing"

func TestPipe2(t *testing.T) {
	var expected Int = 11
	actual := Pipe2[Int, Int, Int](5, func(x Int) Int { return x * 2 }, func(x Int) Int { return x + 1 })
	if expected != actual {
		t.Errorf("Expected %v but got %v", expected, actual)
	}
	t.Logf("TestPipe2: Expected %v, got  %v", expected, actual)
}

func TestPipe3(t *testing.T) {
	// sum of squares even numbers from 1 to 10
	var expected Int = 220
	nums := ListFromArray[Int]([]Int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
	// use map-filter-reduce
	actual := Pipe3[List[Int], List[Int], List[Int], Int](
		nums,
		func(lst List[Int]) List[Int] {
			return Map(lst, func(x Int) Int { return x * x })
		},
		func(lst List[Int]) List[Int] {
			return filter(lst, func(x Int) bool { return x%2 == 0 })
		}, func(lst List[Int]) Int {
			return reduce(lst, func(x, y Int) Int { return x + y }, 0)
		},
	)
	if expected != actual {
		t.Errorf("Expected %v but got %v", expected, actual)
	}
	t.Logf("TestPipe3: Expected %v, got  %v", expected, actual)
}
