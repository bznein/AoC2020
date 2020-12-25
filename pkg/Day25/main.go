package Day25

import (
	"time"

	"github.com/bznein/AoC2020/pkg/input"
	"github.com/bznein/AoC2020/pkg/timing"
)

func Solve(inputF string) (int, int) {
	defer timing.TimeTrack(time.Now())
	n := input.InputToIntSlice(inputF)

	loopSize := 0
	v := 1
	n1Transformed := 1
	n0Transformed := 1
	for {
		if v == n[0] {
			return n1Transformed, -1
		}
		if v == n[1] {
			return n0Transformed, -1
		}
		v *= 7
		v = v % 20201227

		n1Transformed *= n[1]
		n1Transformed = n1Transformed % 20201227

		n0Transformed *= n[0]
		n0Transformed = n0Transformed % 20201227

		loopSize++
	}
}
