package Day16

import (
	"strconv"
	"strings"
	"time"

	"github.com/bznein/AoC2020/pkg/input"
	"github.com/bznein/AoC2020/pkg/timing"
)

type rule map[string]ranges

type ranges struct {
	min1 int
	max1 int
	min2 int
	max2 int
}

func getAssignments(valid []map[string]int, toFind int, alreadyFound map[string]bool, res map[int]string) map[int]string {

	// At each index, look for the list of fields that have count == totFind
	for index, possibilities := range valid {
		found := 0
		candidate := ""
		// For all the fields
		for field, count := range possibilities {
			if _, ok := alreadyFound[field]; ok {
				continue
			}
			if count == toFind {
				found++
				candidate = field
			}
		}
		if found == 1 {
			res[index] = candidate
			alreadyFound[candidate] = true
			return getAssignments(valid, toFind, alreadyFound, res)
		}
	}

	return res
}

func stringToRuleKeyVal(s string) (string, ranges) {
	ruleSplit := strings.Split(s, ": ")
	ruleKey := ruleSplit[0]
	rangeSplit := strings.Split(ruleSplit[1], " or ")
	range1 := strings.Split(rangeSplit[0], "-")
	range2 := strings.Split(rangeSplit[1], "-")
	range1min, _ := strconv.Atoi(range1[0])
	range1max, _ := strconv.Atoi(range1[1])
	range2min, _ := strconv.Atoi(range2[0])
	range2max, _ := strconv.Atoi(range2[1])
	return ruleKey, ranges{
		min1: range1min,
		max1: range1max,
		min2: range2min,
		max2: range2max,
	}
}

func (r *rule) fill(s string) {
	k, v := stringToRuleKeyVal(s)
	(*r)[k] = v
}

func (v ranges) contains(val int) bool {
	return val >= v.min1 && val <= v.max1 || val >= v.min2 && val <= v.max2
}

func Solve(inputF string) (int, int) {
	defer timing.TimeTrack(time.Now())
	part1, part2 := 0, 1
	s := input.InputToStringSlice(inputF)

	rules := rule{}
	nearby := false
	validTickets := 0
	myTicket := []int{}
	var validPositions []map[string]int
	your := false
	for _, ss := range s {
		if !nearby && strings.Contains(ss, "nearby") {
			nearby = true
			continue
		} else if !your && strings.Contains(ss, "your") {
			your = true
			continue
		} else if ss == "" {
			continue
		}
		if !(your || nearby) {
			rules.fill(ss)
		}
		if your {
			valsS := strings.Split(ss, ",")
			for _, v := range valsS {
				vInt, _ := strconv.Atoi(v)
				myTicket = append(myTicket, vInt)
			}
		}
		if nearby {
			valid := true
			vals := strings.Split(ss, ",")
			if validPositions == nil {
				validPositions = make([]map[string]int, len(vals))
				for i := range validPositions {
					validPositions[i] = map[string]int{}
				}
			}
		Vals:
			for _, vv := range vals {
				val, _ := strconv.Atoi(vv)
				for _, v := range rules {
					if v.contains(val) {
						continue Vals
					}
				}
				part1 += val
				valid = false
			}
			if valid {
				validTickets++
				for i, vv := range vals {
					val, _ := strconv.Atoi(vv)
					for k, v := range rules {
						if v.contains(val) {
							validPositions[i][k]++
						}
					}
				}
			}
		}
	}

	assignments := getAssignments(validPositions, validTickets, map[string]bool{}, map[int]string{})

	for i, v := range myTicket {
		if strings.Contains(assignments[i], "departure") {
			part2 *= v
		}
	}

	return part1, part2
}
