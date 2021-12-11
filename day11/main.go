package main

import (
	"adventofcode/utils"
	"fmt"
	"strconv"
)

func FilterLines(lines []string) [][]int {
	octopuses := make([][]int, 0)
	for i := 0; i < len(lines); i++ {
		line := lines[i]
		row := make([]int, len(line))
		for j := 0; j < len(line); j++ {
			value, _ := strconv.Atoi(line[j:j+1])
			row[j] = value
		}
		octopuses = append(octopuses, row)
	}
	return octopuses
}

func OneStep(octopuses [][]int) int {
	flashes := 0
	for i := 0; i < len(octopuses); i++ {
		row := octopuses[i]
		for j := 0; j < len(row); j++ {
			flashes += Increment(octopuses, i, j)
		}
	}
	for i := 0; i < len(octopuses); i++ {
		row := octopuses[i]
		for j := 0; j < len(row); j++ {
			if row[j] > 9 {
				row[j] = 0
			}
		}
	}
	return flashes
}

func Increment(octopuses [][]int, x int, y int) int {
	if x < 0 ||
		y < 0 ||
		x >= len(octopuses) ||
		y >= len(octopuses[x]) {
		return 0
	}
	octopuses[x][y]++
	if octopuses[x][y] == 10 {
		return 1 + Flash(octopuses, x, y)
	}
	return 0
}

func Flash(octopuses [][]int, x int, y int) int {
	flashes := 0
	flashes += Increment(octopuses, x-1, y-1)
	flashes += Increment(octopuses, x-1, y)
	flashes += Increment(octopuses, x-1, y+1)
	flashes += Increment(octopuses, x, y-1)
	flashes += Increment(octopuses, x, y+1)
	flashes += Increment(octopuses, x+1, y-1)
	flashes += Increment(octopuses, x+1, y)
	flashes += Increment(octopuses, x+1, y+1)
	return flashes
}

func main() {
	lines := utils.AsInputList(utils.ReadInput("/Users/dignacio/Documents/adventofcode/day11/input_part_1"))
	octopuses := FilterLines(lines)
	total := 0
	first := -1
	size := len(octopuses) * len(octopuses[0])
	fmt.Println(size)
	fmt.Println(total, octopuses)
	for i := 1; i < 1000000; i++ {
		flashes := OneStep(octopuses)
		total += flashes
		if flashes >= size {
			first = i
			break
		}
	}
	fmt.Println(total, first, octopuses)
}
