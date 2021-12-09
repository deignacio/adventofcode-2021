package main

import (
	"adventofcode/utils"
	"fmt"
	"strconv"
	"strings"
)

func ParseInputs(inputs []string) [][]int {
	values := make([][]int, 0)
	for i := 0; i < len(inputs); i++ {
		input := strings.Split(inputs[i], "")
		value := make([]int, 0)
		for j := 0; j < len(input); j++ {
			val, _ := strconv.Atoi(input[j])
			value = append(value, val)
		}
		values = append(values, value)
	}
	return values
}

func IsLow(depths [][]int, x int, y int) bool {
	current := depths[x][y]
	if x > 0 {
		if depths[x-1][y] <= current {
			return false
		}
	}
	if x + 1 < len(depths) {
		if depths[x+1][y] <= current {
			return false
		}
	}
	if y > 0 {
		if depths[x][y-1] <= current {
			return false
		}
	}
	if y + 1 < len(depths[x]) {
		if depths[x][y+1] <= current {
			return false
		}
	}
	return true
}



func main() {
	inputs := utils.AsInputList(utils.ReadInput("/Users/dignacio/Documents/adventofcode/day09/input_part_1"))
	depths := ParseInputs(inputs)
	lows := make([]int, 0)
	for i := 0; i < len(depths); i++ {
		row := depths[i]
		for j := 0; j < len(row); j++ {
			if IsLow(depths, i, j) {
				lows = append(lows, row[j])
			}
		}
	}
	fmt.Println(lows)
	sum := 0
	for i := 0; i < len(lows); i++ {
		sum += lows[i] + 1
	}
	fmt.Println(sum)
}
