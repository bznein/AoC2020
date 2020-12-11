package Day04

import (
	"fmt"
	"time"

	"github.com/bznein/AoC2020/pkg/input"
	"github.com/bznein/AoC2020/pkg/passport"
	"github.com/bznein/AoC2020/pkg/term"
	"github.com/bznein/AoC2020/pkg/timing"
)

func Solve(inputF string) (int, int) {
	defer timing.TimeTrack(time.Now())
	part1 := 0
	part2 := 0

	if input.Visualize {
		term.Init()
		defer term.Close()
	}
	passports := passport.StringToPassportSlice(inputF)

	for _, p := range passports {

		if input.Visualize {
			p.Print()
			term.Tbprint(8, 26, term.White, term.Black, "-----AIRPORT SECURITY-----")
			term.Separator(8, 26, 6)
			term.Separator(33, 26, 6)
			term.Tbprint(8, 32, term.White, term.Black, "--------------------------")
		}

		if p.HasAllRequiredFields() {
			part1++
			if input.Visualize {
				term.BicolorString(10, 28, "REQUIRED FIELDS = YES", term.White, term.Green, term.Black, 18, 19, 20)
			}
			if p.IsValid() {
				part2++
				if input.Visualize {
					term.BicolorString(10, 30, "VALID FIELDS = YES", term.White, term.Green, term.Black, 15, 16, 17)
				}
			} else if input.Visualize {
				term.BicolorString(10, 30, "VALID FIELDS = NO ", term.White, term.Red, term.Black, 15, 16)
			}
		} else if input.Visualize {
			term.BicolorString(10, 28, "REQUIRED FIELDS = NO ", term.White, term.Red, term.Black, 18, 19)
			term.BicolorString(10, 30, "VALID FIELDS = N/A", term.White, term.Red, term.Black, 15, 16, 17)
		}
		if input.Visualize {
			term.Tbprint(8, 37, term.White, term.Black, fmt.Sprintf("# PASSPORTS WITH REQUIRED FIELDS: %d", part1))
			term.Tbprint(8, 39, term.White, term.Black, fmt.Sprintf("# PASSPORTS WITH VALID FIELDS: %d", part2))
		}
		time.Sleep(time.Millisecond * input.Delay)
	}

	return part1, part2
}
