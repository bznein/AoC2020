package main

import (
	"fmt"
	"sort"

	"github.com/bznein/AoC2020/pkg/algorithm"
	"github.com/bznein/AoC2020/pkg/input"
)

func byOneTimesByThreeDifferences(n []int) int {
	diffs := map[int]int{}
	for i := 0; i < len(n)-1; i++ {
		diffs[n[i+1]-n[i]]++
	}
	return diffs[1] * diffs[3]
}

func possibleArrangements(n []int) int {
	c := make([]int, len(n))
	c[0] = 1
	for i := range n {
		for j := algorithm.Max(0, i-3); j < i; j++ {
			if n[i]-n[j] <= 3 {
				c[i] += c[j]
			}
		}
	}
	return c[len(c)-1]
}

func solve(inputF string) (int, int) {
	part1 := int(0)
	part2 := int(0)

	n := []int{0}
	n = append(n, input.InputToIntSlice(inputF)...)
	sort.Ints(n)
	n = append(n, n[len(n)-1]+3)
	part1 = byOneTimesByThreeDifferences(n)
	part2 = possibleArrangements(n)
	return part1, part2
}

func main() {
	input.ParseFlags()
	inputF := input.ReadInput(fmt.Sprintf("../../inputs/10.txt"))

	p1, p2 := solve(inputF)
	fmt.Printf("Part 1: %d, Part2: %d\n", p1, p2)

}
