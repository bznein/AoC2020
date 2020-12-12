package Day11

import (
	"strings"
	"time"

	"github.com/bznein/AoC2020/pkg/algorithm"
	"github.com/bznein/AoC2020/pkg/input"
	"github.com/bznein/AoC2020/pkg/timing"
)

const (
	floor    = '.'
	empty    = 'L'
	occupied = '#'
)

func occupiedAdjacentSeatsByOne(seats []string, row int, column int) int {
	tot := 0

	upRow := row - 1
	leftCol := column - 1
	downRow := row + 1
	rightCol := column + 1

	canLookUp := upRow >= 0
	canLookDown := downRow < len(seats)
	canLookLeft := leftCol >= 0
	canLookRight := rightCol < len(seats[0])
	canLookUpLeft := canLookUp && canLookLeft
	canLookUpRight := canLookUp && canLookRight
	canLookDownLeft := canLookDown && canLookLeft
	canLookDownRight := canLookDown && canLookRight

	// Look up
	if canLookUp && seats[upRow][column] == occupied {
		tot++
	}

	// Look down
	if canLookDown && seats[downRow][column] == occupied {
		tot++
	}

	// Look right
	if canLookRight && seats[row][rightCol] == occupied {
		tot++
	}

	// Look left
	if canLookLeft && seats[row][leftCol] == occupied {
		tot++
	}

	// Look Upper Left
	if canLookUpLeft && seats[upRow][leftCol] == occupied {
		tot++
	}

	// Look Upper Right
	if canLookUpRight && seats[upRow][rightCol] == occupied {
		tot++
	}

	// Look Lower Left
	if canLookDownLeft && seats[downRow][leftCol] == occupied {
		tot++
	}

	// Look Lower Right
	if canLookDownRight && seats[downRow][rightCol] == occupied {
		tot++
	}

	return tot
}

func occupiedAdjacentSeats(seats []string, row int, column int, maxLook int) int {
	if maxLook == 1 {
		return occupiedAdjacentSeatsByOne(seats, row, column)
	}
	tot := 0

	// Look up
	for i := row - 1; i >= 0; i-- {
		if seats[i][column] == occupied {
			tot++
			break
		} else if seats[i][column] == empty {
			break
		}
	}
	// Look down
	for i := row + 1; i < len(seats); i++ {
		if seats[i][column] == occupied {
			tot++
			break
		} else if seats[i][column] == empty {
			break
		}
	}
	// Look left
	for j := column - 1; j >= 0; j-- {
		if seats[row][j] == occupied {
			tot++
			break
		} else if seats[row][j] == empty {
			break
		}
	}
	// Look right
	for j := column + 1; j < len(seats[row]); j++ {
		if seats[row][j] == occupied {
			tot++
			break
		} else if seats[row][j] == empty {
			break
		}
	}

	// Look Upper Left
	for i, j := row-1, column-1; i >= 0 && j >= 0; i, j = i-1, j-1 {
		if seats[i][j] == occupied {
			tot++
			break
		} else if seats[i][j] == empty {
			break
		}
	}

	// Look Upper Right
	for i, j := row-1, column+1; i >= 0 && j < len(seats[i]); i, j = i-1, j+1 {
		if seats[i][j] == occupied {
			tot++
			break
		} else if seats[i][j] == empty {
			break
		}
	}

	// Look Lower Left
	for i, j := row+1, column-1; i < len(seats) && j >= 0; i, j = i+1, j-1 {
		if seats[i][j] == occupied {
			tot++
			break
		} else if seats[i][j] == empty {
			break
		}
	}

	// Look Lower Right
	for i, j := row+1, column+1; i < len(seats) && j < len(seats[row]); i, j = i+1, j+1 {
		if seats[i][j] == occupied {
			tot++
			break
		} else if seats[i][j] == empty {
			break
		}
	}
	return tot
}

func solveOnePart(seats []string, maxOccupied int, maxLook int) int {
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

	n := input.InputToStringSlice(inputF)
	part1 = solveOnePart(n, 4, 1)
	n = input.InputToStringSlice(inputF)
	part2 = solveOnePart(n, 5, algorithm.Max(len(n), len(n[0])))
	return part1, part2
}