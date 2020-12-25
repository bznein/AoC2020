package main

import (
	"testing"
)

var solveTests = []struct {
	day    int
	part1  int
	part2  int
	part2s string
}{
	{1, 388075, 293450526, ""},
	{2, 474, 745, ""},
	{3, 171, 1206576000, ""},
	{4, 245, 133, ""},
	{5, 955, 569, ""},
	{6, 6782, 3596, ""},
	{7, 248, 57281, ""},
	{8, 1939, 2212, ""},
	{9, 22477624, 2980044, ""},
	{10, 1920, 1511207993344, ""},
	{11, 2346, 2111, ""},
	{12, 1424, 63447, ""},
	{13, 2845, 487905974205117, ""},
	{14, 18630548206046, 4254673508445, ""},
	{15, 211, 2159626, ""},
	{16, 20091, 2325343130651, ""},
	{17, 317, 1692, ""},
	{18, 29839238838303, 201376568795521, ""},
	{19, 222, 339, ""},
	{20, 18482479935793, 2118, ""},
	{21, 2412, -1, "mfp,mgvfmvp,nhdjth,hcdchl,dvkbjh,dcvrf,bcjz,mhnrqp"},
	{22, 34255, 33369, ""},
	{23, 43896725, 2911418906, ""},
	{24, 287, 3636, ""},
	{25, 354320, -1, ""},
}

func TestSolver(t *testing.T) {
	for _, test := range solveTests {
		if test.part2s == "" {
			p1, p2 := solve(test.day)
			if p1 != test.part1 || p2 != test.part2 {
				t.Errorf("Day %d, expected (%d,%d), got (%d,%d)", test.day, test.part1, test.part2, p1, p2)
			}
		} else {
			p1, p2 := solveIntStringOutput(test.day)
			if p1 != test.part1 || p2 != test.part2s {
				t.Errorf("Day %d, expected (%d,%s), got (%d,%s)", test.day, test.part1, test.part2s, p1, p2)
			}

		}
	}
}
