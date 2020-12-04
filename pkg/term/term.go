package term

import (
	"github.com/mattn/go-runewidth"
	"github.com/nsf/termbox-go"
)

const newLine = "                                                                                                                                                                                                                         "

func Tbprint(x, y int, fg, bg termbox.Attribute, msg string) {
	for _, c := range msg {
		termbox.SetCell(x, y, c, fg, bg)
		x += runewidth.RuneWidth(c)
	}
	termbox.Flush()
}

func ClearLine(y int) {
	Tbprint(0, y, termbox.ColorBlack, termbox.ColorBlack, newLine)
}
