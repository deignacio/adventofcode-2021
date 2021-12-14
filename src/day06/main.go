package main

import (
	"adventofcode/utils"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ParseInput(input string) []int {
	fish := make([]int, 0)
	parsed := strings.Split(input, ",")
	for i := 0; i < len(parsed); i++ {
		num, _ := strconv.Atoi(parsed[i])
		fish = append(fish, num)
	}
	return fish
}

func DistroDay(distro map[int]int) map[int]int {
	next := make(map[int]int)
	for i := 0; i <= 8; i++ {
		next[i] = 0
	}
	next[8] = distro[0]
	next[6] = distro[0]
	for i := 1; i <= 8; i++ {
		next[i-1] += distro[i]
	}
	return next
}

func main() {
	inputPath := os.Args[1]
	inputs := utils.AsInputList(utils.ReadInput(inputPath))
	fish := ParseInput(inputs[0])
	distro := make(map[int]int)
	for i := 0; i < len(fish); i++ {
		_, present := distro[fish[i]]
		if present {
			distro[fish[i]]++
		} else {
			distro[fish[i]] = 1
		}
	}
	for i := 0; i < 256; i++ {
		distro = DistroDay(distro)
	}
	fmt.Println(distro)
	count := 0
	for _, v := range distro {
		count += v
	}
	fmt.Println(count)
}
