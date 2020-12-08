package console

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/bznein/AoC2020/pkg/input"
)

type instruction struct {
	command  string
	argument int
}

const (
	noop = "nop"
	jump = "jmp"
	acc  = "acc"
)

type Executor struct {
	ip          int64
	accumulator int64
	execCount   map[int64]int

	instructions []instruction
}

func New(instructions []string) Executor {
	ex := Executor{
		ip:           0,
		accumulator:  0,
		instructions: []instruction{},
		execCount:    map[int64]int{},
	}

	for _, i := range instructions {
		cmd := strings.Split(i, " ")
		arg, _ := strconv.Atoi(cmd[1])
		ex.instructions = append(ex.instructions, instruction{
			command:  cmd[0],
			argument: arg,
		})
	}

	return ex
}

func (ex Executor) getInstruction(index int64) instruction {
	return ex.instructions[index]
}

func (ex *Executor) Run() error {
	for {
		if input.Visualize {
			ex.Print()
		}
		time.Sleep(input.Delay * time.Millisecond)

		if _, ok := ex.execCount[ex.ip]; ok {
			if input.Visualize {
				ex.InfiniteLoopDetected()
			}
			time.Sleep(input.Delay * time.Millisecond * 2000)
			return fmt.Errorf("Infinite Loop Detected!")
		}
		if ex.ip >= int64(len(ex.instructions)) {
			return nil
		}
		ex.execCount[ex.ip]++

		ins := ex.getInstruction(ex.ip)

		switch ins.command {
		case noop:
			ex.ip += 1
		case acc:
			ex.accumulator += int64(ins.argument)
			ex.ip += 1
		case jump:
			ex.ip += int64(ins.argument)
		}
	}
}

func (ex Executor) Peek() int64 {
	return ex.accumulator
}
