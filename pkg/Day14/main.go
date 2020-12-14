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

func getMask1(s string) map[int]rune {
	m := map[int]rune{}
	for i, v := range strings.Split(s, " = ")[1] {
		if _, err := strconv.Atoi(string(v)); err == nil {
			m[i] = v
		}
	}
	return m
}

func getMask2(s string) map[int]rune {
	m := map[int]rune{}
	for i, v := range strings.Split(s, " = ")[1] {
		if v != '0' {
			m[i] = v
		}
	}
	return m
}

func p1(n []string) int {

	values := map[int]int64{}
	var m map[int]rune
	for i := 0; i < len(n); i++ {
		if n[i][1] == 'a' {
			//stupid, fast way to see when it is a mask
			m = getMask1(n[i])
			continue
		}
		cell := 0
		val := 0
		fmt.Sscanf(n[i], "mem[%d] = %d", &cell, &val)
		valBinary := fmt.Sprintf("%036s", strconv.FormatInt(int64(val), 2))
		valBinaryReplaced := s.ReplaceAtMap(valBinary, m)
		values[cell], _ = strconv.ParseInt(valBinaryReplaced, 2, 64)
	}

	part1 := 0
	for _, v := range values {
		part1 += int(v)
	}
	return part1
}

func getAllPossibleStrings(ss string) []string {
	found := false
	retVal := []string{}
	for i, v := range ss {
		if v == 'X' {
			retVal = append(retVal, getAllPossibleStrings(s.ReplaceAtIndex(ss, '0', i))...)
			retVal = append(retVal, getAllPossibleStrings(s.ReplaceAtIndex(ss, '1', i))...)
			found = true
			break
		}
	}
	if !found {
		retVal = append(retVal, ss)
	}
	return retVal
}

func p2(n []string) int {

	values := map[int]int64{}
	var m map[int]rune
	for i := 0; i < len(n); i++ {
		if n[i][1] == 'a' {
			//stupid, fast way to see when it is a mask
			m = getMask2(n[i])
			continue
		}
		cell := 0
		val := 0
		fmt.Sscanf(n[i], "mem[%d] = %d", &cell, &val)
		cellBinary := fmt.Sprintf("%036s", strconv.FormatInt(int64(cell), 2))
		cellBinaryReplaced := s.ReplaceAtMap(cellBinary, m)
		for _, r := range getAllPossibleStrings(cellBinaryReplaced) {
			address, _ := strconv.ParseInt(r, 2, 64)
			values[int(address)] = int64(val)
		}
	}

	part2 := 0
	for _, v := range values {
		part2 += int(v)
	}
	return part2
}

func Solve(inputF string) (int, int) {
	defer timing.TimeTrack(time.Now())
	part1, part2 := 0, -1
	n := input.InputToStringSlice(inputF)
	part1 = p1(n)
	part2 = p2(n)

	return part1, part2
}
