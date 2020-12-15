package Day15

import (
	"fmt"
	"runtime"
	"strconv"
	"strings"
	"time"

	"github.com/bznein/AoC2020/pkg/algorithm"
	"github.com/bznein/AoC2020/pkg/timing"
)

func Solve(inputF string) (int, int) {
	defer timing.TimeTrack(time.Now())
	part1, part2 := -1, -1

	startingNumbers := strings.Split(inputF, ",")

	//	lastTimeSpoken := map[int]int{}
	lastTimeSpoken := make([]int, 30000000)
	for i := range lastTimeSpoken {
		lastTimeSpoken[i] = -1
	}
	for i, v := range startingNumbers {
		val, _ := strconv.Atoi(v)
		lastTimeSpoken[val] = i
	}
	nextNumberSpoken := 0

	max := 0
	for i := len(startingNumbers); i < 30000000-1; i++ {
		last := lastTimeSpoken[nextNumberSpoken]
		max = algorithm.Max(max, nextNumberSpoken)
		lastTimeSpoken[nextNumberSpoken] = i
		if last == -1 {
			nextNumberSpoken = 0
		} else {
			nextNumberSpoken = i - last
		}
		if i == 2018 {
			part1 = nextNumberSpoken
		}
	}

	part2 = nextNumberSpoken
	PrintMemUsage()
	return part1, part2
}
