package main

import (
	"adventofcode/utils"
	"fmt"
	"strings"
)

type cave struct {
	big bool
	name string
}

type path struct {
	src cave
	dst cave
}

type network struct {
	head cave
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
	for len(routes) > 0 && iters < 100 {
		//fmt.Println(len(routes), routes)
		next := make([][]string, 0)
		for j := range routes {
			route := routes[j]
			tail := route[len(route) - 1]
			if tail == "end" {
				valid = append(valid, route)
				continue
			}
			options := caves[tail].options
			for i := range options {
				option := options[i]
				found := false
				for _, existing := range route {
					if option.name == existing && !caves[existing].head.big {
						found = true
					}
				}
				if !found || option.big {
					toAdd := make([]string, len(route))
					copy(toAdd, route)
					toAdd = append(toAdd, option.name)
					next = append(next, toAdd)
				}
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
	lines := utils.AsInputList(utils.ReadInput("/Users/dignacio/Documents/adventofcode/day12/input_part_1"))
	paths := FilterLines(lines)
	//fmt.Println(paths)
	caves := BuildNetwork(paths)
	//fmt.Println(caves)
	journeys := Traverse(caves)
	fmt.Println(len(journeys))
}
