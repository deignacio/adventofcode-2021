package main

import (
	"adventofcode/utils"
	"fmt"
	"strconv"
	"strings"
)


func ParseCommand(input string) (string, int) {
	raw := strings.Split(input, " ")
	magnitude, _ := strconv.Atoi(raw[1])
	return raw[0], magnitude
}

func ParseCommands(inputs []string) (int, int) {
	var depth, distance, aim = 0, 0, 0
	for i:= 0; i < len(inputs); i++ {
		direction, magnitude := ParseCommand(inputs[i])
		if strings.EqualFold(direction, "forward") {
			distance += magnitude
			depth += aim * magnitude
		} else if strings.EqualFold(direction, "down") {
			aim += magnitude
		} else if strings.EqualFold(direction, "up") {
			aim -= magnitude
		}
		fmt.Println("depth", depth, "distance", distance, "aim", aim)
	}

	return depth, distance
}

func main() {
	inputs := utils.AsInputList(utils.ReadInput("/Users/dignacio/Documents/adventofcode/day02/input_part_1"))
	depth, distance := ParseCommands(inputs)
	fmt.Println("Output for partTwo: ", depth * distance)
}
