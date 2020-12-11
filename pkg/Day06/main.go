package Day06

import (
	"fmt"
	"time"

	"github.com/bznein/AoC2020/pkg/input"
	"github.com/bznein/AoC2020/pkg/term"
	"github.com/bznein/AoC2020/pkg/timing"
	"github.com/nsf/termbox-go"
)

type answer string
type group []answer

func inputToGroups(inputD string) []group {
	result := []group{}
	stringSlice := input.InputToStringSlice(inputD)
	emptyGroup := group{}
	for _, s := range stringSlice {
		if s == "" {
			result = append(result, emptyGroup)
			emptyGroup = group{}
			continue
		}
		emptyGroup = append(emptyGroup, answer(s))

	}
	result = append(result, emptyGroup)
	return result
}

func (g group) totalYesAndAnsweredByEveryone() (int, int) {
	res := 0
	p2 := 0

	yes := map[rune]int{}
	for stick, answer := range g {
		for _, ch := range answer {
			if _, ok := yes[ch]; !ok {
				res++
				yes[ch] = 1
			} else {
				yes[ch] += 1
			}
			if input.Visualize {
				printStickFigure(59+yes[ch]*4, 2+(3*int(ch-'a')), colours()[stick%5])
			}
			if yes[ch] == len(g) {
				p2++
			}
			time.Sleep(time.Millisecond * input.Delay)
		}
	}

	return res, p2
}

func printBoard() {
	term.Separator(50, 0, 100)

	term.Tbprint(10, 0, term.Red, term.Black, "GROUP #")
	term.Tbprint(55, 0, term.Red, term.Black, "QUESTIONS")

	for i := 0; i < 26; i++ {
		term.Tbprint(55, 3+(3*i), term.White, term.Black, string(rune('a'+i))+": ")
	}
}

func printStickFigure(x, y int, color termbox.Attribute) {
	term.Tbprint(x, y, color, term.Black, " o ")
	term.Tbprint(x, y+1, color, term.Black, "/|\\ ")
	term.Tbprint(x, y+2, color, term.Black, "/ \\ ")
}

func colours() []termbox.Attribute {
	return []termbox.Attribute{term.Red, term.Green, term.Cyan, term.Yellow, term.White}
}
func printStickFigures(n int) {
	// Stick Figures start at 10,10
	baseX := 10
	baseY := 10
	for i := 0; i < n; i++ {
		relX := baseX + (i * 5)
		if relX > 40 {
			relX = baseX
			baseY += 5
		}
		printStickFigure(relX, baseY, colours()[i%5])
	}

}

func clearAnswers() {
	for i := 2; i < 100; i++ {
		term.Tbprint(59, i, term.White, term.Black, "                                       ")
	}
}

func clearStickFigures() {
	for i := 10; i < 100; i++ {
		term.Tbprint(10, i, term.White, term.Black, "                                       ")
	}
}

func Solve(inputF string) (int, int) {
	defer timing.TimeTrack(time.Now())
	part1 := 0
	part2 := 0

	if input.Visualize {
		term.Init()
		printBoard()
		defer term.Close()
		time.Sleep(time.Millisecond * input.Delay * 3)
	}

	groups := inputToGroups(inputF)
	for num, g := range groups {

		if input.Visualize {
			term.Tbprint(10, 0, term.Red, term.Black, fmt.Sprintf("GROUP #%d", num))
			clearStickFigures()
			clearAnswers()
			printStickFigures(len(g))
		}
		yes, everyone := g.totalYesAndAnsweredByEveryone()
		part1 += yes
		part2 += everyone
		time.Sleep(time.Millisecond * input.Delay * 3)
	}
	return part1, part2
}
