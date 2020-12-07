package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/bznein/AoC2020/pkg/input"
)

type bag struct {
	attribute string
	colour    string
}

type rule map[bag][]containRule

type containRule struct {
	n int
	b bag
}

func target() bag {
	return bag{
		attribute: "shiny",
		colour:    "gold",
	}
}

func (r rule) leadsToTarget(from bag) bool {
	v, ok := r[from]
	if !ok {
		return false
	}
	for _, b := range v {
		if b.b == target() {
			return true
		}
		if r.leadsToTarget(b.b) {
			return true
		}
	}
	return false
}

func (r rule) exploreToEnd(from bag) int {
	n := 1
	v, ok := r[from]
	if !ok {
		return n
	}
	for _, b := range v {
		n += b.n * r.exploreToEnd(b.b)
	}
	return n
}

func solve(inputF string) (int, int) {
	part1 := 0
	part2 := 0

	bags := input.InputToStringSlice(inputF)
	contains := rule{}

	for _, b := range bags {
		splitted := strings.Split(b, " ")
		currentBag := bag{
			attribute: splitted[0],
			colour:    splitted[1],
		}
		commas := strings.Count(b, ",")
		if strings.Contains(b, " no other") {
			continue
		}
		for i := 0; i <= commas; i++ {
			n, _ := strconv.Atoi(splitted[4+(4*i)])
			otherB := bag{
				attribute: splitted[5+(4*i)],
				colour:    splitted[6+(4*i)],
			}
			contains[currentBag] = append(contains[currentBag], containRule{
				n: n,
				b: otherB,
			})
		}
	}

	for k := range contains {
		if contains.leadsToTarget(k) {
			part1++
		}
		if k == target() {
			part2 = contains.exploreToEnd(k) - 1
		}
	}

	return part1, part2
}

func main() {
	input.ParseFlags()
	inputF := input.ReadInput(fmt.Sprintf("../../inputs/7.txt"))
	p1, p2 := solve(inputF)
	fmt.Printf("Part 1: %d, Part2: %d\n", p1, p2)
}
