package Day22

import (
	"strconv"
	"strings"
	"time"

	"github.com/bznein/AoC2020/pkg/input"
	"github.com/bznein/AoC2020/pkg/timing"
)

type gameStatus struct {
	p1 string
	p2 string
}

var gameToStatuses = map[int]map[gameStatus]bool{}

type playerDeck []int

var game = 0
var round = 0

func play(p1, p2 playerDeck) playerDeck {
	if len(p1) == 0 {
		return p2
	}
	if len(p2) == 0 {
		return p1
	}

	p1FirstCard := p1[0]
	p2FirstCard := p2[0]

	if p1FirstCard > p2FirstCard {
		p2 = p2[1:]
		p1 = append(p1[1:], p1FirstCard, p2FirstCard)
		return play(p1, p2)
	} else {
		p1 = p1[1:]
		p2 = append(p2[1:], p2FirstCard, p1FirstCard)
		return play(p1, p2)
	}
}

func playRound(p1, p2 playerDeck) (playerDeck, int) {
	round++
	var sb strings.Builder
	for _, v := range p1 {
		sb.WriteString(strconv.Itoa(v))
		sb.WriteString(",")
	}
	var sb2 strings.Builder
	for _, v := range p2 {
		sb2.WriteString(strconv.Itoa(v))
		sb2.WriteString(",")
	}
	if _, ok := gameToStatuses[game][gameStatus{sb.String(), sb2.String()}]; ok {
		return p1, 1
	}
	gameToStatuses[game][gameStatus{sb.String(), sb2.String()}] = true

	if len(p1) == 0 {
		return p2, 2
	}
	if len(p2) == 0 {
		return p1, 1
	}

	p1C := p1[0]
	p2C := p2[0]

	remainingP1 := len(p1) - 1
	remainingP2 := len(p2) - 1

	var winnerDeck int
	if remainingP1 >= p1C && remainingP2 >= p2C {
		p1Copy := make(playerDeck, p1C)
		p2Copy := make(playerDeck, p2C)

		copy(p1Copy, p1[1:])
		copy(p2Copy, p2[1:])
		_, winnerDeck = playGame2(p1Copy, p2Copy)

	} else if p1C > p2C {
		winnerDeck = 1
	} else {
		winnerDeck = 2
	}
	if winnerDeck == 1 {
		p2 = p2[1:]
		p1 = append(p1[1:], p1C, p2C)
		return playRound(p1, p2)
	} else {
		p1 = p1[1:]
		p2 = append(p2[1:], p2C, p1C)
		return playRound(p1, p2)
	}
}

func playGame1(p1, p2 playerDeck) playerDeck {
	return play(p1, p2)
}

func playGame2(p1, p2 playerDeck) (playerDeck, int) {
	game++
	gameToStatuses[game] = map[gameStatus]bool{}
	p, i := playRound(p1, p2)
	game--
	return p, i
}

func part2(p1, p2 playerDeck, c chan int) {
	winnerDeck, _ := playGame2(p1, p2)
	part2 := 0
	for i, v := range winnerDeck {
		part2 += (len(winnerDeck) - i) * v
	}
	c <- part2
}

func part1(p1, p2 playerDeck, c chan int) {
	winnerDeck := playGame1(p1, p2)
	part1 := 0
	for i, v := range winnerDeck {
		part1 += (len(winnerDeck) - i) * v
	}
	c <- part1
}

func Solve(inputF string) (int, int) {
	defer timing.TimeTrack(time.Now())
	s := input.InputToStringSlice(inputF)

	players := make([]playerDeck, 2)
	player := -1
	for _, ss := range s {
		if ss == "" {
			continue
		}
		if ss[0] == 'P' {
			player++
			continue
		}
		card, _ := strconv.Atoi(ss)
		players[player] = append(players[player], card)
	}

	c1 := make(chan int)
	c2 := make(chan int)
	go part1(players[0], players[1], c1)
	go part2(players[0], players[1], c2)

	part1, part2 := 0, 0
	for {
		if part1 != 0 && part2 != 0 {
			break
		}
		select {
		case p := <-c1:
			part1 = p
		case p := <-c2:
			part2 = p
		}
	}

	return part1, part2
}
