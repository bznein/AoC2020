package Day20

import (
	"fmt"
	"github.com/agrison/go-commons-lang/stringUtils"
	"github.com/bznein/AoC2020/pkg/input"
	"github.com/bznein/AoC2020/pkg/timing"
	"github.com/bznein/AoC2020/pkg/twod"
	"math"
	"strings"
	"time"
)

var width int
var height int

func tiles(s []string) map[int]twod.Grid {
	m := map[int]twod.Grid{}
	var tempGrid twod.Grid
	key := -1
	for _, ss := range s {
		if tempGrid == nil {
			fmt.Sscanf(ss, "Tile %d", &key)
			tempGrid = twod.Grid{}
			continue
		}
		if ss == "" {
			m[key] = tempGrid
			tempGrid = nil
			continue
		}
		tempGrid = append(tempGrid, ss)
	}
	m[key] = tempGrid
	return m
}

func findMatchingStringTile(tileFrom *int, tiles map[int]twod.Grid, s string) int {
	// Loop through all the mak, skip k==tileFrom
	// For each, get the 4 corners, and check if the string, or the string flipped
	// match

	sReverse := stringUtils.Reverse(s)
	for k, v := range tiles {
		if tileFrom != nil && k == *tileFrom {
			continue
		}
		s1 := v.GetRow(0)
		s2 := v.GetRow(len(v) - 1)
		s3 := v.GetColumn(0)
		s4 := v.GetColumn(len(s1) - 1)

		if s == s1 || sReverse == s1 {
			return k
		}
		if s == s2 || sReverse == s2 {
			return k
		}
		if s == s3 || sReverse == s3 {
			return k
		}
		if s == s4 || sReverse == s4 {
			return k
		}
	}
	return -1
}

type finalImage map[twod.Position]twod.Grid

func (img finalImage) print() {
	for i := 0; i < height; i++ {
		for k := 0; k < len(img[twod.Position{}]); k++ {
			for j := 0; j < width; j++ {
				fmt.Printf(" %s", img[twod.Position{i, j}][k])
			}
			fmt.Println()
		}
		fmt.Println()
	}
	fmt.Println()
	fmt.Println()

}

func (img finalImage) toGrid() twod.Grid {
	g := twod.Grid{}
	for i := 0; i < height; i++ {
		for k := 1; k < len(img[twod.Position{}])-1; k++ {
			var sb strings.Builder
			for j := 0; j < width; j++ {
				s := img[twod.Position{i, j}][k]
				sb.WriteString(s[1 : len(s)-1])
			}
			g = append(g, sb.String())
		}
	}
	return g
}

func fillImage(image finalImage, tiles map[int]twod.Grid) twod.Grid {
	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			if i == 0 && j == 0 {

				continue
			}
			var v twod.Grid
			var s string
			// Find this one
			if j != 0 {
				v = image[twod.Position{i, j - 1}]
				s = v.LastCol()

			} else {
				v = image[twod.Position{i - 1, j}]
				s = v.LastRow()
			}
			k := findMatchingStringTile(nil, tiles, s)
			var img twod.Grid
			for i := 0; i <= 3; i++ {
				v := tiles[k]
				t1 := v.Rotate(i)
				t2 := v.Rotate(i).FlipX()
				t3 := v.Rotate(i).FlipY()
				if j != 0 {
					if t1.FirstCol() == s {
						img = t1
						break
					}
					if t2.FirstCol() == s {
						img = t2
						break
					}
					if t3.FirstCol() == s {
						img = t3
						break
					}
				} else {
					if t1.FirstRow() == s {
						img = t1
						break
					}
					if t2.FirstRow() == s {
						img = t2
						break
					}
					if t3.FirstRow() == s {
						img = t3
						break
					}
				}
			}
			delete(tiles, k)
			image[twod.Position{i, j}] = img
		}
	}
	return image.toGrid()
}

func countSeaMonsters(img twod.Grid) int {
	monstersEntries := map[twod.Position]bool{}
	for i, row := range img {
		for j, val := range row {
			if val == '#' {
				pos := twod.Position{i, j}
				if _, ok := monstersEntries[pos]; ok {
					continue
				}
				// I AM NOT PROUD
				if img.IsEntry(i+1, j+1, '#') {

					if img.IsEntry(i+1, j+4, '#') {

						if img.IsEntry(i, j+5, '#') {

							if img.IsEntry(i, j+6, '#') {

								if img.IsEntry(i+1, j+7, '#') {

									if img.IsEntry(i+1, j+10, '#') {

										if img.IsEntry(i, j+11, '#') {

											if img.IsEntry(i, j+12, '#') {

												if img.IsEntry(i+1, j+13, '#') {

													if img.IsEntry(i+1, j+16, '#') {

														if img.IsEntry(i, j+17, '#') {

															if img.IsEntry(i, j+18, '#') {

																if img.IsEntry(i, j+19, '#') {

																	if img.IsEntry(i-1, j+18, '#') {
																		// LOL THIS IS A MOSTER
																		// NOT LET'S MARK ALL THIS ENTRIES
																		monstersEntries[twod.Position{i, j}] = true
																		monstersEntries[twod.Position{i + 1, j + 1}] = true
																		monstersEntries[twod.Position{i + 1, j + 4}] = true
																		monstersEntries[twod.Position{i, j + 5}] = true
																		monstersEntries[twod.Position{i, j + 6}] = true
																		monstersEntries[twod.Position{i + 1, j + 7}] = true
																		monstersEntries[twod.Position{i + 1, j + 10}] = true
																		monstersEntries[twod.Position{i, j + 11}] = true
																		monstersEntries[twod.Position{i, j + 12}] = true
																		monstersEntries[twod.Position{i + 1, j + 13}] = true
																		monstersEntries[twod.Position{i + 1, j + 16}] = true
																		monstersEntries[twod.Position{i, j + 17}] = true
																		monstersEntries[twod.Position{i, j + 18}] = true
																		monstersEntries[twod.Position{i, j + 19}] = true
																		monstersEntries[twod.Position{i - 1, j + 18}] = true
																	}
																}
															}
														}
													}
												}

											}
										}
									}
								}
							}
						}
					}

				}
			}
		}
	}
	tot := 0
	for i, row := range img {
		for j, val := range row {
			if val == '#' {
				pos := twod.Position{i, j}
				if _, ok := monstersEntries[pos]; !ok {
					tot++
				}
			}
		}
	}
	if len(monstersEntries) == 0 {
		return -1
	}
	return tot
}

func Solve(inputF string) (int, int) {
	defer timing.TimeTrack(time.Now())
	part1, part2 := 1, -1
	s := input.InputToStringSlice(inputF)

	tiles := tiles(s)

	width = int(math.Sqrt(float64(len(tiles))))
	height = width
	image := finalImage{}
	del := -1
	for k, v := range tiles {
		tot := 0
		if findMatchingStringTile(&k, tiles, v.FirstCol()) != -1 {
			tot++
		}
		if findMatchingStringTile(&k, tiles, v.LastCol()) != -1 {
			tot++
		}
		if findMatchingStringTile(&k, tiles, v.FirstRow()) != -1 {
			tot++
		}
		if findMatchingStringTile(&k, tiles, v.LastRow()) != -1 {
			tot++
		}
		if tot == 2 {
			part1 *= k
			if _, ok := image[twod.Position{}]; !ok {
				//Arbitratily put this in upper left, the rest will follow
				// Try all combination and see in which case first row and first col do no have matches
				for i := 0; i <= 3; i++ {
					t1 := v.Rotate(i)
					t2 := v.Rotate(i).FlipX()
					t3 := v.Rotate(i).FlipY()

					if findMatchingStringTile(&k, tiles, t1.FirstRow()) == -1 && findMatchingStringTile(&k, tiles, t1.FirstCol()) == -1 {
						v = t1
						break
					}
					if findMatchingStringTile(&k, tiles, t2.FirstRow()) == -1 && findMatchingStringTile(&k, tiles, t2.FirstCol()) == -1 {
						v = t2
						break
					}
					if findMatchingStringTile(&k, tiles, t3.FirstRow()) == -1 && findMatchingStringTile(&k, tiles, t3.FirstCol()) == -1 {
						v = t3
						break
					}
				}
				image[twod.Position{0, 0}] = v
				del = k
			}
		}
	}

	delete(tiles, del)
	g := fillImage(image, tiles)
	for i := 0; i <= 3; i++ {
		t1 := g.Rotate(i)
		t1c := countSeaMonsters(t1)
		if t1c != -1 {
			return part1, t1c
		}

		t2 := g.Rotate(i).FlipX()
		t2c := countSeaMonsters(t2)
		if t2c != -1 {
			return part1, t2c
		}

		t3 := g.Rotate(i).FlipY()
		t3c := countSeaMonsters(t3)
		if t3c != -1 {
			return part1, t3c
		}
	}
	return part1, part2
}
