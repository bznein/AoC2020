package Day08

import (
	"strings"
	"time"

	"github.com/bznein/AoC2020/pkg/console"
	"github.com/bznein/AoC2020/pkg/input"
	"github.com/bznein/AoC2020/pkg/timing"
	visualize "github.com/bznein/AoC2020/pkg/visualize"
)

var mapping = map[string]string{"nop": "jmp", "jmp": "nop"}

func replaceInstruction(source string) string {
	split := strings.Split(source, " ")
	ss := split[0]
	return mapping[ss] + " " + split[1]
}

func listNopOrJmp(slice []string) []int {
	retVal := make([]int, 0, len(slice))
	for i, s := range slice {
		split := strings.Split(s, " ")
		ss := split[0]
		if ss == "nop" || ss == "jmp" {
			retVal = append(retVal, i)
		}
	}
	return retVal
}

func Solve(inputF string) (int64, int64) {
	defer timing.TimeTrack(time.Now())
	part1 := int64(0)
	part2 := int64(0)

	visualize.Init()
	defer visualize.Close()
	a := input.InputToStringSlice(inputF)
	ex := console.New(a)
	ex.Run()
	part1 = ex.Peek()

	originalList := listNopOrJmp(a)
	tries := 0
	for {
		originalStr := a[originalList[tries]]
		a[originalList[tries]] = replaceInstruction(originalStr)
		ex = console.New(a)
		err := ex.Run()
		if err == nil {
			part2 = ex.Peek()
			return part1, part2
		}
		a[originalList[tries]] = originalStr
		tries++
	}
}
