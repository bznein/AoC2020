package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func solve(input string) (int, int) {
	part1 := 0
	part2 := 0
	passwords := strings.Split(input, "\n")
	passwords = passwords[:len(passwords)-1]

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
	input := readInput(fmt.Sprintf("../../inputs/2.txt"))
	p1, p2 := solve(input)
	fmt.Printf("Part 1: %d, Part2: %d\n", p1, p2)
}
