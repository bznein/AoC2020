package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/bznein/AoC2020/pkg/input"
)

type passport struct {
	byr string
	iyr string
	eyr string
	hgt string
	hcl string
	ecl string
	pid string
	cid string
}

func stringToPassportSlice(inputD string) []passport {
	stringSlice := input.InputToStringSlice(inputD)
	result := []passport{}
	emptyPassport := passport{}
	for _, s := range stringSlice {
		if s == "" {
			result = append(result, emptyPassport)
			emptyPassport = passport{}
			continue
		}
		for _, field := range strings.Split(s, " ") {
			v := strings.Split(field, ":")
			key := v[0]
			val := v[1]
			switch key {
			case "ecl":
				emptyPassport.ecl = val
			case "hcl":
				emptyPassport.hcl = val
			case "byr":
				emptyPassport.byr = val
			case "iyr":
				emptyPassport.iyr = val

			case "eyr":
				emptyPassport.eyr = val

			case "hgt":
				emptyPassport.hgt = val

			case "pid":
				emptyPassport.pid = val

			case "cid":
				emptyPassport.cid = val

			}
		}

	}
	result = append(result, emptyPassport)
	return result
}

func (p passport) IsValid() bool {
	return !(p.byr == "" || p.iyr == "" || p.eyr == "" || p.hgt == "" || p.hcl == "" || p.ecl == "" || p.pid == "")
}

func isValidHgt(s string) bool {
	cm := strings.Index(s, "cm")
	in := strings.Index(s, "in")
	if in == -1 && cm == -1 {
		return false
	}
	if in != -1 && cm != -1 {
		return false
	}
	if in != -1 {
		if in != len(s)-2 {
			return false
		}
		hgt, err := strconv.Atoi(s[:in])
		if err != nil || hgt < 59 || hgt > 76 {
			return false
		}
	}
	if cm != -1 {
		if cm != len(s)-2 {
			return false
		}
		hgt, err := strconv.Atoi(s[:cm])
		if err != nil || hgt < 150 || hgt > 193 {
			return false
		}
	}
	return true
}

const alpha = "abcdef0123456789"

func isValidHairColor(s string) bool {
	if len(s) != 7 {
		return false
	}
	if s[0] != '#' {
		return false
	}
	for i := 1; i < 7; i++ {
		if !strings.Contains(alpha, s[i:i+1]) {
			return false
		}
	}
	return true
}

func isValidPid(s string) bool {
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

func (p passport) IsReallyValid() bool {
	/*byr (Birth Year) - four digits; at least 1920 and at most 2002.
	iyr (Issue Year) - four digits; at least 2010 and at most 2020.
	eyr (Expiration Year) - four digits; at least 2020 and at most 2030.
	hgt (Height) - a number followed by either cm or in:
	If cm, the number must be at least 150 and at most 193.
	If in, the number must be at least 59 and at most 76.
	hcl (Hair Color) - a # followed by exactly six characters 0-9 or a-f.
	ecl (Eye Color) - exactly one of: amb blu brn gry grn hzl oth.
	pid (Passport ID) - a nine-digit number, including leading zeroes.
	cid (Country ID) - ignored, missing or not.*/

	b, err := strconv.Atoi(p.byr)
	if err != nil || b < 1920 || b > 2002 {
		//	fmt.Printf("%s is not a valid birth year\n", p.byr)
		return false
	}

	i, err := strconv.Atoi(p.iyr)
	if err != nil || i < 2010 || i > 2020 {
		//	fmt.Printf("%s is not a valid issue year\n", p.iyr)
		return false
	}

	e, err := strconv.Atoi(p.eyr)
	if err != nil || e < 2020 || e > 2030 {
		//	fmt.Printf("%s is not a valid exp year\n", p.eyr)
		return false
	}

	if !isValidHgt(p.hgt) {
		//	fmt.Printf("%s is not a valid hgt\n", p.hgt)
		return false
	}

	if !isValidHairColor(p.hcl) {
		//	fmt.Printf("%s is not a valid hair color year\n", p.hcl)
		return false
	}

	switch p.ecl {
	case "amb":
	case "blu":
	case "gry":
	case "brn":
	case "grn":
	case "hzl":
	case "oth":
	default:
		//	fmt.Printf("%s is not a valid eye color\n", p.ecl)
		return false
	}

	if !isValidPid(p.pid) {
		//	fmt.Printf("%s is not a valid pid\n", p.pid)
		return false
	}

	fmt.Printf("%+v is really valid.\n", p)
	return true
}

func solve(inputF string) (int, int) {
	part1 := 0
	part2 := 0

	passports := stringToPassportSlice(inputF)

	for _, p := range passports {
		if p.IsValid() {
			part1++
			if p.IsReallyValid() {
				part2++
			}
		}
	}

	return part1, part2
}

func main() {
	inputF := input.ReadInput(fmt.Sprintf("../../inputs/4.txt"))
	p1, p2 := solve(inputF)
	fmt.Printf("Part 1: %d, Part2: %d\n", p1, p2)
}
