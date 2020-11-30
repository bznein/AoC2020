package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

var flagtests = []struct {
	in  string
	out int
}{
	{"%a", 2},
	{"%a", 3},
}

func tests() bool {
	passes := true
	for _, tt := range flagtests {
		res := solve(tt.in)
		if res != tt.out {
			fmt.Printf("Expected %v, got %v\n", tt.out, res)
			passes = false
		}
	}
	return passes
}

func solve(input string) int {
	return 2
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
	if !tests() {
		fmt.Println("Tests are not passing!")
		os.Exit(1)
	}
	input := readInput("./input.txt")
	fmt.Printf("Solve: %v", solve(input))
}
