package main

import (
	"adventofcode/utils"
	"container/ring"
	"fmt"
	"math"
	"strconv"
)

func CountIncreasing(inputs []string, windowSize int) int {
	var increasingCount int
	window := ring.New(windowSize)
	var previousSum int64 = math.MaxInt64
	for i:= 0; i < len(inputs); i++ {
		current, _ := strconv.ParseInt(inputs[i], 0, 64)
		window.Value = current
		window = window.Next()
		var currentSum int64 = 0
		window.Do(func(p interface{}) {
			// This is how we invalidate the first two windows that are too small
			if p == nil {
				currentSum = math.MinInt64
			} else {
				currentSum += p.(int64)
			}
		})

		if currentSum > previousSum && previousSum > 0 {
			increasingCount++
		}

		previousSum = currentSum
	}

	return increasingCount
}

func main() {
	inputs := utils.AsInputList(utils.ReadInput("/Users/dignacio/Documents/adventofcode/day01/input_part_1"))
	partOne := CountIncreasing(inputs, 1)
	fmt.Println("Output for partOne: ", partOne)
	partTwo := CountIncreasing(inputs, 3)
	fmt.Println("Output for partTwo: ", partTwo)
}
