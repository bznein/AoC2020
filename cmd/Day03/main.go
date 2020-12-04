package main

import (
	"fmt"
	"time"

	"github.com/bznein/AoC2020/pkg/input"
	"github.com/bznein/AoC2020/pkg/strings"
	"github.com/bznein/AoC2020/pkg/term"
)

type maze []string

type pos struct {
	i int
	j int
}

var try int
var part1 int
var part2 int

func (position *pos) moveBy(i, j int, limit int) {
	position.i += i
	position.j = (position.j + j) % limit
}

func explore(m maze, i, j int) int {
	if input.Visualize {
		term.Init()
		defer term.Close()
		term.Tbprint(60, 20, term.White, term.Black, fmt.Sprintf("Part 1: %d", part1))
		term.Tbprint(60, 21, term.White, term.Black, fmt.Sprintf("Part 2: %d", part2))
	}
	res := 0

	position := pos{
		i: 0,
		j: 0,
	}

	curColor := term.White
	for {
		term.Tbprint(60, try+10, curColor, term.Black, fmt.Sprintf("Trees encountered with slope (%d,%d) = %d", j, i, res))
		// TODO make the 2d maze multicolour (see term.go)
		if input.Visualize {
			c := m[position.i][position.j]
			m[position.i] = strings.ReplaceAtIndex(m[position.i], '@', position.j)
			// TODO see if all of this fits into the terminal, show only a part otherwise
			term.StringSlice(10, 3, term.White, term.Black, m)
			m[position.i] = strings.ReplaceAtIndex(m[position.i], rune(c), position.j)

		}
		if position.i >= len(m) {
			break
		}
		if m[position.i][position.j] == '#' {
			res++
			curColor = term.Red
		} else {
			curColor = term.White
		}
		position.moveBy(i, j, len(m[0]))
		time.Sleep(time.Millisecond * input.Delay)
	}

	return res
}

func solve(inputF string) (int, int) {
	var m maze
	m = input.InputToStringSlice(inputF)
	part1 = 0
	part2 = 0
	try = 0
	part1 = explore(m, 1, 3)
	part2 = part1
	try++
	part2 *= explore(m, 1, 1)
	try++
	part2 *= explore(m, 1, 5)
	try++
	part2 *= explore(m, 1, 7)
	try++
	part2 *= explore(m, 2, 1)
	return part1, part2
}

func main() {
	input.ParseFlags()
	inputF := input.ReadInput(fmt.Sprintf("../../inputs/3.txt"))
	p1, p2 := solve(inputF)
	fmt.Printf("Part 1: %d, Part2: %d\n", p1, p2)
}
