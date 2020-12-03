package main

import (
	"fmt"

	"github.com/bznein/AoC2020/pkg/input"
)

type maze []string

type pos struct {
	i int
	j int
}

func (position *pos) moveBy(i, j int, limit int) {
	position.i += i
	position.j = (position.j + j) % limit
}

func explore(m maze, i, j int) int {
	res := 0

	position := pos{
		i: 0,
		j: 0,
	}

	for {
		if position.i >= len(m) {
			break
		}
		if m[position.i][position.j] == '#' {
			res++
		}
		position.moveBy(i, j, len(m[0]))
	}

	return res
}

func solve(inputF string) (int, int) {
	var m maze
	m = input.InputToStringSlice(inputF)

	part1 := explore(m, 1, 3)
	part2 := part1
	part2 *= explore(m, 1, 1)
	part2 *= explore(m, 1, 5)
	part2 *= explore(m, 1, 7)
	part2 *= explore(m, 2, 1)
	return part1, part2
}

func main() {
	inputF := input.ReadInput(fmt.Sprintf("../../inputs/3.txt"))
	p1, p2 := solve(inputF)
	fmt.Printf("Part 1: %d, Part2: %d\n", p1, p2)
}
