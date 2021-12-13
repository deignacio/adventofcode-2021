package main

import (
	"adventofcode/utils"
	"fmt"
	"image"
	"image/png"
	"os"
	"strconv"
	"strings"
)

type fold struct {
	axis string
	val int
}

func FilterLines(lines []string) (map[int]map[int]bool, []fold) {
	dots := make(map[int]map[int]bool, 0)
	folds := make([]fold, 0)
	for i := 0; i < len(lines); i++ {
		line := lines[i]
		coords := strings.Split(line, ",")
		if len(coords) == 1 {
			instruction := strings.Split(line, " ")
			axisAndVal := strings.Split(instruction[2], "=")
			val, _ := strconv.Atoi(axisAndVal[1])
			folds = append(folds, fold{axis: axisAndVal[0], val: val})
			continue
		}
		x, _ := strconv.Atoi(coords[0])
		xVal, xPres := dots[x]
		if !xPres {
			xVal = make(map[int]bool, 0)
		}
		y, _ := strconv.Atoi(coords[1])
		xVal[y] = true
		dots[x] = xVal
	}

	return dots, folds
}

func ExecuteFold(dots map[int]map[int]bool, action fold, maxX int, maxY int) (map[int]map[int]bool, int, int) {
	if action.axis == "x" {
		return ExecuteXFold(dots, action), (maxX - 1) / 2, maxY
	}
	return ExecuteYFold(dots, action), maxX, (maxY - 1) / 2
}

func ExecuteXFold(dots map[int]map[int]bool, action fold) map[int]map[int]bool{
	next := make(map[int]map[int]bool, 0)
	for x, row := range dots {
		for y, _ := range row {
			if x < action.val {
				next = AddDot(next, x, y)
			} else {
				newX := 2 * action.val - x
				next = AddDot(next, newX, y)
			}
		}
	}

	return next
}

func AddDot(dots map[int]map[int]bool, x int, y int) map[int]map[int]bool {
	xVal, xPres := dots[x]
	if !xPres {
		xVal = make(map[int]bool, 0)
	}
	xVal[y] = true
	dots[x] = xVal
	return dots
}

func ExecuteYFold(dots map[int]map[int]bool, action fold) map[int]map[int]bool {
	next := make(map[int]map[int]bool, 0)
	for x, row := range dots {
		for y, _ := range row {
			if y < action.val {
				next = AddDot(next, x, y)
			} else {
				newY := 2 * action.val - y
				next = AddDot(next, x, newY)
			}
		}
	}

	return next
}

func CountLeft(dots map[int]map[int]bool, maxX int, maxY int) int {
	visible := 0
	for x := 0; x <= maxX; x++ {
		row, xPresent := dots[x]
		if !xPresent {
			continue
		}
		for y := 0; y <= maxY; y++ {
			_, yPresent := row[y]
			if yPresent {
				visible ++
			}
		}
	}
	return visible
}

func Dump(dots map[int]map[int]bool, maxX int, maxY int) {
	img := image.NewRGBA(image.Rect(0, 0, maxX+1, maxY+1))
	i := 0
		for y := 0; y <= maxY; y++ {
			for x := 0; x <= maxX; x++ {
			if !dots[x][y] {
				img.Pix[i] = 0
				img.Pix[i+1] = 0
				img.Pix[i+2] = 0
				img.Pix[i+3] = 255
			} else {
				depth := uint8(128)
				img.Pix[i] = depth
				img.Pix[i+1] = depth
				img.Pix[i+2] = depth
				img.Pix[i+3] = 255
			}
			i += 4
		}
	}
	output, _ := os.Create("/Users/dignacio/Documents/adventofcode/day13/input_part_1.png")
	png.Encode(output, img)
	output.Close()
}

func main() {
	lines := utils.AsInputList(utils.ReadInput("/Users/dignacio/Documents/adventofcode/day13/input_part_1"))
	dots, folds := FilterLines(lines)
	maxX := 0
	maxY := 0
	for x, row := range dots {
		if x > maxX {
			maxX = x
		}
		for y, _ := range row {
			if y > maxY {
				maxY = y
			}
		}
	}
	for _, fold := range folds {
		dots, maxX, maxY = ExecuteFold(dots, fold, maxX, maxY)
		//visible := CountLeft(dots, maxX, maxY)
	}
	fmt.Println(maxX, maxY, dots)
	Dump(dots, maxX, maxY)
}
