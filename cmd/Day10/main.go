package main

import (
	"fmt"
	"sort"

	"github.com/bznein/AoC2020/pkg/algorithm"
	"github.com/bznein/AoC2020/pkg/input"
)

func solve(inputF string) (int, int) {
	n := []int{0}
	n = append(n, input.InputToIntSlice(inputF)...)
	sort.Ints(n)
	n = append(n, n[len(n)-1]+3)
	c := make([]int, len(n))
	c[0] = 1
	diffs := map[int]int{}
	for i := range n {
		for j := algorithm.Max(0, i-3); j < i; j++ {
			if n[i]-n[j] <= 3 {
				c[i] += c[j]
			}
		}
		if i < len(n)-1 {
			diffs[n[i+1]-n[i]]++
		}
	}
	return diffs[1] * diffs[3], c[len(c)-1]
}

func main() {
	input.ParseFlags()
	inputF := input.ReadInput(fmt.Sprintf("../../inputs/10.txt"))

	p1, p2 := solve(inputF)
	fmt.Printf("Part 1: %d, Part2: %d\n", p1, p2)

}
