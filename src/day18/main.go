package main

import (
	"adventofcode/utils"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type snailfish map[string]int

func ParseInputs(inputs []string) []snailfish {
	fish := make([]snailfish, 0)
	for _, raw := range inputs {
		fish = append(fish, ParseSnailfish(raw))
	}
	return fish
}

func ParseSnailfish(raw string) snailfish {
	prefix := ""
	s := make(snailfish)
	for _, t := range raw {
		if t == '[' {
			prefix += "L"
		} else if t == ',' {
			prefix = prefix[:len(prefix)-1] + "R"
		} else if t == ']' {
			prefix = prefix[:len(prefix)-1]
		} else {
			i, _ := strconv.Atoi(string(t))
			s[prefix] = i
		}
	}
	return s
}

func Add(a0 snailfish, a1 snailfish) snailfish {
	next := make(map[string]int)
	for k, v := range a0 {
		next["L"+k] = v
	}
	for k, v := range a1 {
		next["R"+k] = v
	}
	return next
}

func Merge(m0 snailfish, m1 snailfish) snailfish {
	next := make(map[string]int)
	for k, v := range m0 {
		next[k] = v
	}
	for k, v := range m1 {
		next[k] = v
	}
	return next
}

func FindLeft(s snailfish, prefix string) (string, int, bool) {
	lastRight := strings.LastIndex(prefix, "R")
	if lastRight == -1 {
		return "", -1, false
	}

	key := prefix[:lastRight] + "L"
	value, found := s[key]
	for !found {
		key += "R"
		value, found = s[key]
	}
	return key, value, found
}

func FindRight(s snailfish, prefix string) (string, int, bool) {
	lastLeft := strings.LastIndex(prefix, "L")
	if lastLeft == -1 {
		return "", -1, false
	}

	key := prefix[:lastLeft] + "R"
	value, found := s[key]
	for !found {
		key += "L"
		value, found = s[key]
	}
	return key, value, found
}

func Explode(s snailfish, prefix string) bool {
	lKey := prefix + "L"
	rKey := prefix + "R"
	lVal, lPres := s[lKey]
	rVal, rPres := s[rKey]

	if lPres && rPres && len(lKey) > 4 {
		fmt.Println("exploding ", lKey, rKey)
		delete(s, lKey)
		lnKey, _, lnPres := FindLeft(s, lKey)
		if lnPres {
			fmt.Println("\tincrementing left neighbor", lnKey)
			s[lnKey] += lVal
		}
		delete(s, rKey)
		rnKey, _, rnPres := FindRight(s, rKey)
		if rnPres {
			fmt.Println("\tincrementing right neighbor", rnKey)
			s[rnKey] += rVal
		}
		s[prefix] = 0
		return true
	}

	_, present := s[prefix]
	if !present {
		leftModified := Explode(s, prefix+"L")
		if leftModified {
			return true
		}
		rightModified := Explode(s, prefix+"R")
		return rightModified
	}

	return false
}

func Splits(s snailfish, prefix string) bool {
	value, present := s[prefix]
	if !present {
		leftModified := Splits(s, prefix+"L")
		if leftModified {
			return true
		}
		rightModified := Splits(s, prefix+"R")
		return rightModified
	}

	if value >= 10 {
		// split
		fmt.Println("splitting", prefix, value)
		delete(s, prefix)
		if value%2 == 0 {
			s[prefix+"L"] = value / 2
			s[prefix+"R"] = value / 2
		} else {
			s[prefix+"L"] = value / 2
			s[prefix+"R"] = value/2 + 1
		}
		return true
	}

	return false
}

func Magnitude(s snailfish, prefix string) int {
	value, present := s[prefix]
	if !present {
		lVal := Magnitude(s, prefix+"L")
		rVal := Magnitude(s, prefix+"R")
		return 3*lVal + 2*rVal
	}

	return value
}

func Reduce(s snailfish) {
	check := true
	for check {
		fmt.Println("reducing", s)
		check = Explode(s, "")
		if !check {
			check = Splits(s, "")
		}
	}
}

func main() {
	inputPath := os.Args[1]
	lines := utils.AsInputList(utils.ReadInput(inputPath))
	fish := ParseInputs(lines)
	head := fish[0]
	fmt.Println(head)
	for _, f := range fish[1:] {
		fmt.Println("adding", head, f)
		head = Add(head, f)
		Reduce(head)
	}
	fmt.Println(head, Magnitude(head, ""))

	second := ParseInputs(lines)
	head = second[0]
	mags := make(map[string]int)
	sums := make(map[string]snailfish)
	for i := 0; i < len(second); i++ {
		for j := 0; j < len(second); j++ {
			if i == j {
				continue
			}
			key := fmt.Sprintf("%d-%d", i, j)
			sum := Add(second[i], second[j])
			Reduce(sum)
			mag := Magnitude(sum, "")
			fmt.Println(key, sum, mag)
			mags[key] = mag
			sums[key] = sum
		}
	}
	fmt.Println(mags)
	max := 0
	maxKey := ""
	for k, v := range mags {
		if v > max {
			maxKey = k
			max = v
		}
	}
	fmt.Println(maxKey, max, sums[maxKey])
}
