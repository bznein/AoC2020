package twod

import (
	"fmt"
	"github.com/agrison/go-commons-lang/stringUtils"
	"math"
	"strings"
)

type Position struct {
	I int
	J int
}

type Slope struct {
	I int
	J int
}

type Grid []string

func (g Grid) AreValidIndices(i, j int) bool {
	return i >= 0 && i < len(g) && j >= 0 && j < len(g[i])
}

func (g Grid) UnsafeEntry(i, j int) rune {
	return rune(g[i][j])
}

func (g Grid) EntryAt(i, j int) (rune, error) {
	if !g.AreValidIndices(i, j) {
		return '0', fmt.Errorf("Requested invalid indices")
	}
	return rune(g[i][j]), nil
}

func (g Grid) Print() {
	for _, row := range g {
		for _, val := range row {
			fmt.Printf("%c", val)
		}
		fmt.Println()
	}
}

func (g Grid) FirstRow() string {
	return g.GetRow(0)
}

func (g Grid) FirstCol() string {
	return g.GetColumn(0)
}

func (g Grid) LastRow() string {
	return g.GetRow(len(g) - 1)
}

func (g Grid) LastCol() string {
	return g.GetColumn(len(g[0]) - 1)
}

func (g Grid) GetRow(i int) string {
	if i < 0 || i >= len(g) {
		return ""
	}
	return g[i]
}

func (g Grid) GetColumn(j int) string {
	if j < 0 || j >= len(g[0]) {
		return ""
	}
	var sb strings.Builder
	for i := range g {
		sb.WriteRune(rune(g[i][j]))
	}
	return sb.String()
}

func (g Grid) IsEntry(i int, j int, c rune) bool {
	v, err := g.EntryAt(i, j)
	return err == nil && v == c
}

func (p *Position) MoveBySlope(s Slope) {
	p.MoveBy(s.I, s.J)
}

func (p *Position) MoveBySlopeWithWrapping(s Slope, iWrap int, jWrap int) {
	p.MoveWithWrapping(s.I, s.J, iWrap, jWrap)
}

func (p Position) ManhattanDistance(p2 Position) int {
	return int(math.Abs(float64(p.I-p2.I)) + math.Abs(float64(p.J-p2.J)))

}

func (p *Position) MoveBy(i int, j int) {
	p.I += i
	p.J += j
}

func (p *Position) MoveWithWrapping(i int, j int, iWrap int, jWrap int) {
	p.MoveBy(i, j)
	if iWrap != -1 {
		p.I %= iWrap
	}
	if jWrap != -1 {
		p.J %= jWrap
	}
}

func (p *Position) SnapRotate(clockwise bool, degrees int) {
	if degrees != 360 && degrees != 270 && degrees != 180 && degrees != 90 {
		panic("SnapRotate only supports degrees of 90,180,270,360")
	}
	steps := int(degrees / 90)
	if !clockwise {
		steps = 4 - steps
	}
	oldI := p.I
	oldJ := p.J
	switch steps {
	case 1:
		p.J = -oldI
		p.I = oldJ
	case 2: //flips signs
		p.I = -oldI
		p.J = -oldJ
	case 3:
		p.J = oldI
		p.I = -oldJ
	case 4: //Nothing to do here
	}
}

func (g Grid) Rotate(steps int) Grid {
	if steps == 0 || steps == 4 {
		return g
	}
	res := make(Grid, len(g))
	if steps == 1 {
		for i := 0; i < len(g); i++ {
			var sb strings.Builder
			for j := 0; j < len(g[0]); j++ {
				sb.WriteRune(rune(g[len(g)-j-1][i]))
			}
			res[i] = sb.String()
		}
	}
	if steps == 2 {
		for i, s := range g {
			res[len(g)-1-i] = stringUtils.Reverse(s)
		}
	}
	if steps == 3 {
		res = g.Rotate(2)
		res = res.Rotate(1)
	}
	return res
}

func (g Grid) FlipX() Grid {
	res := make(Grid, len(g))
	for i, s := range g {
		res[i] = stringUtils.Reverse(s)
	}
	return res
}

func (g Grid) FlipY() Grid {
	res := make(Grid, len(g))
	for i, s := range g {
		res[len(g)-1-i] = s
	}
	return res
}
