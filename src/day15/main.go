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

func BuildRisks(lines []string, iterCount int) (map[coord]int, int, int) {
	risks := make(map[coord]int)
	width := len(lines[0])
	height := len(lines)
	for i := 0; i < iterCount; i++ {
		for j := 0; j < iterCount; j++ {
			for y, row := range lines {
				for x, val := range row {
					risk, _ := strconv.Atoi(string(val))
					modified := risk + i + j
					if modified > 9 {
						modified -= 9
					}
					risks[coord{x: i*width + x, y: j*height + y}] = modified
				}
			}
		}
	}
	return risks, iterCount * width, iterCount * height
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
	grey := make(map[coord]bool)
	costs := make(map[coord]int)
	for k := range risks {
		costs[k] = math.MaxInt
		unvisited[k] = true
	}
	home := coord{x: 0, y: 0}
	costs[home] = 0
	grey[home] = true
	for len(grey) > 0 {
		current := FindMin(grey, costs)
		cost := costs[current]
		neighbor := coord{x: current.x - 1, y: current.y}
		if unvisited[neighbor] {
			tentative := cost + risks[neighbor]
			if tentative < costs[neighbor] {
				costs[neighbor] = tentative
				grey[neighbor] = true
			}
		}
		neighbor = coord{x: current.x, y: current.y - 1}
		if unvisited[neighbor] {
			tentative := cost + risks[neighbor]
			if tentative < costs[neighbor] {
				costs[neighbor] = tentative
				grey[neighbor] = true
			}
		}
		neighbor = coord{x: current.x + 1, y: current.y}
		if unvisited[neighbor] {
			tentative := cost + risks[neighbor]
			if tentative < costs[neighbor] {
				costs[neighbor] = tentative
				grey[neighbor] = true
			}
		}
		neighbor = coord{x: current.x, y: current.y + 1}
		if unvisited[neighbor] {
			tentative := cost + risks[neighbor]
			if tentative < costs[neighbor] {
				costs[neighbor] = tentative
				grey[neighbor] = true
			}
		}

		delete(unvisited, current)
		delete(grey, current)
	}
	return costs
}

func Dump(cumulative map[coord]int, width int, height int, outputRoot string, label string) {
	out, _ := os.Create(fmt.Sprintf("%s/day15/%s.tsv", outputRoot, label))
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
	risks, width, height := BuildRisks(lines, 5)
	Dump(risks, width, height, os.Args[2], "risks")
	dijkstra := Dijkstra(risks, width, height)
	Dump(dijkstra, width, height, os.Args[2], "finished")
	fmt.Println(dijkstra[coord{x: width - 1, y: height - 1}])
}
