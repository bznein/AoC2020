package main

import (
	"fmt"
	"sort"

	"github.com/bznein/AoC2020/pkg/algorithm"
	"github.com/bznein/AoC2020/pkg/input"
)

func solve(inputF string) (int, int) {
	part1 := -1
	part2 := -1
	ints := input.InputToIntSlice(inputF)
	sort.Ints(ints)
	for idx, i := range ints {
		res := algorithm.BinarySearch(ints, 2020-i)
		if res != -1 && res != idx {
			part1 = ints[res] * i
		}
	}

	//Dumb solution, terrible complexity, don't care
	for idx, i := range ints {
		for idx2, i2 := range ints {
			res := algorithm.BinarySearch(ints, 2020-i-i2)
			if res != -1 && res != idx && res != idx2 {
				part2 = ints[res] * i * i2
			}
		}
	}
	return part1, part2
}

func main() {
	inputF := input.ReadInput(fmt.Sprintf("../../inputs/1.txt"))
	p1, p2 := solve(inputF)
	fmt.Printf("Part 1: %d, Part2: %d\n", p1, p2)
}
