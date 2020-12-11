package Day03

import (
	"time"

	"github.com/bznein/AoC2020/pkg/input"
	"github.com/bznein/AoC2020/pkg/timing"
)

type maze []string

type pos struct {
	i int
	j int
}

var try int
var part1 int
var part2 int

//TODO visualization currently broken due to refatoring

func (position *pos) moveBy(i, j int, limit int) {
	position.i += i
	position.j = (position.j + j) % limit
}

func (position *pos) moveBySlope(s slope, limit int) {
	position.moveBy(s.i, s.j, limit)
}

type slope struct {
	i int
	j int
}

func explore(m maze, slopes []slope) (int, int) {
	p2 := 1

	positions := make([]pos, len(slopes))
	results := make([]int, len(slopes))
	finished := map[int]bool{}
	outside := 0
	for {
		for i, p := range positions {
			if p.i >= len(m) {
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

			if m[p.i][p.j] == '#' {
				results[i]++
			}
			p.moveBySlope(slopes[i], len(m[0]))
			positions[i] = p
		}
	}

}

func Solve(inputF string) (int, int) {
	defer timing.TimeTrack(time.Now())
	var m maze
	m = input.InputToStringSlice(inputF)
	return explore(m, []slope{{1, 3}, {1, 1}, {1, 5}, {1, 7}, {2, 1}})
}
