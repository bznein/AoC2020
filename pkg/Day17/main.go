package Day17

import (
	"time"

	"github.com/bznein/AoC2020/pkg/input"
	"github.com/bznein/AoC2020/pkg/timing"
)

const (
	active    = '#'
	inactive  = '.'
	smallSize = 7
	size      = 12
	steps     = 6
)

type threeDcoords struct {
	x int
	y int
	z int
}

type fourDcoords struct {
	w int
	threeDcoords
}

type threeDGrid map[threeDcoords]rune

type fourDGrid map[fourDcoords]rune

func getNeighbours(t3d threeDGrid, t4d fourDGrid, startPos4d fourDcoords) (int, int) {
	tot3d := 0
	tot4d := 0
	startPos3d := startPos4d.threeDcoords
	w := startPos4d.w
	z := startPos3d.x
	j := startPos3d.z
	i := startPos3d.y
	for ww := w - 1; ww <= w+1; ww++ {
		for zz := z - 1; zz <= z+1; zz++ {
			for jj := j - 1; jj <= j+1; jj++ {
				for ii := i - 1; ii <= i+1; ii++ {
					pos3d := threeDcoords{zz, ii, jj}
					pos4d := fourDcoords{ww, pos3d}
					if w == ww && startPos3d != pos3d && t3d[pos3d] == active {
						tot3d++
					}
					if startPos4d != pos4d && t4d[pos4d] == active {
						tot4d++
					}
				}
			}
		}
	}
	return tot3d, tot4d
}

func oneStep(t1 threeDGrid, t2 fourDGrid) (threeDGrid, fourDGrid, int, int) {
	copyMap1 := threeDGrid{}
	copyMap2 := fourDGrid{}
	totActive3d := 0
	totActive4d := 0
	for w := -smallSize + 1; w < smallSize; w++ {
		for z := -smallSize + 1; z < smallSize; z++ {
			for i := -size + 1; i < size; i++ {
				for j := -size + 1; j < size; j++ {
					pos3d := threeDcoords{z, i, j}
					pos4d := fourDcoords{w, pos3d}
					nActive3d, nActive4d := getNeighbours(t1, t2, pos4d)
					if w == 0 {
						isActive := t1[pos3d] == active
						if (nActive3d == 3 && !isActive) || ((nActive3d == 3 || nActive3d == 2) && isActive) {
							copyMap1[pos3d] = active
							totActive3d++
						} else {
							copyMap1[pos3d] = inactive
						}
					}

					isActive := t2[pos4d] == active
					if (nActive4d == 3 && !isActive) || ((nActive4d == 3 || nActive4d == 2) && isActive) {
						copyMap2[pos4d] = active
						totActive4d++
					} else {
						copyMap2[pos4d] = inactive
					}
				}
			}
		}
	}
	return copyMap1, copyMap2, totActive3d, totActive4d
}

func Solve(inputF string) (int, int) {
	defer timing.TimeTrack(time.Now())

	n := input.InputToStringSlice(inputF)

	t := threeDGrid{}
	t2 := fourDGrid{}

	for w := -smallSize; w <= smallSize; w++ {
		for x := -smallSize; x <= smallSize; x++ {
			for y := -size; y <= size; y++ {
				for z := -size; z <= size; z++ {
					if w == 0 {
						t[threeDcoords{x, y, z}] = inactive
					}
					t2[fourDcoords{w, threeDcoords{x, y, z}}] = inactive
				}
			}

		}
	}

	for i := range n {
		for j := range n[i] {
			if n[i][j] == active {
				t[threeDcoords{0, i, j}] = rune(n[i][j])
				t2[fourDcoords{0, threeDcoords{0, i, j}}] = rune(n[i][j])
			}
		}
	}

	part1 := 0
	part2 := 0
	for i := 0; i < steps; i++ {
		t, t2, part1, part2 = oneStep(t, t2)
	}

	return part1, part2
}
