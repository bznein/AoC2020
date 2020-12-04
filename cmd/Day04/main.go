package main

import (
	"fmt"

	"github.com/bznein/AoC2020/pkg/input"
	"github.com/bznein/AoC2020/pkg/passport"
)

func solve(inputF string) (int, int) {
	part1 := 0
	part2 := 0

	passports := passport.StringToPassportSlice(inputF)

	for _, p := range passports {
		if p.HasAllRequiredFields() {
			part1++
			if p.IsValid() {
				part2++
			}
		}
	}

	return part1, part2
}

func main() {
	inputF := input.ReadInput(fmt.Sprintf("../../inputs/4.txt"))
	p1, p2 := solve(inputF)
	fmt.Printf("Part 1: %d, Part2: %d\n", p1, p2)
}
