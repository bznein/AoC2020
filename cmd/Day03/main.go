package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

type maze []string

type pos struct {
	i int
	j int
}

func (position *pos) moveBy(i, j int, limit int) {
	position.i += i
	position.j = (position.j + j) % limit
}

func explore(m maze, i, j int) int {
	res := 0

	position := pos{
		i: 0,
		j: 0,
	}

	for {
		if position.i >= len(m) {
			break
		}
		if m[position.i][position.j] == '#' {
			res++
		}
		position.moveBy(i, j, len(m[0]))
	}

	return res
}

func solve(input string) (int, int) {
	var m maze
	m = strings.Split(input, "\n")

	m = m[:len(m)-1]
	part1 := explore(m, 1, 3)
	part2 := part1
	part2 *= explore(m, 1, 1)
	part2 *= explore(m, 1, 5)
	part2 *= explore(m, 1, 7)
	part2 *= explore(m, 2, 1)
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
	input := readInput(fmt.Sprintf("../../inputs/3.txt"))
	p1, p2 := solve(input)
	fmt.Printf("Part 1: %d, Part2: %d\n", p1, p2)
}
