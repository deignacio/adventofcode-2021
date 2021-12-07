package main

import (
	"adventofcode/utils"
	"fmt"
	"math"
	"strconv"
	"strings"
)

func ParseInput(input string) []int {
	crabs := make([]int, 0)
	parsed := strings.Split(input, ",")
	for i := 0; i < len(parsed); i++ {
		num, _ := strconv.Atoi(parsed[i])
		crabs = append(crabs, num)
	}
	return crabs
}

func FindCost(crabs []int, pos int) int {
	sum := 0
	for i := 0; i < len(crabs); i++ {
		seed := int(math.Abs(float64(crabs[i] - pos)))
		sum += TriNum(seed)
	}
	return sum
}

func TriNum(seed int) int {
	sum := 0
	for i := 0; i <= seed; i++ {
		sum += i
	}
	return sum
}

func main() {
	inputs := utils.AsInputList(utils.ReadInput("/Users/dignacio/Documents/adventofcode/day07/input_part_1"))
	crabs := ParseInput(inputs[0])
	min := math.MaxInt
	max := math.MinInt
	for i := 0; i < len(crabs); i++ {
		if crabs[i] < min {
			min = crabs[i]
		}
		if crabs[i] > max {
			max = crabs[i]
		}
	}
	costs := make([]int, max - min + 1)
	lowestCost := math.MaxInt
	for i := 0; i <= max - min; i++ {
		cost := FindCost(crabs, min + i)
		if cost < lowestCost {
			lowestCost = cost
		}
		costs[i] = cost
	}
	fmt.Println(costs)
	fmt.Println(lowestCost)
}
