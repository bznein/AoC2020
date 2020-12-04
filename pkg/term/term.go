package term

import (
	"github.com/mattn/go-runewidth"
	"github.com/nsf/termbox-go"
)

const (
	newLine = "                                                                                                                                                                                                                                                                                                                                                                                                 "

	Black = termbox.ColorBlack
	Green = termbox.ColorGreen
	Red   = termbox.ColorRed
	White = termbox.ColorWhite
)

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

func Init() {
	termbox.Init()
	termbox.Clear(White, Black)
	termbox.Flush()
}

func Close() {
	termbox.Close()
}

func Separator(x int) {
	for i := 1; i <= 5; i++ {
		Tbprint(x, i, White, Black, "|")
	}
}

func BicolorString(x, y int, s string, baseColor termbox.Attribute, alternateColor termbox.Attribute, bg termbox.Attribute, indices ...int) {
	for i, c := range s {
		correctC := baseColor
		for _, idx := range indices {
			if idx == i {
				correctC = alternateColor
				break
			}
		}
		termbox.SetCell(x, y, c, correctC, bg)
		x += runewidth.RuneWidth(c)
	}
	termbox.Flush()
}
