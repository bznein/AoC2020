package algorithm

import (
	"testing"
)

var searchTests = []struct {
	in     []int
	search int
	out    int
}{
	{[]int{1, 2, 3, 4}, 2, 1},
	{[]int{1, 2, 3, 4}, 4, 3},
	{[]int{1, 2, 3, 4}, -2, -1},
	{[]int{}, 2, -1},
}

func TestBinarySearch(t *testing.T) {
	for _, test := range searchTests {
		res := BinarySearch(test.in, test.search)
		if res != test.out {
			t.Errorf("Error, with input %v, searching for %v, expected %d, got %d", test.in, test.search, test.out, res)
		}
	}
}
