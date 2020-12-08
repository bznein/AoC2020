package term

import (
	"strings"

	"github.com/mattn/go-runewidth"
	"github.com/nsf/termbox-go"
)

const (
	newLine = "                                                                                                                                                                                                                                                                                                                                                                                                 "

	Black  = termbox.ColorBlack
	Green  = termbox.ColorGreen
	Red    = termbox.ColorRed
	White  = termbox.ColorWhite
	Gray   = termbox.ColorDarkGray
	Yellow = termbox.ColorYellow
	Cyan   = termbox.ColorCyan
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

func ClearArea(x, y int, w, h int) {
	var sb strings.Builder
	for i := 0; i < w; i++ {
		sb.WriteString(" ")
	}
	for i := 0; i < h; i++ {
		Tbprint(x, y+i, White, Black, sb.String())
	}
}

func Init() {
	termbox.Init()
	termbox.Clear(White, Black)
	termbox.Flush()
}

func Close() {
	termbox.Close()
}

func HSeparator(x int, y int, l int, c rune) {
	for i := 1; i <= l; i++ {
		Tbprint(x, y, White, Black, string(c))
		x += runewidth.RuneWidth(c)
	}
}

func Separator(x int, y int, l int) {
	for i := 1; i <= l; i++ {
		Tbprint(x, y+i, White, Black, "|")
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

// TODO nicer interface to have this multicolour
func StringSlice(x, y int, fg, bg termbox.Attribute, slice []string) {
	for i, s := range slice {
		Tbprint(x, y+i, fg, bg, s)
	}
}
