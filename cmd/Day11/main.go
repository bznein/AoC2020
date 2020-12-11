package main

import (
	"fmt"
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

func occupiedAdjacentSeats(seats []string, row int, column int, maxLook int) int {
	tot := 0

	// Look up
	for i := row - 1; i >= 0 && row-i <= maxLook; i-- {
		if seats[i][column] == occupied {
			tot++
			break
		} else if seats[i][column] == empty {
			break
		}
	}
	// Look down
	for i := row + 1; i < len(seats) && i-row <= maxLook; i++ {
		if seats[i][column] == occupied {
			tot++
			break
		} else if seats[i][column] == empty {
			break
		}
	}
	// Look left
	for j := column - 1; j >= 0 && column-j <= maxLook; j-- {
		if seats[row][j] == occupied {
			tot++
			break
		} else if seats[row][j] == empty {
			break
		}
	}
	// Look right
	for j := column + 1; j < len(seats[row]) && j-column <= maxLook; j++ {
		if seats[row][j] == occupied {
			tot++
			break
		} else if seats[row][j] == empty {
			break
		}
	}

	// Look Upper Left
	for i, j := row-1, column-1; i >= 0 && j >= 0 && row-i <= maxLook; i, j = i-1, j-1 {
		if seats[i][j] == occupied {
			tot++
			break
		} else if seats[i][j] == empty {
			break
		}
	}

	// Look Upper Right
	for i, j := row-1, column+1; i >= 0 && j < len(seats[i]) && row-i <= maxLook; i, j = i-1, j+1 {
		if seats[i][j] == occupied {
			tot++
			break
		} else if seats[i][j] == empty {
			break
		}
	}

	// Look Lower Left
	for i, j := row+1, column-1; i < len(seats) && j >= 0 && i-row <= maxLook; i, j = i+1, j-1 {
		if seats[i][j] == occupied {
			tot++
			break
		} else if seats[i][j] == empty {
			break
		}
	}

	// Look Lower Right
	for i, j := row+1, column+1; i < len(seats) && j < len(seats[row]) && i-row <= maxLook; i, j = i+1, j+1 {
		if seats[i][j] == occupied {
			tot++
			break
		} else if seats[i][j] == empty {
			break
		}
	}
	return tot
}

func totalOccupied(seats []string) int {
	tot := 0
	for _, r := range seats {
		for _, c := range r {
			if c == occupied {
				tot++
			}
		}
	}
	return tot
}

//TODO just one function

func solveOnePart(seats []string, maxOccupied int, maxLook int) int {
	defer timing.TimeTrack(time.Now())
	for {
		seatsTemp := make([]string, len(seats))
		changes := false
		totOccupied := 0
		for i, row := range seats {
			var sb strings.Builder
			for j, column := range row {
				switch column {
				case floor:
					sb.WriteRune(floor)
				case empty:
					if occupiedAdjacentSeats(seats, i, j, maxLook) == 0 {
						sb.WriteRune(occupied)
						totOccupied++
						changes = true
					} else {
						sb.WriteRune(empty)
					}
				case occupied:
					if occupiedAdjacentSeats(seats, i, j, maxLook) >= maxOccupied {
						sb.WriteRune(empty)
						changes = true
					} else {
						sb.WriteRune(occupied)
						totOccupied++
					}
				}
			}
			seatsTemp[i] = sb.String()
		}
		if !changes {
			return totOccupied

		}
		copy(seats, seatsTemp)
	}
}

func solve(inputF string) (int, int) {
	part1 := 0
	part2 := 0

	n := input.InputToStringSlice(inputF)
	part1 = solveOnePart(n, 4, 1)
	n = input.InputToStringSlice(inputF)
	part2 = solveOnePart(n, 5, algorithm.Max(len(n), len(n[0])))
	return part1, part2
}

func main() {
	input.ParseFlags()
	inputF := input.ReadInput(fmt.Sprintf("../../inputs/11.txt"))

	p1, p2 := solve(inputF)
	fmt.Printf("Part 1: %d, Part2: %d\n", p1, p2)

}
