package Day13

import (
	"math"
	"math/big"
	"strconv"
	"strings"
	"time"

	"github.com/bznein/AoC2020/pkg/input"
	"github.com/bznein/AoC2020/pkg/timing"
)

func Solve(inputF string) (int, int) {
	defer timing.TimeTrack(time.Now())
	part2 := -1

	s := input.InputToStringSlice(inputF)
	depTime, _ := strconv.Atoi(s[0])

	busTimes := map[int]int{}
	minWait := math.MaxInt32
	totProd := 1
	minBusIdx := -1
	for i, v := range strings.Split(s[1], ",") {
		if val, err := strconv.Atoi(v); err == nil {
			waitTime := (val) - depTime%val
			if waitTime < minWait {
				minWait = waitTime
				minBusIdx = val
			}
			busTimes[i] = -val
			totProd *= val
		}
	}

	// b1..bn are the keys in busTimes
	// n1..nn are the values in bustimes

	bigN := map[int]int{}
	y := map[int]int{}

	for i, v := range busTimes {
		bigN[i] = totProd / v
		bigInt := big.NewInt(int64(bigN[i]))
		y[i] = int(bigInt.ModInverse(bigInt, big.NewInt(int64(v))).Int64())
	}

	x0 := 0
	for i := range busTimes {
		x0 += i * bigN[i] * y[i]
	}
	part2 = x0 % totProd

	return minWait * minBusIdx, -part2
}
