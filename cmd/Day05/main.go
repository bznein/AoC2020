package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
	"time"

	"github.com/bznein/AoC2020/pkg/input"
	"github.com/bznein/AoC2020/pkg/term"
)

func column(boardingPass string) int {
	colString := strings.ReplaceAll(strings.ReplaceAll(boardingPass[7:], "R", "1"), "L", "0")
	val, _ := strconv.ParseInt(colString, 2, 64)
	return int(val)

}

func row(boardingPass string) int {
	rowString := strings.ReplaceAll(strings.ReplaceAll(boardingPass[:7], "B", "1"), "F", "0")
	val, _ := strconv.ParseInt(rowString, 2, 64)

	return int(val)
}

func seatID(boardingPass string) int {
	return row(boardingPass)*8 + column(boardingPass)
}

func solve(inputF string) (int, int) {
	part1 := int(0)
	part2 := int(0)
	min := math.MaxInt64

	aeroplane := term.Aeroplane{
		X:           40,
		Y:           25,
		Rows:        128,
		SeatsPerRow: 8,
	}
	if input.Visualize {
		term.Init()
		aeroplane.PaintInit()
		term.Separator(aeroplane.X+4, aeroplane.Y+20, 3)
		term.Tbprint(aeroplane.X+4, aeroplane.Y+20, term.White, term.Black, "/-----------\\")
		term.Tbprint(aeroplane.X+5, aeroplane.Y+21, term.White, term.Black, " YOUR SEAT")
		term.Tbprint(aeroplane.X+5, aeroplane.Y+22, term.White, term.Black, " ROW: ???")
		term.Tbprint(aeroplane.X+5, aeroplane.Y+23, term.White, term.Black, " COL: ???")
		term.Separator(aeroplane.X+16, aeroplane.Y+20, 3)
		term.Tbprint(aeroplane.X+4, aeroplane.Y+24, term.White, term.Black, "\\-----------/")
		time.Sleep(time.Second * 5)
		defer term.Close()
	}
	seats := map[int]bool{}

	boardingPasses := input.InputToStringSlice(inputF)

	for _, b := range boardingPasses {
		if input.Visualize {
			term.Tbprint(aeroplane.X+5, aeroplane.Y+18, term.White, term.Black, fmt.Sprintf("HIGHEST NUMBERED SEAT: %d", part1))
		}
		time.Sleep(time.Millisecond * input.Delay)
		v := seatID(b)
		if v > part1 {
			part1 = v
		} else if v < min {
			min = v
		}
		seats[v] = true
		if input.Visualize {
			term.Tbprint(aeroplane.X+4, aeroplane.Y-15, term.White, term.Black, "/----------------\\")
			term.Tbprint(aeroplane.X+6, aeroplane.Y-13, term.White, term.Black, "BOARDING PASS")
			term.BicolorString(aeroplane.X+6, aeroplane.Y-11, b, term.Yellow, term.Cyan, term.Black, 7, 8, 9)
			term.BicolorString(aeroplane.X+6, aeroplane.Y-9, fmt.Sprintf("ROW: %d", row(b)), term.White, term.Yellow, term.Black, 5, 6, 7)
			term.BicolorString(aeroplane.X+6, aeroplane.Y-8, fmt.Sprintf("COL: %d", column(b)), term.White, term.Cyan, term.Black, 5, 6, 7)
			term.Separator(aeroplane.X+4, aeroplane.Y-15, 8)
			term.Separator(aeroplane.X+21, aeroplane.Y-15, 8)
			term.Tbprint(aeroplane.X+4, aeroplane.Y-6, term.White, term.Black, "\\----------------/")

			aeroplane.TakeSeat(row(b), column(b))
		}
	}
	for i := min + 1; i < part1; i++ {
		if _, ok := seats[i]; !ok {
			part2 = i
			break
		}
	}
	if input.Visualize {

		for i := 0; i < min; i++ {
			col := i % 8
			row := (i - col) / 8
			aeroplane.SeatUnavailable(row, col)
		}
		for i := part1 + 1; i < 128*8; i++ {
			col := i % 8
			row := (i - col) / 8
			aeroplane.SeatUnavailable(row, col)
		}
		col := part2 % 8
		row := (part2 - col) / 8
		term.Tbprint(aeroplane.X+5, aeroplane.Y+21, term.White, term.Black, " YOUR SEAT")
		term.Tbprint(aeroplane.X+5, aeroplane.Y+22, term.White, term.Black, fmt.Sprintf(" ROW: %d ", row))
		term.Tbprint(aeroplane.X+5, aeroplane.Y+23, term.White, term.Black, fmt.Sprintf(" COL: %d  ", col))

		term.Tbprint(aeroplane.X+45, aeroplane.Y+21, term.White, term.Black, "UNAVAILABLE SEATS")
		minCol := (min - 1) % 8
		minRow := (min - 1 - minCol) / 8
		term.Tbprint(aeroplane.X+45, aeroplane.Y+22, term.White, term.Black, fmt.Sprintf("FROM [%d-%d] TO [%d-%d]", 0, 0, minRow, minCol))
		maxCol := (part1 + 1) % 8
		maxRow := (part1 + 1 - maxCol) / 8
		term.Tbprint(aeroplane.X+45, aeroplane.Y+23, term.White, term.Black, fmt.Sprintf("FROM [%d-%d] TO [%d-%d]", maxRow, maxCol, 127, 7))

		time.Sleep(time.Second * 15)
	}

	return part1, part2

}

func main() {
	input.ParseFlags()
	inputF := input.ReadInput(fmt.Sprintf("../../inputs/5.txt"))
	p1, p2 := solve(inputF)
	fmt.Printf("Part 1: %d, Part2: %d\n", p1, p2)
}
