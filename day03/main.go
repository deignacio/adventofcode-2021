package main

import (
	"adventofcode/utils"
	"fmt"
	"math"
	"strings"
)


func BuildGamma(counts map[int]int, size int) float64 {
	majority := size / 2
	var value float64 = 0
	for i := 0; i < size; i++ {
		count, present := counts[i]
		if present && count > majority {
			value += math.Pow(2, float64(i))
		}
	}
	return value
}

func BuildEpsilon(counts map[int]int, size int) float64 {
	majority := size / 2
	var value float64 = 0
	for i := 0; i < size; i++ {
		count, present := counts[i]
		if present && count < majority {
			value += math.Pow(2, float64(i))
		}
	}
	return value
}

func ParseCommands(inputs []string) map[int]int {
	counts := make(map[int]int)
	for i:= 0; i < len(inputs); i++ {
		input := inputs[i]
		inputLen := len(input)
		for j:= 0; j < inputLen; j++ {
			bitPosition := inputLen - j - 1
			bitValue := input[j:j+1]
			_, present := counts[bitPosition]
			if strings.EqualFold("0", bitValue) {
				if present {
					counts[bitPosition]++
				} else {
					counts[bitPosition] = 1
				}
			}
		}
	}

	return counts
}

func main() {
	inputs := utils.AsInputList(utils.ReadInput("/Users/dignacio/Documents/adventofcode/day03/input_part_1"))
	counts := ParseCommands(inputs)
	numInputs := len(inputs)
	fmt.Println(counts)
	gamma := int(BuildGamma(counts, numInputs))
	epsilon := int(BuildEpsilon(counts, numInputs))
	fmt.Println("Output for partOne: ", gamma * epsilon)


}
