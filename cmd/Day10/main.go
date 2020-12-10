package main

import (
	"fmt"
	"sort"

	"github.com/bznein/AoC2020/pkg/input"
)

func byOneTimesByThreeDifferences(n []int) int {
	byOne := 0
	byThree := 0
	for i := 0; i < len(n)-1; i++ {
		if n[i+1]-n[i] == 1 {
			byOne++
		} else if n[i+1]-n[i] == 3 {
			byThree++
		}
	}
	return byOne * byThree
}

func possibleArrangements(n []int) int {
	c := make([]int, len(n))
	c[0] = 1
	for i := range n {
		for j := 0; j < i; j++ {
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
