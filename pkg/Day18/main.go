package Day18

import (
	"time"

	p1 "github.com/bznein/AoC2020/pkg/Day18/part1"
	p2 "github.com/bznein/AoC2020/pkg/Day18/part2"
	"github.com/bznein/AoC2020/pkg/input"
	"github.com/bznein/AoC2020/pkg/timing"
)

func Solve(inputF string) (int, int) {
	defer timing.TimeTrack(time.Now())

	part1 := 0
	part2 := 0

	s := input.InputToStringSlice(inputF)

	c1 := make(chan int)
	go p1.SolvePart1(s, c1)
	c2 := make(chan int)
	go p2.SolvePart2(s, c2)
	for {
		if part1 != 0 && part2 != 0 {
			break
		}
		select {
		case p := <-c1:
			part1 = p
		case p := <-c2:
			part2 = p
		}
	}

	return part1, part2
}
