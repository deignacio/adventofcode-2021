package main

import (
	"adventofcode/utils"
	"fmt"
	"os"
	"strings"
)

func BuildRules(lines []string) map[string]string {
	rules := make(map[string]string)
	for _, rule := range lines {
		split := strings.Split(rule, " ")
		rules[split[0]] = split[2]
	}
	return rules
}

func ApplyRules(polymer string, rules map[string]string) string {
	next := ""
	current := polymer[0:1]
	for _, c := range polymer[1:] {
		end := string(c)
		next = next + current
		candidate := current + end
		toInsert, present := rules[candidate]
		if present {
			next = next + toInsert
		}
		current = end
	}
	next = next + string(polymer[len(polymer)-1])
	return next
}

func CountPairs(polymer string) map[string]int {
	count := make(map[string]int)
	for _, e := range polymer {
		count[string(e)]++
	}
	return count
}

func main() {
	inputPath := os.Args[1]
	lines := utils.AsInputList(utils.ReadInput(inputPath))
	polymer := lines[0]
	rules := BuildRules(lines[2:])
	fmt.Println(polymer, rules)
	numSteps := 10
	for i := 0; i < numSteps; i++ {
		polymer = ApplyRules(polymer, rules)
		// fmt.Println(polymer)
		fmt.Println(i, len(polymer))
	}
	count := CountPairs(polymer)
	fmt.Println(count)
}
