package passport

import (
	"strconv"
	"strings"

	"github.com/bznein/AoC2020/pkg/input"
)

const (
	birthYear     = "byr"
	issueYear     = "iyr"
	expYear       = "eyr"
	height        = "hgt"
	hairColour    = "hcl"
	eyeColour     = "ecl"
	pid           = "pid"
	cid           = "cid"
	validPidChars = "abcdef0123456789"
)

var validEyeColours = [...]string{"amb", "blu", "gry", "brn", "grn", "hzl", "oth"}

type passport map[string]string

func StringToPassportSlice(inputD string) []passport {
	result := []passport{}
	stringSlice := input.InputToStringSlice(inputD)
	emptyPassport := passport{}
	for _, s := range stringSlice {
		if s == "" {
			result = append(result, emptyPassport)
			emptyPassport = passport{}
			continue
		}
		for _, field := range strings.Split(s, " ") {
			v := strings.Split(field, ":")
			emptyPassport[v[0]] = v[1]
		}

	}
	result = append(result, emptyPassport)
	return result
}

func (p passport) HasAllRequiredFields() bool {
	return p[birthYear] != "" &&
		p[issueYear] != "" &&
		p[expYear] != "" &&
		p[height] != "" &&
		p[hairColour] != "" &&
		p[eyeColour] != "" &&
		p[pid] != ""
}

func (p passport) heightHasValidMeasure(unitPosition, min, max int) bool {
	if unitPosition != len(p[height])-2 {
		return false
	}
	hgt, err := strconv.Atoi(p[height][:unitPosition])
	if err != nil || hgt < min || hgt > max {
		return false
	}
	return true
}

func (p passport) hasValidHeight() bool {
	s := p[height]
	cm := strings.Index(s, "cm")
	in := strings.Index(s, "in")
	if (in == -1 && cm == -1) || (in != -1 && cm != -1) {
		return false
	}
	if in != -1 {
		return p.heightHasValidMeasure(in, 59, 76)
	}
	if cm != -1 {
		return p.heightHasValidMeasure(cm, 150, 193)
	}
	return true
}

func (p passport) hasValidHairColour() bool {
	s := p[hairColour]
	if len(s) != 7 {
		return false
	}
	if s[0] != '#' {
		return false
	}
	for i := 1; i < 7; i++ {
		if !strings.Contains(validPidChars, s[i:i+1]) {
			return false
		}
	}
	return true
}

func (p passport) hasValidPid() bool {
	s := p[pid]
	if len(s) != 9 {
		return false
	}
	for _, c := range s {
		if _, err := strconv.Atoi(string(c)); err != nil {
			return false
		}
	}
	return true
}

func (p passport) hasValidDate(key string, min int, max int) bool {
	v, err := strconv.Atoi(p[key])
	return err == nil && v >= min && v <= max
}

func (p passport) hasValidEyeColour() bool {
	for _, colour := range validEyeColours {
		if p[eyeColour] == colour {
			return true
		}
	}
	return false
}

func (p passport) IsValid() bool {
	if !p.hasValidDate(birthYear, 1920, 2002) {
		return false
	}
	if !p.hasValidDate(issueYear, 2010, 2020) {
		return false
	}
	if !p.hasValidDate(expYear, 2020, 2030) {
		return false
	}
	if !p.hasValidHeight() {
		return false
	}
	if !p.hasValidHairColour() {
		return false
	}
	if !p.hasValidPid() {
		return false
	}
	if !p.hasValidEyeColour() {
		return false
	}
	return true
}
