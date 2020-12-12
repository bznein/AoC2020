package twod

import (
	"fmt"
	"math"
)

type Position struct {
	I int
	J int
}

func (p Position) ManhattanDistance(p2 Position) int {

	fmt.Printf("p: %v\n", p)
	fmt.Printf("p2: %v\n", p2)
	return int(math.Abs(float64(p.I-p2.I)) + math.Abs(float64(p.J-p2.J)))

}

func (p *Position) MoveBy(i int, j int) {
	p.I += i
	p.J += j
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
