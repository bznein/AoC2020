package Day12

import (
	"strconv"
	"time"

	"github.com/bznein/AoC2020/pkg/input"
	"github.com/bznein/AoC2020/pkg/timing"
	"github.com/bznein/AoC2020/pkg/twod"
)

type command struct {
	move rune
	arg  int
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
	position       twod.Position
	direction      dir
	wayPointoffset twod.Position
}

func (s *ship) moveBy(i, j int) {
	s.position.MoveBy(i, j)
}

func (s *ship) moveWaypointBy(i, j int) {
	s.wayPointoffset.MoveBy(i, j)
}

func (s *ship) executeCommandP1(c command) {
	switch c.move {
	case 'N':
		s.moveBy(-c.arg, 0)
	case 'S':
		s.moveBy(c.arg, 0)
	case 'E':
		s.moveBy(0, c.arg)
	case 'W':
		s.moveBy(0, -c.arg)

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
		s.moveWaypointBy(-c.arg, 0)
	case 'S':
		s.moveWaypointBy(c.arg, 0)
	case 'E':
		s.moveWaypointBy(0, c.arg)
	case 'W':
		s.moveWaypointBy(0, -c.arg)

	case 'F':
		s.moveBy(s.wayPointoffset.I*c.arg, s.wayPointoffset.J*c.arg)
	case 'R', 'L':
		s.wayPointoffset.SnapRotate(c.move == 'R', c.arg)
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

	curShip := ship{position: twod.Position{I: 0, J: 0}, direction: 'E'}
	p2Ship := ship{position: twod.Position{I: 0, J: 0}, direction: 'E', wayPointoffset: twod.Position{I: -1, J: 10}}

	for _, c := range commands {
		curShip.executeCommandP1(c)
		p2Ship.executeCommandP2(c)
	}

	part1 = twod.Position{}.ManhattanDistance(curShip.position)
	part2 = twod.Position{}.ManhattanDistance(p2Ship.position)

	return part1, part2
}
