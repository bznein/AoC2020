package Day10

import (
	"time"

	"github.com/bznein/AoC2020/pkg/algorithm"
	"github.com/bznein/AoC2020/pkg/input"
	"github.com/bznein/AoC2020/pkg/timing"
)

func Solve(inputF string) (int, int) {
	defer timing.TimeTrack(time.Now())
	n := []int{0}
	n = append(n, input.InputToIntSlice(inputF)...)
	n = algorithm.CountingSort(n)
	n = append(n, n[len(n)-1]+3)
	c := make([]int, len(n))
	c[0] = 1
	diffs := map[int]int{}
	for i := range n {
		for j := algorithm.Max(0, i-3); j < i; j++ {
			if n[i]-n[j] <= 3 {
				c[i] += c[j]
			}
		}
		if i < len(n)-1 {
			diffs[n[i+1]-n[i]]++
		}
	}
	return diffs[1] * diffs[3], c[len(c)-1]
}
