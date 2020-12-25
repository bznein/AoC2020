package Day25

import (
	"time"

	"github.com/bznein/AoC2020/pkg/input"
	"github.com/bznein/AoC2020/pkg/timing"
)

func transform(loopSize int, subject int) int {
	v := 1
	for i := 0; i < loopSize; i++ {
		v *= subject
		v = v % 20201227
	}
	return v
}

func Solve(inputF string) (int, int) {
	defer timing.TimeTrack(time.Now())
	n := input.InputToIntSlice(inputF)

	loopSize := 0
	v := 1
	for {
		if v == n[0] {
			return transform(loopSize, n[1]), -1
		}
		if v == n[1] {
			return transform(loopSize, n[0]), -1
		}
		v *= 7
		v = v % 20201227
		loopSize++
	}
}
