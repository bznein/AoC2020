package Day19

import (
	"strings"
	"time"

	"github.com/bznein/AoC2020/pkg/input"
	"github.com/bznein/AoC2020/pkg/timing"
	"regexp"
)

type ruleType []string
type regexRuleType []*regexp.Regexp

func getRules(s []string) (ruleType, int) {
	l := 0
	for i, line := range s {
		if line == "" {
			l = i
			break
		}
	}
	rules := make(ruleType, l)
	for i := 0; i < l; i++ {
		split := strings.Split(s[i], ":")
		rules[input.AsInt(split[0])] = "(" + split[1][1:] + ")"
	}
	return rules, l + 1
}

func replaceAndMatch(rules ruleType, messages []string) int {
	regexRules := make(regexRuleType, len(rules))
	replacer := strings.NewReplacer("\"", "", " ", "")
	for k := range rules {
		regexRules[k] = regexp.MustCompile("^" + replacer.Replace(rules[k]) + "$")
	}
	return matches(regexRules[0], messages)
}

func matches(r *regexp.Regexp, messages []string) int {
	tot := 0
	for _, message := range messages {
		if r.MatchString(message) {
			tot++
		}
	}
	return tot
}

func p1(rules ruleType, messages []string, c chan int, re *regexp.Regexp) {
	for {
		finish := true
		for k := range rules {
			rules[k] = re.ReplaceAllStringFunc(rules[k], func(s string) string {
				finish = false
				return rules[input.AsInt(s)]
			})
		}
		if finish {
			c <- replaceAndMatch(rules, messages)
			return
		}
	}
}

func p2(rules ruleType, messages []string, c chan int, re *regexp.Regexp) {

	counts := make([]int, 4)
	maxDepth := 6

	rules[8] = "(42 | 42 8)"
	rules[11] = "(42 31 | 42 11 31)"

	for {
		finish := true

		for k := range rules {
			rules[k] = re.ReplaceAllStringFunc(rules[k], func(s string) string {
				finish = false
				i := input.AsInt(s)
				if i == 8 || i == 11 {
					counts[i-8]++
					if counts[i-8] > maxDepth {
						return "foo"
					}
				}
				return rules[i]
			})
		}

		if finish {
			c <- replaceAndMatch(rules, messages)
			return
		}
	}

}

func Solve(inputF string) (int, int) {
	defer timing.TimeTrack(time.Now())

	part1 := 0
	part2 := 0

	s := input.InputToStringSlice(inputF)
	r, k := getRules(s)
	r2 := make(ruleType, len(r))
	copy(r2, r)
	re := regexp.MustCompile(`\d+`)

	c1 := make(chan int)
	go p1(r, s[k:], c1, re)
	c2 := make(chan int)
	go p2(r2, s[k:], c2, re)

	for {
		if part1 != 0 && part2 != 0 {
			return part1, part2
		}
		select {
		case p := <-c1:
			part1 = p
		case p := <-c2:
			part2 = p
		}
	}

}
