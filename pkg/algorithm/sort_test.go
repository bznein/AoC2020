package algorithm

import (
	"reflect"
	"testing"
)

var sortTests = []struct {
	in  []int
	out []int
}{
	{[]int{1, 6, 2, 3}, []int{1, 2, 3, 6}},
	{[]int{1, -6, 2, 3}, []int{-6, 1, 2, 3}},
	{[]int{1, 2, 3}, []int{1, 2, 3}},
	{[]int{}, []int{}},
}

func TestCountingSort(t *testing.T) {
	for _, test := range sortTests {
		res := CountingSort(test.in)
		if !reflect.DeepEqual(res, test.out) {
			t.Errorf("Error, with input %v expected %v, got %v", test.in, test.out, res)
		}
	}
}
