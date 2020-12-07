package main

import (
	"fmt"

	"github.com/bznein/AoC2020/pkg/input"
)

func solve(inputF string) (int, int) {
	part1 := 0
	part2 := 0

	a := input.InputToStringSlice(inputF)
	a := input.InputToIntSlice(inputF)
	a := input.

	return part1, part2
}

func main() {
	input.ParseFlags()
	inputF := input.ReadInput(fmt.Sprintf("../../inputs/7.txt"))
	p1, p2 := solve(inputF)
	fmt.Printf("Part 1: %d, Part2: %d\n", p1, p2)
}
