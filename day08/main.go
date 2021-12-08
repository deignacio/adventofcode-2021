package main

import (
	"adventofcode/utils"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func ParseInputs(inputs []string) ([][]string, [][]string) {
	signals := make([][]string, 0)
	values := make([][]string, 0)
	for j := 0; j < len(inputs); j++ {
		parsed := strings.Split(inputs[j], " ")
		signals = append(signals, make([]string, 0))
		values = append(values, make([]string, 0))
		fmt.Println(parsed)
		foundPipe := false
		for i := 0; i < len(parsed); i++ {
			if strings.EqualFold(parsed[i], "|") {
				foundPipe = true
				continue
			}
			if foundPipe {
				values[j] = append(values[j], parsed[i])
			} else {
				signals[j] = append(signals[j], parsed[i])
			}
		}
	}
	return signals, values
}

func MapStems(digit string) map[string]bool {
	explode := strings.Split(digit, "")
	stems := make(map[string]bool, 7)
	for i := 0; i < len(explode); i++ {
		stems[explode[i]] = true
	}
	return stems
}

func Sort(value string) string {
	explode := strings.Split(value, "")
	sort.Strings(explode)
	return strings.Join(explode, "")
}

func HasStems(stems map[string]bool, desired map[string]bool) bool {
	hasDesired := true
	for k, _ := range desired {
		if !stems[k] {
			return false
		}
	}
	return hasDesired
}

func IdentifyDigits(signals []string, values []string) int {
	digits := make(map[int]map[string]bool, 10)
	byStem := make(map[string]int, 10)

	// obvious
	for i := 0; i < len(signals); i++ {
		signal := Sort(signals[i])
		signalLen := len(signal)
		stems := MapStems(signal)
		if signalLen == 2 {
			digits[1] = stems
			byStem[signal] = 1
		} else if signalLen == 3 {
			digits[7] = stems
			byStem[signal] = 7
		} else if signalLen == 4 {
			digits[4] = stems
			byStem[signal] = 4
		} else if signalLen == 7 {
			digits[8] = stems
			byStem[signal] = 8
		}
	}

	for i := 0; i < len(signals); i++ {
		signal := Sort(signals[i])
		signalLen := len(signal)
		stems := MapStems(signal)
		if signalLen == 5 {
			// 2, 3, 5
			// 3 has 1
			if HasStems(stems, digits[1]) {
				digits[3] = stems
				byStem[signal] = 3
				continue
			}
			// 5 has 4 - 1
			desired := make(map[string]bool, 4)
			for k, _ := range digits[4] {
				desired[k] = true
			}
			for k, _ := range digits[1] {
				delete(desired, k)
			}
			if HasStems(stems, desired) {
				digits[5] = stems
				byStem[signal] = 5
				continue
			}
			// 2 has only 2 overlap with 4
			// or, the only one left
			digits[2] = stems
			byStem[signal] = 2
		} else if signalLen == 6 {
			// 0, 6, 9
			if !HasStems(stems, digits[1]) {
				digits[6] = stems
				byStem[signal] = 6
			} else if HasStems(stems, digits[4]) {
				digits[9] = stems
				byStem[signal] = 9
			} else {
				digits[0] = stems
				byStem[signal] = 0
			}
		}
	}

	output := ""
	for i := 0; i < len(values); i++ {
		value := Sort(values[i])
		output += strconv.Itoa(byStem[value])
	}
	outputValue, _ := strconv.Atoi(output)
	return outputValue
}

func main() {
	inputs := utils.AsInputList(utils.ReadInput("/Users/dignacio/Documents/adventofcode/day08/input_part_1"))
	signals, values := ParseInputs(inputs)
	outputs := make([]int, len(signals))
	sum := 0
	for i := 0; i < len(signals); i++ {
		outputs[i] = IdentifyDigits(signals[i], values[i])
		sum += outputs[i]
	}
	fmt.Println(sum)
}
