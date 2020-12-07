package visualize

import (
	"fmt"
	"time"

	"github.com/bznein/AoC2020/pkg/input"
	term "github.com/bznein/AoC2020/pkg/term"
	"github.com/nsf/termbox-go"
)

func ShowSearchLine(target int, i int, i2 *int) {
	if input.Visualize {
		if i2 == nil {
			term.Tbprint(0, 0, term.White, term.Black, fmt.Sprintf("Current number: %d - Searching for: %d", i, target-i))
		} else {

			term.Tbprint(0, 0, term.White, term.Black, fmt.Sprintf("Current pair of numbers: %d,%d - Searching for: %d", i, *i2, target-i-*i2))
		}
	}
}

func ShowCurrentResult(part1 int, part2 *int) {
	if input.Visualize {
		if part2 == nil {
			term.Tbprint(0, 7, term.White, term.Black, fmt.Sprintf("Part 1: %d, Part2: n/a", part1))
		} else {
			term.Tbprint(0, 7, term.White, term.Black, fmt.Sprintf("Part 1: %d, Part2: %d", part1, part2))
		}
	}
}

func ShowResult(found bool) {
	var res string
	var c termbox.Attribute
	if found {
		res = "Found!"
		c = term.Green
	} else {
		res = "not found :("
		c = term.Red
	}
	if input.Visualize {
		term.Tbprint(0, 5, c, termbox.ColorBlack, res)
		time.Sleep(time.Millisecond * input.Delay)
	}
}

func ShowBinarySearch(firstColour termbox.Attribute, secondColour termbox.Attribute, lower []int, mid int, upper []int) {
	if input.Visualize {
		term.ClearLine(3)
		term.Tbprint(0, 3, firstColour, termbox.ColorBlack, fmt.Sprintf("%v", lower))
		term.Tbprint(33, 3, termbox.ColorWhite, termbox.ColorBlack, fmt.Sprintf("%d", mid))
		term.Tbprint(45, 3, secondColour, termbox.ColorBlack, fmt.Sprintf("%v", upper))
		term.ClearLine(5)
		time.Sleep(time.Millisecond * input.Delay)
	}
}
