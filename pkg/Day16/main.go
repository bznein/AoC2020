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

func getAssignments(valid map[int]map[string]int, toFind int, alreadyFound map[string]bool, res map[int]string) map[int]string {

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

func Solve(inputF string) (int, int) {
	defer timing.TimeTrack(time.Now())
	part1, part2 := 0, 1
	s := input.InputToStringSlice(inputF)

	nearbyTickets := []string{}
	rules := rule{}
	nearby := false
	myTicket := []int{}
	your := false
	for _, ss := range s {
		if strings.Contains(ss, "nearby") {
			nearby = true
			continue
		}
		if strings.Contains(ss, "your") {
			your = true
			continue
		}
		if ss == "" {
			continue
		}
		if !(your || false) {
			ruleSplit := strings.Split(ss, ": ")
			ruleKey := ruleSplit[0]
			rangeSplit := strings.Split(ruleSplit[1], " or ")
			range1 := strings.Split(rangeSplit[0], "-")
			range2 := strings.Split(rangeSplit[1], "-")
			range1min, _ := strconv.Atoi(range1[0])
			range1max, _ := strconv.Atoi(range1[1])
			range2min, _ := strconv.Atoi(range2[0])
			range2max, _ := strconv.Atoi(range2[1])
			rules[ruleKey] = ranges{
				min1: range1min,
				max1: range1max,
				min2: range2min,
				max2: range2max,
			}

		}
		if your {
			valsS := strings.Split(ss, ",")
			for _, v := range valsS {
				vInt, _ := strconv.Atoi(v)
				myTicket = append(myTicket, vInt)
			}
		}
		if nearby {
			nearbyTickets = append(nearbyTickets, ss)
		}
	}

	validTickets := []string{}
	for _, ticket := range nearbyTickets {
		valid := true
		vals := strings.Split(ticket, ",")
	Vals:
		for _, vv := range vals {
			val, _ := strconv.Atoi(vv)
			for _, v := range rules {
				if (val >= v.min1 && val <= v.max1) || (val >= v.min2 && val <= v.max2) {
					continue Vals
				}
			}
			part1 += val
			valid = false
		}
		if valid {
			validTickets = append(validTickets, ticket)
		}
	}

	validPositions := map[int]map[string]int{}

	for _, ticket := range validTickets {
		vals := strings.Split(ticket, ",")
		for i, vv := range vals {
			val, _ := strconv.Atoi(vv)
			for k, v := range rules {
				if (val >= v.min1 && val <= v.max1) || (val >= v.min2 && val <= v.max2) {
					if _, ok := validPositions[i]; !ok {
						validPositions[i] = map[string]int{}
					}
					if _, ok := validPositions[i][k]; !ok {
						validPositions[i][k] = 0
					}
					validPositions[i][k]++
				}
			}
		}
	}

	assignments := getAssignments(validPositions, len(validTickets), map[string]bool{}, map[int]string{})

	for i, v := range myTicket {
		if strings.Contains(assignments[i], "departure") {
			part2 *= v
		}
	}

	return part1, part2
}
