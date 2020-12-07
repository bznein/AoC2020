package main

import (
	"fmt"
	"sort"

	"github.com/bznein/AoC2020/pkg/algorithm"
	"github.com/bznein/AoC2020/pkg/input"
	"github.com/bznein/AoC2020/pkg/visualize"
)

const (
	target = 2020
)

func solve(inputF string) (int, int) {
	part1 := -1
	ints := input.InputToIntSlice(inputF)
	sort.Ints(ints)
	visualize.Init()
	defer visualize.Close()
	for idx, i := range ints {
		if i > target {
			continue
		}
		day.ShowSearchLine(target, i, nil)
		res := algorithm.BinarySearch(ints, target-i)
		if res != -1 && res != idx {
			part1 = ints[res] * i
			day.ShowCurrentResult(part1, nil)
			break
		}
	}

	//Dumb solution, terrible complexity, don't care
	for idx, i := range ints {
		for idx2, i2 := range ints {
			if i+i2 > target {
				continue
			}
			day.ShowSearchLine(target, i, &i2)
			res := algorithm.BinarySearch(ints, target-i-i2)
			if res != -1 && res != idx && res != idx2 {
				p2 := ints[res] * i * i2
				day.ShowCurrentResult(part1, &p2)
				return part1, p2
			}
		}
	}
	return part1, -1
}

func main() {
	input.ParseFlags()
	inputF := input.ReadInput(fmt.Sprintf("../../inputs/1.txt"))
	p1, p2 := solve(inputF)
	fmt.Printf("Part 1: %d, Part2: %d\n", p1, p2)
}
