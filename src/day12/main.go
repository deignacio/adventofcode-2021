package main

import (
	"adventofcode/utils"
	"fmt"
	"os"
	"strings"
)

type cave struct {
	big  bool
	name string
}

type path struct {
	src cave
	dst cave
}

type network struct {
	head    cave
	options []cave
}

func FilterLines(lines []string) []path {
	paths := make([]path, 0)
	for i := 0; i < len(lines); i++ {
		line := lines[i]
		paths = append(paths, MakePath(line))
	}
	return paths
}

func MakePath(line string) path {
	parts := strings.Split(line, "-")
	return path{src: MakeCave(parts[0]), dst: MakeCave(parts[1])}
}

func MakeCave(name string) cave {
	return cave{name: name, big: strings.ToUpper(name) == name}
}

func BuildNetwork(paths []path) map[string]network {
	caves := make(map[string]network, 0)
	for i := 0; i < len(paths); i++ {
		srcVal, srcPresent := caves[paths[i].src.name]

		if srcPresent {
			nextOptions := append(srcVal.options, paths[i].dst)
			caves[paths[i].src.name] = network{head: paths[i].src, options: nextOptions}
		} else {
			nextOptions := []cave{paths[i].dst}
			caves[paths[i].src.name] = network{head: paths[i].src, options: nextOptions}
		}
		dstVal, dstPresent := caves[paths[i].dst.name]
		if dstPresent {
			nextOptions := append(dstVal.options, paths[i].src)
			caves[paths[i].dst.name] = network{head: paths[i].dst, options: nextOptions}
		} else {
			nextOptions := []cave{paths[i].src}
			caves[paths[i].dst.name] = network{head: paths[i].dst, options: nextOptions}
		}
	}
	return caves
}

func Traverse(caves map[string]network) []string {
	routes := [][]string{{caves["start"].head.name}}
	valid := make([][]string, 0)
	iters := 0
	for len(routes) > 0 && iters < 1000 {
		//fmt.Println(len(routes), routes)
		next := make([][]string, 0)
		for j := range routes {
			route := routes[j]
			tail := route[len(route)-1]
			if tail == "end" {
				valid = append(valid, route)
				continue
			}

			options := caves[tail].options
			for i := range options {
				option := options[i]
				if option.name == "start" {
					continue
				}
				toAdd := make([]string, len(route))
				copy(toAdd, route)
				toAdd = append(toAdd, option.name)
				counts := make(map[string]int, 0)
				for _, stop := range toAdd {
					_, present := counts[stop]
					if present && strings.ToLower(stop) == stop {
						counts[stop] += 1
					} else {
						counts[stop] = 1
					}
				}
				numTwos := 0
				numMore := 0
				for _, count := range counts {
					if count == 2 {
						numTwos++
					} else if count > 2 {
						numMore++
					}
				}
				if numTwos > 1 || numMore > 0 {
					continue
				}
				next = append(next, toAdd)
			}
		}
		routes = next
		iters++
	}
	journeys := make(map[string]string, 0)
	for _, v := range valid {
		journey := strings.Join(v, ",")
		journeys[journey] = journey
	}
	asArray := make([]string, 0)
	for _, journey := range journeys {
		asArray = append(asArray, journey)
	}
	return asArray
}

func main() {
	inputPath := os.Args[1]
	lines := utils.AsInputList(utils.ReadInput(inputPath))
	paths := FilterLines(lines)
	//fmt.Println(paths)
	caves := BuildNetwork(paths)
	//fmt.Println(caves)
	journeys := Traverse(caves)
	fmt.Println(len(journeys))
}
