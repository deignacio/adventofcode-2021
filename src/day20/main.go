package main

import (
	"adventofcode/utils"
	"fmt"
	"image"
	"image/color"
	"image/color/palette"
	"image/gif"
	"os"
)

func ParseInputs(inputs []string, stepCount int) (map[string]bool, []bool, int, int) {
	rule := make([]bool, len(inputs[0]))
	for i, c := range inputs[0] {
		rule[i] = c == '#'
	}
	buffer := 2*stepCount + 2
	grid := make(map[string]bool)
	for y, row := range inputs[2:] {
		for x, c := range row {
			key := fmt.Sprintf("%d-%d", x+buffer, y+buffer)
			grid[key] = c == '#'
		}
	}
	return grid, rule, len(inputs[2]) + 2*buffer, len(inputs) - 2 + 2*buffer
}

func Enhance(grid map[string]bool, rule []bool, width int, height int, overflow bool) (map[string]bool, int, int) {
	next := make(map[string]bool)
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			dest := fmt.Sprintf("%d-%d", x, y)
			neighbors := GetNeighbors(grid, x, y, width, height, overflow)
			// fmt.Println(dest, neighbors)
			next[dest] = rule[neighbors]
		}
	}
	return next, width, height
}

func GetNeighbors(grid map[string]bool, x int, y int, width int, height int, overflow bool) int {
	val := 0
	for j := y - 1; j <= y+1; j++ {
		for i := x - 1; i <= x+1; i++ {
			val = val * 2
			key := fmt.Sprintf("%d-%d", i, j)
			n, pres := grid[key]
			if pres {
				if n {
					val++
				}
			} else if i < 0 || i >= width || j < 0 || j >= height {
				if overflow {
					val++
				}
			}
		}
	}
	return val
}

func Dump(grid map[string]bool, width int, height int, step int) *image.Paletted {
	scale := 10
	w := width * scale
	h := height * scale
	img := image.NewPaletted(image.Rect(0, 0, w, h), palette.Plan9)
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			key := fmt.Sprintf("%d-%d", x, y)
			brightness := uint8(0)
			if grid[key] {
				brightness = uint8(255)
			}
			for dx := 0; dx < scale; dx++ {
				for dy := 0; dy < scale; dy++ {
					img.Set(x*scale+dx, y*scale+dy, color.RGBA{
						brightness,
						brightness,
						brightness,
						255,
					})
				}
			}
		}
	}
	// outputPath := fmt.Sprintf("%s/day20/grid-%d.gif", os.Args[2], step)
	// out, _ := os.Create(outputPath)
	// defer out.Close()
	// gif.Encode(out, img, nil)
	return img
}

func CountLights(grid map[string]bool, width int, height int, buffer int) int {
	sum := 0
	for x := buffer; x < width-buffer; x++ {
		for y := buffer; y < height-buffer; y++ {
			key := fmt.Sprintf("%d-%d", x, y)
			if grid[key] {
				sum++
			}
		}
	}

	return sum
}

func main() {
	inputPath := os.Args[1]
	stepCount := 50
	lines := utils.AsInputList(utils.ReadInput(inputPath))
	grid, rule, width, height := ParseInputs(lines, stepCount)
	fmt.Println(rule)
	overflow := rule[0]
	images := make([]*image.Paletted, 0)
	delays := make([]int, 0)

	for i := 0; i < stepCount; i++ {
		images = append(images, Dump(grid, width, height, i))
		delays = append(delays, 10)
		if stepCount%2 == 0 {
			overflow = rule[0]
		} else {
			overflow = rule[511]
		}
		fmt.Println(i, CountLights(grid, width, height, i))
		grid, width, height = Enhance(grid, rule, width, height, overflow)
	}
	images = append(images, Dump(grid, width, height, stepCount))
	delays = append(delays, 10)

	outputPath := fmt.Sprintf("%s/day20/grid.gif", os.Args[2])
	out, _ := os.Create(outputPath)
	defer out.Close()
	gif.EncodeAll(out, &gif.GIF{
		Image: images,
		Delay: delays,
	})
	fmt.Println(stepCount, CountLights(grid, width, height, stepCount))
}
