package Day24

import (
	"time"

	"github.com/bznein/AoC2020/pkg/input"
	"github.com/bznein/AoC2020/pkg/timing"
	"github.com/bznein/AoC2020/pkg/twod"
)

const (
	days     = 100
	tileSize = 66
)

func day(tilesIn map[twod.Position]bool) map[twod.Position]bool {
	tilesOut := map[twod.Position]bool{}

	for i := -tileSize; i <= tileSize; i++ {
		for j := -tileSize; j <= tileSize; j++ {
			k := twod.Position{I: i, J: j}
			v := tilesIn[k]
			adjacentBlack := 0
			if val, ok := tilesIn[k.East()]; ok && !val {
				adjacentBlack++
			}
			if val, ok := tilesIn[k.West()]; ok && !val {
				adjacentBlack++
			}
			if val, ok := tilesIn[k.SouthWest()]; ok && !val {
				adjacentBlack++
			}
			if val, ok := tilesIn[k.SouthEast()]; ok && !val {
				adjacentBlack++
			}
			if val, ok := tilesIn[k.NorthEast()]; ok && !val {
				adjacentBlack++
			}
			if val, ok := tilesIn[k.NorthWest()]; ok && !val {
				adjacentBlack++
			}

			if !v {
				if adjacentBlack == 0 || adjacentBlack > 2 {
					tilesOut[k] = true
				} else {
					tilesOut[k] = false
				}
			} else {
				if adjacentBlack == 2 {
					tilesOut[k] = false
				} else {
					tilesOut[k] = true
				}
			}
		}

	}
	return tilesOut
}

func blackTiles(tiles map[twod.Position]bool) int {
	tot := 0
	for _, v := range tiles {
		if !v {
			tot++
		}
	}
	return tot
}

func Solve(inputF string) (int, int) {
	defer timing.TimeTrack(time.Now())
	part1, part2 := 0, -1
	s := input.InputToStringSlice(inputF)

	//TRUE: WHITE SIDE UP
	tiles := map[twod.Position]bool{}

	for i := -tileSize; i <= tileSize; i++ {
		for j := -tileSize; j <= tileSize; j++ {
			tiles[twod.Position{I: i, J: j}] = true
		}
	}

	for _, tile := range s {
		pos := twod.Position{}
		for i := 0; i < len(tile); i++ {
			switch tile[i] {
			case 'e':
				pos.MoveEast()
			case 'w':
				pos.MoveWest()
			case 'n':
				i++
				if tile[i] == 'w' {
					pos.MoveNorthWest()
				} else {
					pos.MoveNorthEast()
				}
			case 's':
				i++
				if tile[i] == 'w' {
					pos.MoveSouthWest()
				} else {
					pos.MoveSouhtEast()
				}
			}
		}
		if _, ok := tiles[pos]; ok {
			tiles[pos] = !tiles[pos]
		}
	}

	part1 = blackTiles(tiles)

	for i := 1; i <= days; i++ {
		tiles = day(tiles)
	}
	part2 = blackTiles(tiles)
	return part1, part2
}
