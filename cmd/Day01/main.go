package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"sort"
	"strconv"
	"strings"
)

func binarySearch(a []int, search int) int {
	mid := len(a) / 2
	switch {
	case len(a) == 0:
		return -1 // not found
	case a[mid] > search:
		return binarySearch(a[:mid], search)
	case a[mid] < search:
		result := binarySearch(a[mid+1:], search)
		if result >= 0 { // if anything but the -1 "not found" result
			return result + mid + 1
		}
		return -1
	default: // a[mid] == search
		return mid
	}
}

func solve(input string) (int, int) {
	part1 := -1
	part2 := -1
	ints := []int{}
	for _, s := range strings.Split(input, "\n") {
		i, _ := strconv.Atoi(s)
		ints = append(ints, i)
	}
	ints = ints[:len(ints)-1]
	sort.Ints(ints)
	for idx, i := range ints {
		res := binarySearch(ints, 2020-i)
		if res != -1 && res != idx {
			part1 = ints[res] * i
		}
	}

	//Dumb solution, terrible complexity, don't care
	for idx, i := range ints {
		for idx2, i2 := range ints {
			res := binarySearch(ints, 2020-i-i2)
			if res != -1 && res != idx && res != idx2 {
				part2 = ints[res] * i * i2
			}
		}
	}
	return part1, part2
}

func readInput(path string) string {
	// Read entire file content, giving us little control but
	// making it very simple. No need to close the file.
	content, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}

	// Convert []byte to string and print to screen
	text := string(content)
	return text
}

func main() {
	input := readInput(fmt.Sprintf("../../inputs/1.txt"))
	p1, p2 := solve(input)
	fmt.Printf("Part 1: %d, Part2: %d\n", p1, p2)
}
