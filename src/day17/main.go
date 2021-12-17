package main

import (
	"adventofcode/utils"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type aim struct {
	x int
	y int
}

type trajectory struct {
	dx int
	dy int
}

type coord struct {
	x int
	y int
}

type target struct {
	minX int
	maxX int
	minY int
	maxY int
}

func GetTarget(raw string) target {
	words := strings.Split(raw, " ")
	xWord := words[2]
	xWords := strings.Split(xWord[2:len(xWord)-1], ".")
	minX, _ := strconv.Atoi(xWords[0])
	maxX, _ := strconv.Atoi(xWords[2])
	yWord := words[3]
	yWords := strings.Split(yWord[2:], ".")
	minY, _ := strconv.Atoi(yWords[0])
	maxY, _ := strconv.Atoi(yWords[2])
	return target{
		minX: minX,
		maxX: maxX,
		minY: minY,
		maxY: maxY,
	}

}

func HitTarget(c coord, t target) bool {
	return c.x >= t.minX &&
		c.x <= t.maxX &&
		c.y >= t.minY &&
		c.y <= t.maxY
}

func PassedTarget(c coord, t target) bool {
	return c.x > t.maxX || c.y < t.minY
}

func WillHit(a aim, t target) (bool, []coord) {
	coords := make([]coord, 0)
	c := coord{x: 0, y: 0}
	coords = append(coords, c)
	traj := trajectory{dx: a.x, dy: a.y}
	for !PassedTarget(c, t) {
		next := coord{
			x: c.x + traj.dx,
			y: c.y + traj.dy,
		}
		c = next
		coords = append(coords, c)
		if HitTarget(next, t) {
			return true, coords
		}
		newDx := traj.dx
		if newDx > 0 {
			newDx--
		} else if newDx < 0 {
			newDx++
		}
		newDy := traj.dy - 1
		traj = trajectory{dx: newDx, dy: newDy}
	}
	return false, coords
}

func Triangular(max int) int {
	i := 1
	s := 1
	for s < max {
		i++
		s += i
	}
	return i
}

func FindApogee(path []coord) int {
	max := math.MinInt
	for _, c := range path {
		if c.y > max {
			max = c.y
		}
	}
	return max
}

func main() {
	inputPath := os.Args[1]
	lines := utils.AsInputList(utils.ReadInput(inputPath))
	raw := lines[0]
	target := GetTarget(raw)
	minXAim := Triangular(target.minX)
	maxXAim := target.maxX
	minYAim := target.minY
	maxYAim, _ := strconv.Atoi(os.Args[2])
	hits := make(map[aim][]coord)
	for x := minXAim; x <= maxXAim; x++ {
		for y := minYAim; y < maxYAim; y++ {
			a := aim{x: x, y: y}
			willHit, coords := WillHit(a, target)
			if willHit {
				hits[a] = coords
			}
		}
	}
	max := math.MinInt
	for _, path := range hits {
		local := FindApogee(path)
		if local > max {
			max = local
		}
	}
	fmt.Println(max)
	fmt.Println(len(hits))
}
