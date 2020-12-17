package Day17

import (
	"fmt"
	"time"

	"github.com/bznein/AoC2020/pkg/input"
	"github.com/bznein/AoC2020/pkg/timing"
)

const (
	active   = '#'
	inactive = '.'
	//TODO this can be computed!
	smallSize = 7
	size      = 12
)

//TODO move in 3D package
type threeDGrid map[int]map[int]map[int]rune
type fourDGrid map[int]threeDGrid

func (t threeDGrid) print() {
	for z := -smallSize; z <= smallSize; z++ {
		fmt.Printf("z=%d\n", z)
		for i := -size; i <= size; i++ {
			for j := -size; j <= size; j++ {
				fmt.Printf("%s", string(t[z][i][j]))
			}
			fmt.Println("")
		}
	}
}

func (t fourDGrid) print() {
	for w := -smallSize; w <= smallSize; w++ {
		fmt.Printf("W=%d", w)
		t[w].print()
		fmt.Println("")
	}
}

func (t fourDGrid) get80Neighbours(w, z, i, j int) int {
	tot := 0
	for ww := w - 1; ww <= w+1; ww++ {
		for zz := z - 1; zz <= z+1; zz++ {
			for jj := j - 1; jj <= j+1; jj++ {
				for ii := i - 1; ii <= i+1; ii++ {
					if t[ww][zz][ii][jj] == active && !(ww == w && ii == i && jj == j && zz == z) {
						tot++
					}
				}
			}
		}
	}
	return tot
}

func (t threeDGrid) get26Neighbours(z, i, j int) int {
	tot := 0
	for zz := z - 1; zz <= z+1; zz++ {
		for jj := j - 1; jj <= j+1; jj++ {
			for ii := i - 1; ii <= i+1; ii++ {
				if t[zz][ii][jj] == active && !(ii == i && jj == j && zz == z) {
					tot++
				}
			}
		}
	}
	return tot
}

func (t threeDGrid) totalActive() int {
	tot := 0
	for x := -size; x <= size; x++ {
		for y := -size; y <= size; y++ {
			for z := -size; z <= size; z++ {
				if t[x][y][z] == active {
					tot++
				}
			}
		}
	}
	return tot
}

func (source threeDGrid) copy() threeDGrid {
	t := threeDGrid{}
	for x := -size; x <= size; x++ {
		if _, ok := t[x]; !ok {

			t[x] = map[int]map[int]rune{}
		}
		for y := -size; y <= size; y++ {
			if _, ok := t[x][y]; !ok {
				t[x][y] = map[int]rune{}
			}
			for z := -size; z <= size; z++ {
				t[x][y][z] = source[x][y][z]
			}
		}
	}
	return t
}

func (source fourDGrid) copy() fourDGrid {
	t := fourDGrid{}
	for w := -smallSize; w <= smallSize; w++ {
		t[w] = source[w].copy()
	}
	return t
}

func (source fourDGrid) totalActive() int {
	tot := 0
	for w := -smallSize; w <= smallSize; w++ {
		tot += source[w].totalActive()
	}
	return tot
}

func (t threeDGrid) oneStep() threeDGrid {
	copyMap := t.copy()
	for z := -size + 1; z < size; z++ {
		for i := -size + 1; i < size; i++ {
			for j := -size + 1; j < size; j++ {
				nActive := t.get26Neighbours(z, i, j)
				if t[z][i][j] == active {
					if nActive != 2 && nActive != 3 {
						copyMap[z][i][j] = inactive
					}
				} else {
					if nActive == 3 {
						copyMap[z][i][j] = active
					}
				}
			}
		}
	}
	return copyMap
}

func (t fourDGrid) oneStep() fourDGrid {
	copyMap := t.copy()
	for w := -smallSize + 1; w < smallSize; w++ {
		for z := -size + 1; z < size; z++ {
			for i := -size + 1; i < size; i++ {
				for j := -size + 1; j < size; j++ {
					nActive := t.get80Neighbours(w, z, i, j)
					if t[w][z][i][j] == active {
						if nActive != 2 && nActive != 3 {
							copyMap[w][z][i][j] = inactive
						}
					} else {
						if nActive == 3 {
							copyMap[w][z][i][j] = active
						}
					}
				}
			}
		}
	}
	return copyMap
}

func Solve(inputF string) (int, int) {
	defer timing.TimeTrack(time.Now())

	n := input.InputToStringSlice(inputF)

	t := threeDGrid{}
	t2 := fourDGrid{}

	for w := -smallSize; w <= smallSize; w++ {
		if _, ok := t2[w]; !ok {
			t2[w] = threeDGrid{}
		}
		for x := -size; x <= size; x++ {
			if _, ok := t[x]; !ok {
				t[x] = map[int]map[int]rune{}
			}
			if _, ok := t2[w][x]; !ok {
				t2[w][x] = map[int]map[int]rune{}
			}
			for y := -size; y <= size; y++ {
				if _, ok := t[x][y]; !ok {
					t[x][y] = map[int]rune{}
				}
				if _, ok := t2[w][x][y]; !ok {
					t2[w][x][y] = map[int]rune{}
				}
				for z := -size; z <= size; z++ {
					t[x][y][z] = inactive
					t2[w][x][y][z] = inactive
				}
			}

		}
	}

	for i := range n {
		for j := range n[i] {
			if n[i][j] == active {
				t[0][i][j] = rune(n[i][j])
				t2[0][0][i][j] = rune(n[i][j])
			}
		}
	}

	for i := 0; i < 6; i++ {
		t = t.oneStep()
		t2 = t2.oneStep()
	}

	return t.totalActive(), t2.totalActive()
}
