package main

import (
	"adventofcode/utils"
	"fmt"
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

func main() {
	inputs := utils.AsInputList(utils.ReadInput("/Users/dignacio/Documents/adventofcode/day08/input_part_1"))
	signals, values := ParseInputs(inputs)
	found := 0
	for i := 0; i < len(signals); i++ {
		for j := 0; j < len(signals[i]); j++ {
			signalLen := len(signals[i][j])
			if signalLen == 2 ||
				signalLen == 3 ||
				signalLen == 4 ||
				signalLen == 7 {
				found++
			}
		}
	}
	for i := 0; i < len(values); i++ {
		for j := 0; j < len(values[i]); j++ {
			valueLen := len(values[i][j])
			if valueLen == 2 ||
				valueLen == 3 ||
				valueLen == 4 ||
				valueLen == 7 {
				found++
			}
		}
	}
	fmt.Println(len(signals), len(values))
	fmt.Println(found)
}
