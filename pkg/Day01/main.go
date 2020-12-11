package Day01

import (
	"sort"
	"time"

	"github.com/bznein/AoC2020/pkg/algorithm"
	"github.com/bznein/AoC2020/pkg/input"
	"github.com/bznein/AoC2020/pkg/timing"
	"github.com/bznein/AoC2020/pkg/visualize"
	day "github.com/bznein/AoC2020/pkg/visualize/Day01"
)

const (
	target = 2020
)

func Solve(inputF string) (int, int) {
	defer timing.TimeTrack(time.Now())
	part1 := -1
	part2 := -1
	ints := input.InputToIntSlice(inputF)
	sort.Ints(ints)
	visualize.Init()
	defer visualize.Close()
	for idx, i := range ints {
		if i > target {
			continue
		}
		if part2 == -1 {
			lower := idx + 1
			upper := len(ints) - 1
			for lower < upper {
				sum := i + ints[lower] + ints[upper]
				if sum == target {
					part2 = i * ints[lower] * ints[upper]
					day.ShowCurrentResult(part1, &part2)
					if part1 == -1 {
						return part1, part2
					}
					break
				} else if sum < target {
					lower++
				} else {
					upper--
				}
			}
		}
		if part1 == -1 {
			day.ShowSearchLine(target, i, nil)
			res := algorithm.BinarySearch(ints, target-i)
			if res != -1 && res != idx {
				part1 = ints[res] * i
				day.ShowCurrentResult(part1, nil)
				if part2 != -1 {
					break
				}
			}
		}
	}

	return part1, part2
}
