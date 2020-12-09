package main

import (
	"fmt"

	"github.com/bznein/AoC2020/pkg/algorithm"
	"github.com/bznein/AoC2020/pkg/input"
)

type preambleSums []int

const (
	preambleLength = 5
)

func getPreambleSums(n []int) preambleSums {
	sums := preambleSums{}
	for i, n1 := range n {
		for j, n2 := range n {
			if i != j {
				sums = append(sums, n1+n2)
			}
		}
	}
	return sums
}

func (p preambleSums) contains(n int) bool {
	for _, v := range p {
		if v == n {
			return true
		}
	}
	return false
}

func consecutiveSumTo(numbers []int, from int, n int) int {
	size := 0
	sum := 0
	for {
		if from+size >= len(numbers) {
			return 0
		}
		sum += numbers[from+size]
		if sum == n {
			return size
		}
		if sum > n {
			return 0
		}
		size++
	}
}

func minMaxSlice(slice []int) (int, int) {
	min := slice[0]
	max := slice[0]
	for _, v := range slice {
		min = algorithm.Min(min, v)
		max = algorithm.Max(max, v)
	}
	return min, max
}

func solve(inputF string) (int, int) {
	part1 := int(0)
	part2 := int(0)

	numbers := input.InputToIntSlice(inputF)

	for i := preambleLength; i < len(numbers); i++ {
		p := getPreambleSums(numbers[i-preambleLength : i])
		if !p.contains(numbers[i]) {
			part1 = numbers[i]
			break
		}
	}

	for i := range numbers {
		if size := consecutiveSumTo(numbers, i, part1); size != 0 {
			min, max := minMaxSlice(numbers[i : i+size+1])
			part2 = min + max
			break
		}
	}

	return part1, part2
}

func main() {
	input.ParseFlags()
	inputF := input.ReadInput(fmt.Sprintf("../../inputs/9.txt"))
	if input.Visualize {
		visualSolve(inputF)
	} else {
		p1, p2 := solve(inputF)
		fmt.Printf("Part 1: %d, Part2: %d\n", p1, p2)
	}
}
