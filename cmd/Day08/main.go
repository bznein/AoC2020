package main

import (
	"fmt"
	"strings"

	"github.com/bznein/AoC2020/pkg/console"
	"github.com/bznein/AoC2020/pkg/input"
	visualize "github.com/bznein/AoC2020/pkg/visualize"
)

func changeIthNopOrJmp(target int, slice []string) []string {
	found := 0
	retVal := make([]string, len(slice))
	copy(retVal, slice)
	mapping := map[string]string{"nop": "jmp", "jmp": "nop"}
	for i, s := range retVal {
		split := strings.Split(s, " ")
		ss := split[0]
		if ss == "nop" || ss == "jmp" {
			if found == target {
				retVal[i] = mapping[ss] + " " + split[1]
				return retVal
			}
			found++
		}
	}
	return retVal
}

func solve(inputF string) (int64, int64) {
	part1 := int64(0)
	part2 := int64(0)

	visualize.Init()
	defer visualize.Close()
	a := input.InputToStringSlice(inputF)
	ex := console.New(a)
	ex.Run()
	part1 = ex.Peek()

	tries := 0
	for {
		ex = console.New(changeIthNopOrJmp(tries, a))
		err := ex.Run()
		if err == nil {
			part2 = ex.Peek()
			return part1, part2
		}
		tries++
	}
}

func main() {
	input.ParseFlags()
	inputF := input.ReadInput(fmt.Sprintf("../../inputs/8.txt"))
	p1, p2 := solve(inputF)
	fmt.Printf("Part 1: %d, Part2: %d\n", p1, p2)
}
