package main

import (
	"adventofcode/utils"
	"fmt"
	"math"
	"os"
	"strconv"
)

type coord struct {
	x int
	y int
}

func BuildRisks(lines []string) (map[coord]int, int, int) {
	risks := make(map[coord]int)
	for y, row := range lines {
		for x, val := range row {
			risk, _ := strconv.Atoi(string(val))
			risks[coord{x: x, y: y}] = risk
		}
	}
	return risks, len(lines[0]), len(lines)
}

func TraverseRisks(risks map[coord]int, width int, height int) map[coord]int {
	cumulative := make(map[coord]int)
	current := coord{x: width - 1, y: height - 1}
	cumulative[current] = risks[current]
	for y := height - 2; y >= 0; y-- {
		current = coord{x: width - 1, y: y}
		below := coord{x: width - 1, y: y + 1}
		cumulative[current] = risks[current] + cumulative[below]
	}
	for x := width - 2; x >= 0; x-- {
		current = coord{x: x, y: height - 1}
		right := coord{x: x + 1, y: height - 1}
		cumulative[current] = risks[current] + cumulative[right]
	}
	for x := width - 2; x >= 0; x-- {
		for y := height - 2; y >= 0; y-- {
			current = coord{x: x, y: y}
			below := coord{x: x, y: y + 1}
			right := coord{x: x + 1, y: y}
			choice := int(math.Min(float64(cumulative[below]), float64(cumulative[right])))
			cumulative[current] = risks[current] + choice
		}
	}
	home := coord{x: 0, y: 0}
	cumulative[home] -= risks[home]
	return cumulative
}

func TraverseForward(risks map[coord]int, width int, height int) map[coord]int {
	cumulative := make(map[coord]int)
	current := coord{x: 0, y: 0}
	cumulative[current] = 0
	for y := 1; y < height; y++ {
		current = coord{x: 0, y: y}
		above := coord{x: 0, y: y - 1}
		cumulative[current] = risks[current] + cumulative[above]
	}
	for x := 1; x < width; x++ {
		current = coord{x: x, y: 0}
		left := coord{x: x - 1, y: 0}
		cumulative[current] = risks[current] + cumulative[left]
	}
	for x := 1; x < width; x++ {
		for y := 1; y < height; y++ {
			current = coord{x: x, y: y}
			above := coord{x: x, y: y - 1}
			left := coord{x: x - 1, y: y}
			choice := int(math.Min(float64(cumulative[above]), float64(cumulative[left])))
			cumulative[current] = risks[current] + choice
		}
	}
	return cumulative
}

func FindMin(unvisisted map[coord]bool, costs map[coord]int) coord {
	best := coord{}
	cost := math.MaxInt
	for k := range unvisisted {
		if costs[k] < cost {
			best = k
			cost = costs[k]
		}
	}
	return best
}

func Dijkstra(risks map[coord]int, width int, height int) map[coord]int {
	unvisited := make(map[coord]bool)
	costs := make(map[coord]int)
	for k := range risks {
		costs[k] = math.MaxInt
		unvisited[k] = true
	}
	home := coord{x: 0, y: 0}
	costs[home] = 0
	for len(unvisited) > 0 {
		current := FindMin(unvisited, costs)
		cost := costs[current]
		neighbors := make([]coord, 0)
		if current.x > 0 {
			neighbors = append(neighbors, coord{x: current.x - 1, y: current.y})
		}
		if current.y > 0 {
			neighbors = append(neighbors, coord{x: current.x, y: current.y - 1})
		}
		if current.x+1 < width {
			neighbors = append(neighbors, coord{x: current.x + 1, y: current.y})
		}
		if current.y+1 < height {
			neighbors = append(neighbors, coord{x: current.x, y: current.y + 1})
		}
		for _, neighbor := range neighbors {
			if unvisited[neighbor] {
				tentative := cost + risks[neighbor]
				if tentative < costs[neighbor] {
					costs[neighbor] = tentative
				}
			}
		}

		delete(unvisited, current)
	}
	return costs
}

func Dump(cumulative map[coord]int, width int, height int, outputRoot string) {
	out, _ := os.Create(fmt.Sprintf("%s/day15/cumulative.tsv", outputRoot))
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			out.WriteString(fmt.Sprintf("%04d\t", cumulative[coord{x: x, y: y}]))
		}
		out.WriteString("\n")
	}
	out.Close()
}

func main() {
	inputPath := os.Args[1]
	lines := utils.AsInputList(utils.ReadInput(inputPath))
	risks, width, height := BuildRisks(lines)
	dijkstra := Dijkstra(risks, width, height)
	Dump(dijkstra, width, height, os.Args[2])
	fmt.Println(dijkstra[coord{x: width - 1, y: height - 1}])
}
