package algorithm

func minMaxInSlice(n []int) (int, int) {
	if len(n) == 0 {
		return 0, 0
	}
	min := n[0]
	max := n[0]
	for _, k := range n {
		min = Min(min, k)
		max = Max(max, k)
	}
	return min, max
}

func CountingSort(n []int) []int {
	offset, max := minMaxInSlice(n)

	sortedSlice := make([]int, len(n))
	counts := make([]int, max-offset+1)

	for _, v := range n {
		counts[v-offset]++
	}

	idx := 0
	for k, v := range counts {
		for i := 0; i < v; i++ {
			sortedSlice[idx] = k + offset
			idx++
		}
	}
	return sortedSlice
}
