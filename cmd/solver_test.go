package main

import (
	"testing"
)

var solveTests = []struct {
	day   int
	part1 int
	part2 int
}{
	{1, 388075, 293450526},
	{2, 474, 745},
	{3, 171, 1206576000},
	{4, 245, 133},
	{5, 955, 569},
	{6, 6782, 3596},
	{7, 248, 57281},
	{8, 1939, 2212},
	{9, 22477624, 2980044},
	{10, 1920, 1511207993344},
	{11, 2346, 2111},
	{12, 1424, 63447},
	{13, 2845, 487905974205117},
}

func TestSolver(t *testing.T) {
	for _, test := range solveTests {
		p1, p2 := solve(test.day)
		if p1 != test.part1 || p2 != test.part2 {
			t.Errorf("Day %d, expected (%d,%d), got (%d,%d)", test.day, test.part1, test.part2, p1, p2)
		}
	}
}
