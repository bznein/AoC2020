package console

import (
	"fmt"

	"github.com/bznein/AoC2020/pkg/term"
)

func (c Executor) printSkeleton() {
	term.Tbprint(0, 0, term.White, term.Black, "                         ______                     ")
	term.Tbprint(0, 1, term.White, term.Black, " _________        .---\"\"\"      \"\"\"---.              ")
	term.Tbprint(0, 2, term.White, term.Black, ":______.-':      :  .--------------.  :             ")
	term.Tbprint(0, 3, term.White, term.Black, "| ______  |      | :                : |             ")
	term.Tbprint(0, 4, term.White, term.Black, "|         |      | :                : |             ")
	term.Tbprint(0, 5, term.White, term.Black, "|:______B:|      | |                | |             ")
	term.Tbprint(0, 6, term.White, term.Black, "|:______B:|      | |                | |             ")
	term.Tbprint(0, 7, term.White, term.Black, "|:______B:|      | |                | |             ")
	term.Tbprint(0, 8, term.White, term.Black, "|         |      | |                | |             ")
	term.Tbprint(0, 9, term.White, term.Black, "|:_____:  |      | |                | |             ")
	term.Tbprint(0, 10, term.White, term.Black, "|    ==   |      | :                : |             ")
	term.Tbprint(0, 11, term.White, term.Black, "|       O |      :  '--------------'  :             ")
	term.Tbprint(0, 12, term.White, term.Black, "|       o |      :'---...______...---'              ")
	term.Tbprint(0, 13, term.White, term.Black, "|       o |-._.-i___/'             \\._              ")
	term.Tbprint(0, 14, term.White, term.Black, "|'-.____o_|   '-.   '-...______...-'  `-._          ")
	term.Tbprint(0, 15, term.White, term.Black, ":_________:      `.____________________   `-.___.-. ")
	term.Tbprint(0, 16, term.White, term.Black, "                 .'.eeeeeeeeeeeeeeeeee.'.      :___:")
	term.Tbprint(0, 17, term.White, term.Black, "               .'.eeeeeeeeeeeeeeeeeeeeee.'.         ")
	term.Tbprint(0, 18, term.White, term.Black, "              :____________________________:")
}

func (c Executor) Print() {
	c.printSkeleton()
	term.Tbprint(22, 4, term.White, term.Black, "INSTRUCTION")
	ins := c.getInstruction(c.ip)
	term.Tbprint(22, 5, term.White, term.Black, fmt.Sprintf("%s %+d", ins.command, ins.argument))
	term.Tbprint(22, 7, term.White, term.Black, fmt.Sprintf("IP: %d", c.ip))
	term.Tbprint(22, 9, term.White, term.Black, fmt.Sprintf("ACC: %d", c.accumulator))
}

func (c Executor) InfiniteLoopDetected() {
	term.ClearArea(22, 4, 12, 6)
	term.Tbprint(22, 4, term.Red, term.Black, "INFINITE")
	term.Tbprint(22, 5, term.Red, term.Black, "LOOP")
	term.Tbprint(22, 6, term.Red, term.Black, "DETECTED")
	term.Tbprint(22, 7, term.Red, term.Black, ":(")
}

func (c Executor) TerminedSuccesfully() {
	term.Tbprint(22, 4, term.Green, term.Black, "PROGRAM")
	term.Tbprint(22, 5, term.Green, term.Black, "EXITED")
	term.Tbprint(22, 6, term.Green, term.Black, "SUCCESFULLY")
	term.Tbprint(22, 7, term.Green, term.Black, ":)")
}
