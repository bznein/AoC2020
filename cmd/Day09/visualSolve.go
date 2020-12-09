package main

import (
	"fmt"
	"time"

	"github.com/bznein/AoC2020/pkg/input"
	"github.com/bznein/AoC2020/pkg/term"
)

func visualSolvePart1(n []int) int {
	ret := 0
TargetSearch:
	for i := preambleLength; i < len(n); i++ {
		term.Tbprint(10, 5, term.White, term.Black, fmt.Sprintf("TARGET %d                   ", n[i]))
		for j := i - preambleLength; j < i; j++ {
			for k := j + 1; k < i; k++ {
				res := n[j] + n[k]
				term.Tbprint(15, 8, term.White, term.Black, fmt.Sprintf("%d           ", n[j]))
				term.Tbprint(15, 9, term.White, term.Black, "+")
				term.Tbprint(15, 10, term.White, term.Black, fmt.Sprintf("%d           ", n[k]))
				term.Tbprint(15, 11, term.White, term.Black, "=")
				term.Tbprint(15, 12, term.White, term.Black, fmt.Sprintf("%d           ", res))
				if res == n[i] {
					term.Tbprint(15, 15, term.Red, term.Black, "NUMBERS SUM TO THE TARGET, CONTINUING WITH THE NEXT ONE...        ")
					time.Sleep(time.Millisecond * input.Delay * 5)
					continue TargetSearch
				} else {
					term.Tbprint(15, 15, term.Green, term.Black, "NUMBERS DO NOT SUM TO THE TARGET, MAYBE THIS IS THE RIGHT ONE?         ")
					time.Sleep(time.Millisecond * input.Delay)
				}
				time.Sleep(time.Millisecond * input.Delay)
			}
		}
		term.Tbprint(15, 15, term.Green, term.Black, fmt.Sprintf("NO PAIR FOUND THAT SUMS TO THIS NUMBER, THE SOLUTION FOR PART 1 IS %d", n[i]))
		ret = n[i]
		break
	}

	time.Sleep(time.Second * 5)
	return ret
}

func visualSolvePart2(n []int, target int) {

	term.Tbprint(10, 25, term.White, term.Black, fmt.Sprintf("PART 2 TARGET %d", target))

External:
	for i := 0; i < len(n)-preambleLength; i++ {
		sum := 0
		totNums := 0
		term.ClearArea(10, 27, 100, 100)
		for {
			sum += n[i+totNums]

			if totNums > 0 {
				term.Tbprint(10, 30+(2*totNums), term.White, term.Black, "+")
			}
			term.Tbprint(10, 30+(2*totNums+1), term.White, term.Black, fmt.Sprintf("%d", n[i+totNums]))
			term.Tbprint(10, 30, term.White, term.Black, fmt.Sprintf("SUM: %d", sum))
			if sum > target {
				term.Tbprint(10, 27, term.Red, term.Black, "SUM OVER TARGET, LET'S RESTART            ")
				time.Sleep(time.Millisecond * input.Delay)
				break
			} else if sum == target {
				term.Tbprint(10, 27, term.Green, term.Black, "SUM REACHED, NOW I HAVE TO ADD MIN AND MAX")
				break External
			}
			time.Sleep(time.Millisecond * input.Delay)
			totNums++
		}
	}
	time.Sleep(time.Second * 5)
}

func visualSolve(inputF string) {

	term.Init()
	defer term.Close()
	n := input.InputToIntSlice(inputF)
	visualSolvePart2(n, visualSolvePart1(n))
}
