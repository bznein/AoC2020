package Day15

import (
	"strconv"
	"strings"
	"time"

	"github.com/bznein/AoC2020/pkg/timing"
)

func Solve(inputF string) (int, int) {
	defer timing.TimeTrack(time.Now())
	part1, part2 := -1, -1

	startingNumbers := strings.Split(inputF, ",")

	lastTimeSpoken := make([]int, 30000000)
	for i, v := range startingNumbers {
		val, _ := strconv.Atoi(v)
		lastTimeSpoken[val] = i + 1
	}
	nextNumberSpoken := 0

	for i := len(startingNumbers); i < 30000000-1; i++ {
		last := lastTimeSpoken[nextNumberSpoken]

		lastTimeSpoken[nextNumberSpoken] = i + 1
		if last == 0 {
			nextNumberSpoken = 0
		} else {
			nextNumberSpoken = i + 1 - last
		}
		if i == 2018 {
			part1 = nextNumberSpoken
		}
	}

	part2 = nextNumberSpoken
	return part1, part2
}
