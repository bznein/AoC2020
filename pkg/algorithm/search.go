package algorithm

func BinarySearch(a []int, search int) int {
	mid := len(a) / 2
	switch {
	case len(a) == 0:
		return -1 // not found
	case a[mid] > search:
		return BinarySearch(a[:mid], search)
	case a[mid] < search:
		result := BinarySearch(a[mid+1:], search)
		if result >= 0 { // if anything but the -1 "not found" result
			return result + mid + 1
		}
		return -1
	default: // a[mid] == search
		return mid
	}
}
