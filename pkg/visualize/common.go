package visualize

import (
	"github.com/bznein/AoC2020/pkg/input"
	term "github.com/bznein/AoC2020/pkg/term"
)

func Init() {
	if input.Visualize {
		term.Init()
	}
}

func Close() {

	if input.Visualize {
		term.Close()
	}
}
