package main

import (
	"fmt"
	"sort"
	"time"

	"github.com/bznein/AoC2020/pkg/algorithm"
	"github.com/bznein/AoC2020/pkg/input"
	term "github.com/bznein/AoC2020/pkg/term"
	termT "github.com/nsf/termbox-go"
)

const (
	target = 2020
)

func solve(inputF string) (int, int) {
	part1 := -1
	ints := input.InputToIntSlice(inputF)
	sort.Ints(ints)
	if input.Visualize {
		termT.Init()
		defer termT.Close()
		termT.Clear(termT.ColorWhite, termT.ColorBlack)
		termT.Flush()
	}
	for idx, i := range ints {
		if i > target {
			continue
		}
		if input.Visualize {
			term.Tbprint(0, 0, termT.ColorWhite, termT.ColorBlack, fmt.Sprintf("Current number: %d - Searching for: %d", i, target-i))
			termT.Flush()
		}
		res := algorithm.BinarySearch(ints, target-i)
		if res != -1 && res != idx {
			part1 = ints[res] * i
			if input.Visualize {
				term.Tbprint(0, 7, termT.ColorWhite, termT.ColorBlack, fmt.Sprintf("Part 1: %d, Part2: n/a", part1))
				time.Sleep(time.Millisecond * 1)
				termT.Flush()
			}
			break
		}
	}

	//Dumb solution, terrible complexity, don't care
	for idx, i := range ints {
		for idx2, i2 := range ints {
			if i+i2 > target {
				continue
			}
			if input.Visualize {
				term.Tbprint(0, 0, termT.ColorWhite, termT.ColorBlack, fmt.Sprintf("Current pair of numbers: %d,%d - Searching for: %d", i, i2, target-i-i2))
				time.Sleep(time.Millisecond * 1)
			}
			res := algorithm.BinarySearch(ints, target-i-i2)
			if res != -1 && res != idx && res != idx2 {
				if input.Visualize {
					term.Tbprint(0, 8, termT.ColorWhite, termT.ColorBlack, fmt.Sprintf("Part 1: %d, Part2: %d", part1, ints[res]*i*i2))
					time.Sleep(time.Millisecond * 1)
					termT.Flush()
				}
				return part1, ints[res] * i * i2
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
