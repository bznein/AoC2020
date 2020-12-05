package term

import (
	"fmt"
	"strings"
)

type seat struct {
	x int
	y int
}

const (
	wingWidth  = 70
	tailHeight = 15
	tailWidth  = 30
	frontWidth = 30
)

type Aeroplane struct {
	X           int
	Y           int
	Rows        int
	SeatsPerRow int
}

func (a Aeroplane) paintColumn(i int) {
	var sb strings.Builder
	for r := 0; r < a.Rows; r++ {
		sb.WriteString("O ")
	}
	Tbprint(a.X, a.Y+i, Green, Black, sb.String())
}

func (a Aeroplane) TakeSeat(row, column int) {
	if column >= a.SeatsPerRow/2 {
		column++
	}
	Tbprint(a.X+(row*2), a.Y+2+column, Red, Black, "X")
}

func (a Aeroplane) SeatUnavailable(row, column int) {
	if column >= a.SeatsPerRow/2 {
		column++
	}
	Tbprint(a.X+(row*2), a.Y+2+column, Gray, Black, "-")
}

func (a Aeroplane) paintRowsNumbers() {
	var sb1, sb2, sb3 strings.Builder
	for r := 0; r < a.Rows; r++ {

		v := r
		digit := v % 10
		v /= 10
		tens := v % 10
		v /= 10
		hundreds := v % 10
		sb3.WriteString(fmt.Sprintf("%d ", digit))
		if tens != 0 || hundreds != 0 {
			sb2.WriteString(fmt.Sprintf("%d ", tens))
		} else {
			sb2.WriteString("  ")
		}
		if hundreds != 0 {
			sb1.WriteString(fmt.Sprintf("%d ", hundreds))
		} else {
			sb1.WriteString("  ")
		}
	}
	Tbprint(a.X, a.Y+13, White, Black, sb1.String())
	Tbprint(a.X, a.Y+14, White, Black, sb2.String())
	Tbprint(a.X, a.Y+15, White, Black, sb3.String())
}

func (a Aeroplane) paintColumnsNumbers() {
	offset := 2
	for i := 0; i < a.SeatsPerRow; i++ {
		if i == a.SeatsPerRow/2 {
			offset = 3
		}
		Tbprint(a.X+a.Rows*2+1, a.Y+i+offset, White, Black, fmt.Sprintf("%d", i))
	}
}

func (a Aeroplane) PaintInit() {

	HSeparator(a.X, a.Y, a.Rows*2, '‾')
	HSeparator(a.X, a.Y+12, a.Rows*2, '_')
	a.paintWings()
	a.paintRear()
	a.paintFront()
	offset := 2
	for i := 0; i < a.SeatsPerRow; i++ {
		a.paintColumn(i + offset)
		if i == a.SeatsPerRow/2-1 {
			offset = 3
		}
	}
}

func (a Aeroplane) paintWings() {
	for i := 0; i < wingWidth; i++ {
		var sb strings.Builder
		for j := 0; j < i; j++ {
			sb.WriteString(" ")
		}
		sb.WriteString("/")
		for j := i; j < wingWidth-1; j++ {
			sb.WriteString(" ")
		}
		sb.WriteString("|")
		Tbprint(a.X+60, a.Y-1-i, White, Black, sb.String())
	}
	for i := 0; i < wingWidth; i++ {
		var sb strings.Builder
		for j := 0; j < i; j++ {
			sb.WriteString(" ")
		}
		sb.WriteString("\\")
		for j := i; j < wingWidth-1; j++ {
			sb.WriteString(" ")
		}
		sb.WriteString("|")
		Tbprint(a.X+60, a.Y+13+i, White, Black, sb.String())
	}
}

func (a Aeroplane) paintRear() {
	Separator(a.X+256, a.Y-2, 14)

	// Upper side
	for i := 0; i < tailHeight; i++ {
		var sb strings.Builder
		for j := 0; j < i; j++ {
			sb.WriteString(" ")
		}
		sb.WriteString("/")
		for j := 0; j <= tailWidth; j++ {
			if i < tailHeight-1 {
				sb.WriteString(" ")
			} else {
				sb.WriteString("‾")
			}
		}
		sb.WriteString("/")
		Tbprint(a.X+256, a.Y-1-i, White, Black, sb.String())
	}
	for i := 0; i < 7; i++ {
		var sb strings.Builder
		for j := 0; j < tailWidth-i; j++ {
			sb.WriteString(" ")
		}

		sb.WriteString("/")
		Tbprint(a.X+257, a.Y+i, White, Black, sb.String())
	}

	// Lower side
	for i := 0; i < tailHeight; i++ {
		var sb strings.Builder
		for j := 0; j < i; j++ {
			sb.WriteString(" ")
		}
		sb.WriteString("\\")
		for j := 0; j < tailWidth; j++ {
			if i < tailHeight-1 {
				sb.WriteString(" ")
			} else {
				sb.WriteString("_")
			}
		}
		sb.WriteString("\\")
		Tbprint(a.X+256, a.Y+13+i, White, Black, sb.String())
	}
	for i := 0; i < 7; i++ {
		var sb strings.Builder
		for j := 0; j < tailWidth-i; j++ {
			sb.WriteString(" ")
		}

		sb.WriteString("\\")
		Tbprint(a.X+257, a.Y+13-i, White, Black, sb.String())
	}

}

func (a Aeroplane) paintFront() {
	// Upper part
	for i := 0; i < 7; i++ {
		var sb strings.Builder
		for j := 0; j < i; j++ {
			sb.WriteString(" ")
		}
		sb.WriteString("/")
		if i == 6 {
			for j := 0; j < frontWidth; j++ {
				sb.WriteString("‾")
			}
		}
		Tbprint(a.X-6-frontWidth, a.Y+6-i, White, Black, sb.String())
	}

	// Lower part
	for i := 0; i < 6; i++ {
		var sb strings.Builder
		for j := 0; j < i; j++ {
			sb.WriteString(" ")
		}
		sb.WriteString("\\")
		if i == 5 {
			for j := 0; j < frontWidth; j++ {
				sb.WriteString("_")
			}
		}
		Tbprint(a.X-6-frontWidth, a.Y+7+i, White, Black, sb.String())
	}
	Separator(a.X-5, a.Y-1, 13)

}
