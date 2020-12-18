package Day18

import (
	"strconv"
	"strings"
	"time"

	"github.com/bznein/AoC2020/pkg/input"
	"github.com/bznein/AoC2020/pkg/timing"
	"github.com/golang-collections/collections/stack"
)

type op string

const (
	add = "+"
	mul = "*"
)

func (o op) execute(a int, b int) int {
	switch o {
	case add:
		return a + b
	case mul:
		return a * b
	default:
		panic("AAAAAA")
	}
}

func matchClosedParentheses(s []string, from int) int {
	openInTheMiddle := 0
	for i := from; i < len(s); i++ {
		if s[i] == "(" {
			openInTheMiddle++
		}
		if s[i] == ")" {
			if openInTheMiddle == 0 {
				return i
			}
			openInTheMiddle--
		}
	}
	return -1
}

func subEquationSolve(s []string, from, to int) int {
	res := stack.New()
	currentOp := op("")
	for i := from; i < to; i++ {
		val, err := strconv.Atoi(s[i])
		if err == nil {
			if currentOp == "" {
				res.Push(val)
			} else {
				v1, _ := res.Pop().(int)
				res.Push(currentOp.execute(v1, val))
			}
		} else if s[i] == add || s[i] == mul {
			currentOp = op(s[i])
		} else {
			if s[i] != "(" {
				panic("Not open paren!")
			}
			nextP := matchClosedParentheses(s, i+1)
			val := subEquationSolve(s, i+1, nextP)
			if currentOp == "" {
				res.Push(val)
			} else {
				v1, _ := res.Pop().(int)
				res.Push(currentOp.execute(v1, val))
			}
			i = nextP
		}
	}
	val, _ := res.Peek().(int)
	return val
}

func Solve(inputF string) (int, int) {
	defer timing.TimeTrack(time.Now())
	part1, part2 := 0, -1

	s := input.InputToStringSlice(inputF)

	for _, ss := range s {
		s1 := strings.ReplaceAll(ss, "(", "( ")
		s1 = strings.ReplaceAll(s1, ")", " )")
		s = strings.Split(s1, " ")
		part1 += subEquationSolve(s, 0, len(s))
	}

	return part1, part2
}
