package main

import (
	"adventofcode/utils"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func HasMajority(inputs []string, position int) bool {
	count := 0
	for i := 0; i < len(inputs); i++ {
		input := inputs[i]
		bitValue := input[position : position+1]
		if strings.EqualFold("0", bitValue) {
			count++
		}
	}
	majority := len(inputs) / 2
	return count > majority
}

func FilterByPrefix(inputs []string, prefix string) []string {
	filtered := make([]string, 0, len(inputs))
	for i := 0; i < len(inputs); i++ {
		input := inputs[i]
		if strings.HasPrefix(input, prefix) {
			filtered = append(filtered, input)
		}
	}
	return filtered
}

func FindOxygen(inputs []string) string {
	filtered := inputs
	bitPosition := 0
	prefix := ""
	for len(filtered) > 1 {
		if HasMajority(filtered, bitPosition) {
			prefix = prefix + "0"
		} else {
			prefix = prefix + "1"
		}
		filtered = FilterByPrefix(filtered, prefix)
		bitPosition++
	}

	return filtered[0]
}

func FindScrubber(inputs []string) string {
	filtered := inputs
	bitPosition := 0
	prefix := ""
	for len(filtered) > 1 {
		if HasMajority(filtered, bitPosition) {
			prefix = prefix + "1"
		} else {
			prefix = prefix + "0"
		}
		filtered = FilterByPrefix(filtered, prefix)
		bitPosition++
	}

	return filtered[0]
}

func main() {
	inputPath := os.Args[1]
	inputs := utils.AsInputList(utils.ReadInput(inputPath))
	rawOxygen := FindOxygen(inputs)
	rawScrubber := FindScrubber(inputs)
	parsedOxygen, _ := strconv.ParseInt(rawOxygen, 2, 64)
	parsedScrubber, _ := strconv.ParseInt(rawScrubber, 2, 64)
	consumption := parsedOxygen * parsedScrubber
	fmt.Println("Output for partTwo: ", consumption)

}
