package Day23

import (
	"container/ring"
	"time"

	"github.com/bznein/AoC2020/pkg/input"
	"github.com/bznein/AoC2020/pkg/timing"
)

func newDest(dest int, elems int) int {
	return (dest-2+elems)%elems + 1
}

func playGame(inputF string, elems int, moves int) int {
	cups := ring.New(elems)
	positions := make([]*ring.Ring, elems+1)
	for i := 1; i <= elems; i++ {
		if i <= len(inputF) {
			cups.Value = int(inputF[i-1] - '0')
		} else {
			cups.Value = i
		}
		positions[cups.Value.(int)] = cups
		cups = cups.Next()
	}

	for i := 0; i < moves; i++ {
		pickedUp := cups.Unlink(3)
		val := cups.Value.(int)
		dest := newDest(val, elems)

		unavail := map[int]bool{}
		for i := 0; i < 3; i++ {
			unavail[pickedUp.Value.(int)] = true
			pickedUp = pickedUp.Next()
		}

		for unavail[dest] {
			dest = newDest(elems+dest, elems)
		}

		positions[dest].Link(pickedUp)
		cups = cups.Next()
	}

	p1 := positions[1]
	if elems > len(inputF) {
		return p1.Next().Value.(int) * p1.Move(2).Value.(int)
	}
	res := 0
	p1.Unlink(len(inputF) - 1).Do(func(p interface{}) {
		res = res*10 + p.(int)
	})
	return res
}

func Solve(inputF string) (int, int) {
	defer timing.TimeTrack(time.Now())
	part1, part2 := 0, -1
	n := input.InputToStringSlice(inputF)

	part1 = playGame(n[0], len(n[0]), 100)
	part2 = playGame(n[0], 1000000, 10000000)
	return part1, part2
}
