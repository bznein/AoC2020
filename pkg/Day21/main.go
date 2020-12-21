package Day21

import (
	"sort"
	"strings"
	"time"

	"github.com/bznein/AoC2020/pkg/input"
	"github.com/bznein/AoC2020/pkg/timing"
)

func contains(ss []string, s string) int {
	for i, sss := range ss {
		if s == sss {
			return i
		}
	}
	return -1
}

func remove(s []string, i int) []string {
	s[len(s)-1], s[i] = s[i], s[len(s)-1]
	return s[:len(s)-1]
}

func fillAllergenes(maybe map[string][]string, temp map[string]string) map[string]string {
	if len(maybe) == 1 {
		return temp
	}
	for k, v := range maybe {
		if len(v) == 1 {
			temp[v[0]] = k
			delete(maybe, k)
			for k2, v2 := range maybe {
				if idx := contains(v2, v[0]); idx != -1 {
					maybe[k2] = remove(v2, idx)
				}
			}
			return fillAllergenes(maybe, temp)
		}
	}
	return temp
}

func Solve(inputF string) (int, string) {
	defer timing.TimeTrack(time.Now())
	part1 := 0

	allIngredients := map[string]bool{}
	mayBeContainedIn := map[string][]string{}
	s := input.InputToStringSlice(inputF)
	for _, ss := range s {
		ingSplit := strings.Split(ss, " (")
		ingredients := strings.Split(ingSplit[0], " ")
		for _, i := range ingredients {
			allIngredients[i] = true
		}
		l := len(ingSplit[1])
		allergenes := strings.Split(ingSplit[1][9:l-1], ", ")
		for _, all := range allergenes {
			if _, ok := mayBeContainedIn[all]; !ok {
				mayBeContainedIn[all] = ingredients
			} else {
				newSlice := make([]string, 0, len(ingredients))
				for _, ing := range mayBeContainedIn[all] {
					if idx := contains(ingredients, ing); idx != -1 {
						newSlice = append(newSlice, ing)
					}
				}
				mayBeContainedIn[all] = newSlice
			}
		}
	}

	notInAny := []string{}
Search:
	for v := range allIngredients {

		for _, all := range mayBeContainedIn {
			for _, ingredient := range all {
				if ingredient == v {
					continue Search
				}
			}
		}
		notInAny = append(notInAny, v)
	}

	for _, ss := range s {
		ingSplit := strings.Split(ss, " (")
		ingredients := strings.Split(ingSplit[0], " ")
		for _, ing := range ingredients {
			if contains(notInAny, ing) != -1 {
				part1++
			}
		}
	}

	finalAllergenes := fillAllergenes(mayBeContainedIn, map[string]string{})

	finalIng := []string{}
	for v := range finalAllergenes {
		finalIng = append(finalIng, v)
	}
	sort.Slice(finalIng, func(i, j int) bool {
		return finalAllergenes[finalIng[i]] < finalAllergenes[finalIng[j]]
	})

	return part1, strings.Join(finalIng, ",")
}
