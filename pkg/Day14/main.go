package Day14

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/bznein/AoC2020/pkg/input"
	s "github.com/bznein/AoC2020/pkg/strings"
	"github.com/bznein/AoC2020/pkg/timing"
)

func getMasks(s string) (map[int]rune, map[int]rune) {
	m1 := map[int]rune{}
	m2 := map[int]rune{}
	for i, v := range strings.Split(s, " = ")[1] {
		if _, err := strconv.Atoi(string(v)); err == nil {
			m1[i] = v
		}
		if v != '0' {
			m2[i] = v
		}
	}
	return m1, m2
}

func getAllPossibleStrings(ss string) []string {
	idx := -1
	for i, v := range ss {
		if v == 'X' {
			idx = i
			break
		}
	}
	if idx == -1 {
		return []string{ss}
	}

	return append(getAllPossibleStrings(s.ReplaceAtIndex(ss, '0', idx)), getAllPossibleStrings(s.ReplaceAtIndex(ss, '1', idx))...)
}

func Solve(inputF string) (int, int) {
	defer timing.TimeTrack(time.Now())
	n := input.InputToStringSlice(inputF)

	values1 := map[int]int64{}
	var m1 map[int]rune
	values2 := map[int]int64{}
	var m2 map[int]rune

	for i := 0; i < len(n); i++ {
		if n[i][1] == 'a' {
			//stupid, fast way to see when it is a mask
			m1, m2 = getMasks(n[i])
			continue
		}
		cell := 0
		val := 0
		fmt.Sscanf(n[i], "mem[%d] = %d", &cell, &val)
		valBinary := fmt.Sprintf("%036s", strconv.FormatInt(int64(val), 2))
		valBinaryReplaced := s.ReplaceAtMap(valBinary, m1)
		values1[cell], _ = strconv.ParseInt(valBinaryReplaced, 2, 64)
		cellBinary := fmt.Sprintf("%036s", strconv.FormatInt(int64(cell), 2))
		cellBinaryReplaced := s.ReplaceAtMap(cellBinary, m2)
		for _, r := range getAllPossibleStrings(cellBinaryReplaced) {
			address, _ := strconv.ParseInt(r, 2, 64)
			values2[int(address)] = int64(val)
		}
	}

	part1 := 0
	for _, v := range values1 {
		part1 += int(v)
	}
	part2 := 0
	for _, v := range values2 {
		part2 += int(v)
	}
	return part1, part2
}
