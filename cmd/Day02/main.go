package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/bznein/AoC2020/pkg/input"
	"github.com/bznein/AoC2020/pkg/term"
)

const (
	g = term.Green
	b = term.Black
	r = term.Red
	w = term.White
)

func solve(inputF string) (int, int) {
	if input.Visualize {
		term.Init()
		defer term.Close()
	}
	part1 := 0
	part2 := 0
	passwords := input.InputToStringSlice(inputF)

	for _, pwd := range passwords {
		if input.Visualize {
			term.Tbprint(10, 8, w, b, fmt.Sprintf("Part 1: %d, Part 2: %d", part1, part2))

		}
		parts := strings.Split(pwd, " ")
		nums := strings.Split(parts[0], "-")
		min, _ := strconv.Atoi(nums[0])
		max, _ := strconv.Atoi(nums[1])

		letter := parts[1][0]
		pw := parts[2]
		occurrences := strings.Count(pw, string(letter))
		if occurrences >= min && occurrences <= max {
			part1++
			if input.Visualize {
				term.ClearLine(3)
				term.Separator(40, 0, 5)
				term.Tbprint(10, 3, g, b, pw)
				term.Tbprint(35, 1, g, b, nums[0])
				term.Tbprint(35, 2, g, b, "<=")
				term.Tbprint(35, 3, g, b, fmt.Sprintf("%d", occurrences))
				term.Tbprint(35, 4, g, b, "<=")
				term.Tbprint(35, 5, g, b, nums[1])
			}
		} else if input.Visualize {
			lowerC := g
			upperC := r
			if occurrences < min {
				lowerC = r
				upperC = g
			}
			term.ClearLine(3)
			term.Separator(40, 0, 5)
			term.Tbprint(10, 3, r, b, pw)
			term.Tbprint(35, 1, lowerC, b, nums[0])
			term.Tbprint(35, 2, lowerC, b, "<=")
			term.Tbprint(35, 3, r, b, fmt.Sprintf("%d", occurrences))
			term.Tbprint(35, 4, upperC, b, "<=")
			term.Tbprint(35, 5, upperC, b, nums[1])
		}
		if (pw[min-1] == letter && pw[max-1] != letter) || (pw[min-1] != letter && pw[max-1] == letter) {
			correctIndex := min - 1
			if pw[max-1] == letter {
				correctIndex = max - 1
			}
			if input.Visualize {
				term.BicolorString(45, 3, pw, w, g, b, correctIndex)
			}
			part2++
		} else if input.Visualize {
			term.BicolorString(45, 3, pw, w, r, b, min-1, max-1)
		}
		time.Sleep(time.Millisecond * input.Delay)

	}
	return part1, part2

}

func main() {
	input.ParseFlags()
	inputF := input.ReadInput(fmt.Sprintf("../../inputs/2.txt"))
	p1, p2 := solve(inputF)
	fmt.Printf("Part 1: %d, Part2: %d\n", p1, p2)
}
