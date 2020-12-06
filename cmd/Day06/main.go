package main

import (
	"fmt"

	"github.com/bznein/AoC2020/pkg/input"
)

type answer string
type group []answer

func inputToGroups(inputD string) []group {
	result := []group{}
	stringSlice := input.InputToStringSlice(inputD)
	emptyGroup := group{}
	for _, s := range stringSlice {
		if s == "" {
			result = append(result, emptyGroup)
			emptyGroup = group{}
			continue
		}
		emptyGroup = append(emptyGroup, answer(s))

	}
	result = append(result, emptyGroup)
	return result
}

func (g group) totalYesAndAnsweredByEveryone() (int, int) {
	res := 0
	p2 := 0

	yes := map[rune]int{}
	for _, answer := range g {
		for _, ch := range answer {
			if _, ok := yes[ch]; !ok {
				res++
				yes[ch] = 1
			} else {
				yes[ch] += 1
			}
			if yes[ch] == len(g) {
				p2++
			}
		}
	}

	return res, p2
}

func solve(inputF string) (int, int) {
	part1 := 0
	part2 := 0

	groups := inputToGroups(inputF)
	for _, g := range groups {
		yes, everyone := g.totalYesAndAnsweredByEveryone()
		part1 += yes
		part2 += everyone
	}
	return part1, part2
}

func main() {
	input.ParseFlags()
	inputF := input.ReadInput(fmt.Sprintf("../../inputs/6.txt"))
	p1, p2 := solve(inputF)
	fmt.Printf("Part 1: %d, Part2: %d\n", p1, p2)
}
