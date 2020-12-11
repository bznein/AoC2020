package algorithm

import (
	"github.com/bznein/AoC2020/pkg/term"
	visualize "github.com/bznein/AoC2020/pkg/visualize/Day01"
)

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

// Note: assumes search is NOT in the array
func ClosestTwoElements(a []int, search int) (int, int) {
	if len(a) == 0 {
		return -1, -1
	}
	if len(a) == 1 {
		if a[0] < search {
			return a[0], -1
		}
		return -1, a[0]
	}
	mid := len(a) / 2
	midV := a[mid]
	midPrevV := a[mid-1]

	if midPrevV < search && midV > search {
		return midPrevV, midV
	}

	if midPrevV < search && mid < search {
		return ClosestTwoElements(a[mid+1:], search)
	}
	return ClosestTwoElements(a[:mid], search)
}
