package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/bznein/AoC2020/pkg/input"
)

func solve(inputF string) (int, int) {
	part1 := 0
	part2 := 0
	passwords := input.InputToStringSlice(inputF)

	for _, pwd := range passwords {
		parts := strings.Split(pwd, " ")
		nums := strings.Split(parts[0], "-")
		min, _ := strconv.Atoi(nums[0])
		max, _ := strconv.Atoi(nums[1])
		letter := parts[1][0]
		pw := parts[2]
		occurrences := strings.Count(pw, string(letter))
		if occurrences >= min && occurrences <= max {
			part1++
		}
		if pw[min-1] == letter && pw[max-1] != letter {
			part2++
		} else if pw[min-1] != letter && pw[max-1] == letter {
			part2++
		}

	}
	return part1, part2

}

func main() {
	inputF := input.ReadInput(fmt.Sprintf("../../inputs/2.txt"))
	p1, p2 := solve(inputF)
	fmt.Printf("Part 1: %d, Part2: %d\n", p1, p2)
}
