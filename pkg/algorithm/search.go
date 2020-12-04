package algorithm

import (
	"fmt"
	"time"

	"github.com/bznein/AoC2020/pkg/input"
	"github.com/bznein/AoC2020/pkg/term"
	"github.com/nsf/termbox-go"
)

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func BinarySearch(a []int, search int) int {
	mid := len(a) / 2

	lower := max(0, mid-5)
	upper := min(mid+5, len(a))
	switch {
	case len(a) == 0:
		if input.Visualize {
			term.Tbprint(0, 5, termbox.ColorRed, termbox.ColorBlack, "not found :(")
			termbox.Flush()
			time.Sleep(time.Millisecond * input.Delay)
		}
		return -1 // not found
	case a[mid] > search:
		if input.Visualize {
			term.ClearLine(3)
			term.Tbprint(0, 3, termbox.ColorRed, termbox.ColorBlack, fmt.Sprintf("%v", a[lower:mid]))
			term.Tbprint(33, 3, termbox.ColorWhite, termbox.ColorBlack, fmt.Sprintf("%d", a[mid]))
			term.Tbprint(45, 3, termbox.ColorGreen, termbox.ColorBlack, fmt.Sprintf("%v", a[mid+1:upper]))
			term.ClearLine(5)
			time.Sleep(time.Millisecond * input.Delay)
		}
		return BinarySearch(a[:mid], search)
	case a[mid] < search:
		result := BinarySearch(a[mid+1:], search)
		if input.Visualize {
			term.ClearLine(3)
			term.Tbprint(0, 3, termbox.ColorGreen, termbox.ColorBlack, fmt.Sprintf("%v", a[lower:mid]))
			term.Tbprint(33, 3, termbox.ColorWhite, termbox.ColorBlack, fmt.Sprintf("%d", a[mid]))
			term.Tbprint(45, 3, termbox.ColorRed, termbox.ColorBlack, fmt.Sprintf("%v", a[mid+1:upper]))
			term.ClearLine(5)
			time.Sleep(time.Millisecond * input.Delay)
		}
		if result >= 0 { // if anything but the -1 "not found" result
			return result + mid + 1
		}
		return -1
	default: // a[mid] == search
		if input.Visualize {
			term.Tbprint(0, 5, termbox.ColorGreen, termbox.ColorBlack, "Found!")
		}
		return mid
	}
}
