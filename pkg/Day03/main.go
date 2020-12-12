package Day03

import (
	"time"

	"github.com/bznein/AoC2020/pkg/input"
	"github.com/bznein/AoC2020/pkg/timing"
	"github.com/bznein/AoC2020/pkg/twod"
)

//TODO visualization currently broken due to refatoring

func explore(m twod.Grid, slopes []twod.Slope) (int, int) {
	p2 := 1

	positions := make([]twod.Position, len(slopes))
	results := make([]int, len(slopes))
	finished := map[int]bool{}
	outside := 0
	for {
		for i, p := range positions {
			if p.I >= len(m) {
				if _, ok := finished[i]; !ok {
					p2 *= results[i]
					outside++
				}
				finished[i] = true
				if outside == len(positions) {
					return results[0], p2
				}
				continue
			}

			if m.IsEntry(p.I, p.J, '#') {
				results[i]++
			}
			p.MoveBySlopeWithWrapping(slopes[i], -1, len(m[0]))
			positions[i] = p
		}
	}

}

func Solve(inputF string) (int, int) {
	defer timing.TimeTrack(time.Now())
	m := twod.Grid(input.InputToStringSlice(inputF))
	return explore(m, []twod.Slope{{1, 3}, {1, 1}, {1, 5}, {1, 7}, {2, 1}})
}
