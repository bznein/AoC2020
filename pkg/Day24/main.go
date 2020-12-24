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

func day(tilesIn [][]bool) ([][]bool, int) {
	tilesOut := make([][]bool, tileSize*2+2)
	for i := range tilesOut {
		tilesOut[i] = make([]bool, tileSize*2+2)
	}
	blackTiles := 0
	for i := -tileSize + 1; i <= tileSize; i++ {
		for j := -tileSize + 1; j <= tileSize; j++ {
			k := twod.Position{I: i + tileSize, J: j + tileSize}
			adjacentBlack := 0
			east := k.East()
			west := k.West()
			southWest := k.SouthWest()
			southEast := k.SouthEast()
			northEast := k.NorthEast()
			northWest := k.NorthWest()
			if tilesIn[east.I][east.J] {
				adjacentBlack++
			}
			if tilesIn[west.I][west.J] {
				adjacentBlack++
			}
			if tilesIn[southWest.I][southWest.J] {
				adjacentBlack++
			}
			if tilesIn[southEast.I][southEast.J] {
				adjacentBlack++
			}
			if tilesIn[northEast.I][northEast.J] {
				adjacentBlack++
			}
			if tilesIn[northWest.I][northWest.J] {
				adjacentBlack++
			}

			if tilesIn[k.I][k.J] {
				if adjacentBlack == 0 || adjacentBlack > 2 {
					tilesOut[k.I][k.J] = false
				} else {
					tilesOut[k.I][k.J] = true
					blackTiles++
				}
			} else {
				if adjacentBlack == 2 {
					tilesOut[k.I][k.J] = true
					blackTiles++
				} else {
					tilesOut[k.I][k.J] = false
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

	tiles := make([][]bool, tileSize*2+2)
	for i := range tiles {
		tiles[i] = make([]bool, tileSize*2+2)
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
		pos.MoveBy(tileSize, tileSize)
		tiles[pos.I][pos.J] = !tiles[pos.I][pos.J]
		if tiles[pos.I][pos.J] {
			part1++
		} else {
			part1--
		}

	}

	for i := 1; i <= days; i++ {
		tiles, part2 = day(tiles)
	}
	return part1, part2
}
