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

func day(tilesIn map[twod.Position]bool) (map[twod.Position]bool, int) {
	tilesOut := map[twod.Position]bool{}
	blackTiles := 0
	for i := -tileSize; i <= tileSize; i++ {
		for j := -tileSize; j <= tileSize; j++ {
			k := twod.Position{I: i, J: j}
			v := tilesIn[k]
			adjacentBlack := 0
			if val, ok := tilesIn[k.East()]; ok && val {
				adjacentBlack++
			}
			if val, ok := tilesIn[k.West()]; ok && val {
				adjacentBlack++
			}
			if val, ok := tilesIn[k.SouthWest()]; ok && val {
				adjacentBlack++
			}
			if val, ok := tilesIn[k.SouthEast()]; ok && val {
				adjacentBlack++
			}
			if val, ok := tilesIn[k.NorthEast()]; ok && val {
				adjacentBlack++
			}
			if val, ok := tilesIn[k.NorthWest()]; ok && val {
				adjacentBlack++
			}

			if v {
				if adjacentBlack == 0 || adjacentBlack > 2 {
					tilesOut[k] = false
				} else {
					tilesOut[k] = true
					blackTiles++
				}
			} else {
				if adjacentBlack == 2 {
					tilesOut[k] = true
					blackTiles++
				} else {
					tilesOut[k] = false
				}
			}
		}

	}
	return tilesOut, blackTiles
}

func Solve(inputF string) (int, int) {
	defer timing.TimeTrack(time.Now())
	part1, part2 := 0, 0
	s := input.InputToStringSlice(inputF)

	tiles := map[twod.Position]bool{}

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
		tiles[pos] = !tiles[pos]
		if !tiles[pos] {
			part1--
		} else {
			part1++
		}

	}

	for i := 1; i <= days; i++ {
		tiles, part2 = day(tiles)
	}
	return part1, part2
}
