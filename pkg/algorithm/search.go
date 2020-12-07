package algorithm

import (
	"github.com/bznein/AoC2020/pkg/term"
	visualize "github.com/bznein/AoC2020/pkg/visualize/Day01"
)

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func Max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func BinarySearch(a []int, search int) int {
	mid := len(a) / 2

	lower := Max(0, mid-5)
	upper := Min(mid+5, len(a))
	switch {
	case len(a) == 0:
		visualize.ShowResult(false)
		return -1 // not found
	case a[mid] > search:
		visualize.ShowBinarySearch(term.Red, term.Green, a[lower:mid], mid, a[mid+1:upper])
		return BinarySearch(a[:mid], search)
	case a[mid] < search:
		result := BinarySearch(a[mid+1:], search)
		visualize.ShowBinarySearch(term.Green, term.Red, a[lower:mid], mid, a[mid+1:upper])
		if result >= 0 { // if anything but the -1 "not found" result
			return result + mid + 1
		}
		return -1
	default: // a[mid] == search
		visualize.ShowResult(true)
		return mid
	}
}
