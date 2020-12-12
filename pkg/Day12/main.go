package Day12

import (
	"math"
	"strconv"
	"time"

	"github.com/bznein/AoC2020/pkg/input"
	"github.com/bznein/AoC2020/pkg/timing"
)

//TODO refactor all 2D logic into one package

type pos struct {
	i int
	j int
}

type command struct {
	move rune
	arg  int
}

func (p pos) manhattanDistance(p2 pos) int {

	return int(math.Abs(float64(p.i-p2.i)) + math.Abs(float64(p.j-p2.j)))

}

type dir rune

var posMapping = map[dir][]dir{'N': {'E', 'S', 'W'}, 'E': {'S', 'W', 'N'}, 'S': {'W', 'N', 'E'}, 'W': {'N', 'E', 'S'}}

func (d dir) rotate(towards rune, steps int) dir {

	if towards == 'L' {
		steps = 4 - steps
	}
	return posMapping[d][steps-1]
}

type ship struct {
	position       pos
	direction      dir
	wayPointoffset pos
}

func (s *ship) executeCommandP1(c command) {
	switch c.move {
	case 'N':
		s.position.i -= c.arg
	case 'S':
		s.position.i += c.arg
	case 'E':
		s.position.j += c.arg
	case 'W':
		s.position.j -= c.arg
	case 'F':
		s.executeCommandP1(command{
			move: rune(s.direction),
			arg:  c.arg,
		})
	case 'R', 'L':
		steps := int(c.arg / 90)
		s.direction = s.direction.rotate(c.move, steps)

	}
}

func (s *ship) executeCommandP2(c command) {
	switch c.move {
	case 'N':
		s.wayPointoffset.i -= c.arg
	case 'S':
		s.wayPointoffset.i += c.arg
	case 'E':
		s.wayPointoffset.j += c.arg
	case 'W':
		s.wayPointoffset.j -= c.arg
	case 'F':
		s.position.i += (s.wayPointoffset.i * c.arg)
		s.position.j += (s.wayPointoffset.j * c.arg)

	case 'R', 'L':
		steps := int(c.arg / 90)
		if c.move == 'L' {
			steps = 4 - steps
		}
		oldI := s.wayPointoffset.i
		oldJ := s.wayPointoffset.j
		switch steps {
		case 1:
			s.wayPointoffset.j = -oldI
			s.wayPointoffset.i = oldJ
		case 2: //flips signs
			s.wayPointoffset.i = -oldI
			s.wayPointoffset.j = -oldJ
		case 3:
			s.wayPointoffset.j = oldI
			s.wayPointoffset.i = -oldJ
		case 4: //Nothing to do here
		}

	}
}

func Solve(inputF string) (int, int) {
	defer timing.TimeTrack(time.Now())
	part1, part2 := -1, -1

	s := input.InputToStringSlice(inputF)

	commands := make([]command, len(s))

	for i, cmd := range s {
		move := cmd[0]
		arg, _ := strconv.Atoi(cmd[1:])
		commands[i] = command{
			move: rune(move),
			arg:  arg,
		}
	}

	curShip := ship{position: pos{i: 0, j: 0}, direction: 'E'}
	p2Ship := ship{position: pos{i: 0, j: 0}, direction: 'E', wayPointoffset: pos{i: -1, j: 10}}

	for _, c := range commands {
		curShip.executeCommandP1(c)
		p2Ship.executeCommandP2(c)
	}

	part1 = pos{}.manhattanDistance(curShip.position)
	part2 = pos{}.manhattanDistance(p2Ship.position)

	return part1, part2
}
