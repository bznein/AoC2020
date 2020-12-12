package Day11

import (
	"strings"
	"time"

	"github.com/bznein/AoC2020/pkg/algorithm"
	"github.com/bznein/AoC2020/pkg/input"
	"github.com/bznein/AoC2020/pkg/timing"
	"github.com/bznein/AoC2020/pkg/twod"
)

const (
	floor    = '.'
	empty    = 'L'
	occupied = '#'
)

func occupiedAdjacentSeatsByOne(seats twod.Grid, row int, column int) int {
	tot := 0

	// Look up
	if seats.IsEntry(row-1, column, occupied) {
		tot++
	}

	// Look down
	if seats.IsEntry(row+1, column, occupied) {
		tot++
	}

	// Look right
	if seats.IsEntry(row, column+1, occupied) {
		tot++
	}

	// Look leftt
	if seats.IsEntry(row, column-1, occupied) {
		tot++
	}

	// Look Upper Left
	if seats.IsEntry(row-1, column-1, occupied) {
		tot++
	}

	// Look Upper Right
	if seats.IsEntry(row-1, column+1, occupied) {
		tot++
	}

	// Look Lower Left
	if seats.IsEntry(row+1, column-1, occupied) {
		tot++
	}

	// Look Lower Right
	if seats.IsEntry(row+1, column+1, occupied) {
		tot++
	}

	return tot
}

func isOccupiedOrAtLeastNotFLoor(seats twod.Grid, i int, j int) (bool, bool) {
	c := seats.UnsafeEntry(i, j)
	return c == occupied, c != floor
}

func occupiedAdjacentSeats(seats twod.Grid, row int, column int, maxLook int) int {
	if maxLook == 1 {
		return occupiedAdjacentSeatsByOne(seats, row, column)
	}
	tot := 0

	// Look up
	for i := row - 1; i >= 0; i-- {
		incr, shouldBreak := isOccupiedOrAtLeastNotFLoor(seats, i, column)
		if incr {
			tot++
		}
		if shouldBreak {
			break
		}
	}
	// Look down
	for i := row + 1; i < len(seats); i++ {
		incr, shouldBreak := isOccupiedOrAtLeastNotFLoor(seats, i, column)
		if incr {
			tot++
		}
		if shouldBreak {
			break
		}
	}
	// Look left
	for j := column - 1; j >= 0; j-- {
		incr, shouldBreak := isOccupiedOrAtLeastNotFLoor(seats, row, j)
		if incr {
			tot++
		}
		if shouldBreak {
			break
		}
	}
	// Look right
	for j := column + 1; j < len(seats[row]); j++ {
		incr, shouldBreak := isOccupiedOrAtLeastNotFLoor(seats, row, j)
		if incr {
			tot++
		}
		if shouldBreak {
			break
		}
	}

	// Look Upper Left
	for i, j := row-1, column-1; i >= 0 && j >= 0; i, j = i-1, j-1 {
		incr, shouldBreak := isOccupiedOrAtLeastNotFLoor(seats, i, j)
		if incr {
			tot++
		}
		if shouldBreak {
			break
		}
	}

	// Look Upper Right
	for i, j := row-1, column+1; i >= 0 && j < len(seats[i]); i, j = i-1, j+1 {
		incr, shouldBreak := isOccupiedOrAtLeastNotFLoor(seats, i, j)
		if incr {
			tot++
		}
		if shouldBreak {
			break
		}
	}

	// Look Lower Left
	for i, j := row+1, column-1; i < len(seats) && j >= 0; i, j = i+1, j-1 {
		incr, shouldBreak := isOccupiedOrAtLeastNotFLoor(seats, i, j)
		if incr {
			tot++
		}
		if shouldBreak {
			break
		}
	}

	// Look Lower Right
	for i, j := row+1, column+1; i < len(seats) && j < len(seats[row]); i, j = i+1, j+1 {
		incr, shouldBreak := isOccupiedOrAtLeastNotFLoor(seats, i, j)
		if incr {
			tot++
		}
		if shouldBreak {
			break
		}
	}
	return tot
}

func solveOnePart(seats twod.Grid, maxOccupied int, maxLook int) int {
	for {
		changes := false
		totOccupied := 0
		sbs := make([]strings.Builder, len(seats))
		for i, row := range seats {
			for j, column := range row {
				switch column {
				case floor:
					sbs[i].WriteRune(floor)
				case empty:
					if occupiedAdjacentSeats(seats, i, j, maxLook) == 0 {
						sbs[i].WriteRune(occupied)
						totOccupied++
						changes = true
					} else {
						sbs[i].WriteRune(empty)
					}
				case occupied:
					if occupiedAdjacentSeats(seats, i, j, maxLook) >= maxOccupied {
						sbs[i].WriteRune(empty)
						changes = true
					} else {
						sbs[i].WriteRune(occupied)
						totOccupied++
					}
				}
			}
		}
		if !changes {
			return totOccupied

		}
		for i := range seats {
			seats[i] = sbs[i].String()
		}
	}
}

func Solve(inputF string) (int, int) {
	defer timing.TimeTrack(time.Now())
	part1 := 0
	part2 := 0

	n := twod.Grid(input.InputToStringSlice(inputF))
	n2 := make(twod.Grid, len(n))
	copy(n2, n)
	part1 = solveOnePart(n, 4, 1)
	part2 = solveOnePart(n2, 5, algorithm.Max(len(n), len(n[0])))
	return part1, part2
}
