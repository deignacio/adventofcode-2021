package main

import (
	"adventofcode/utils"
	"fmt"
	"math"
	"os"
	"strings"
)

func BuildRules(lines []string) map[string]string {
	rules := make(map[string]string)
	for _, rule := range lines {
		split := strings.Split(rule, " ")
		step := split[0]
		rules[step] = split[2]
	}
	return rules
}

func TranslatePolymer(polymer string) (map[string]int, byte, byte) {
	generation := make(map[string]int)
	for i := 0; i+1 < len(polymer); i++ {
		step := polymer[i:i+2]
		_, present := generation[step]
		if present {
			generation[step]++
		} else {
			generation[step] = 1
		}
	}
	return generation, polymer[0], polymer[len(polymer)-1]
}

func LanternPolymerIteration(generation map[string]int, rules map[string]string) map[string]int {
	next := make(map[string]int)
	for step, count := range generation {
		rule, hasRule := rules[step]
		if hasRule {
			head := step[0:1] + rule
			_, hasHead := next[head]
			if hasHead {
				next[head] += count
			} else {
				next[head] = count
			}
			tail := rule + step[1:2]
			_, hasTail := next[tail]
			if hasTail {
				next[tail] += count
			} else {
				next[tail] = count
			}
		}

	}
	return next
}

func CountElements(generation map[string]int, first byte, last byte) map[byte]int {
	counts := make(map[byte]int)
	counts[first] = 1
	counts[last] = 1
	for step, count := range generation {
		head := step[0]
		_, hasHead := counts[head]
		if hasHead {
			counts[head] += count
		} else {
			counts[head] = count
		}
		tail := step[1]
		_, hasTail := counts[tail]
		if hasTail {
			counts[tail] += count
		} else {
			counts[tail] = count
		}
	}
	for step := range counts {
		counts[step] /= 2
	}
	return counts
}

func main() {
	inputPath := os.Args[1]
	lines := utils.AsInputList(utils.ReadInput(inputPath))
	generation, first, last := TranslatePolymer(lines[0])
	rules := BuildRules(lines[2:])
	numSteps := 40
	fmt.Println(0, generation)
	for i := 1; i <= numSteps; i++ {
		generation = LanternPolymerIteration(generation, rules)
		fmt.Println(i, generation)
	}
	counts := CountElements(generation, first, last)
	max := math.MinInt
	min := math.MaxInt
	for _, val := range counts {
		if val > max {
			max = val
		}
		if val < min {
			min = val
		}
	}
	fmt.Println(counts, max, min, max-min)
}
