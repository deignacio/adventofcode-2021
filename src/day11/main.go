package main

import (
	"adventofcode/utils"
	"fmt"
	"image"
	"image/color"
	"image/color/palette"
	"image/gif"
	"os"
	"strconv"
)

func FilterLines(lines []string) [][]int {
	octopuses := make([][]int, 0)
	for i := 0; i < len(lines); i++ {
		line := lines[i]
		row := make([]int, len(line))
		for j := 0; j < len(line); j++ {
			value, _ := strconv.Atoi(line[j : j+1])
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

func Dump(octopuses [][]int) *image.Paletted {
	scale := 25
	w := len(octopuses) * scale
	h := len(octopuses[0]) * scale
	img := image.NewPaletted(image.Rect(0, 0, w, h), palette.Plan9)
	for x := 0; x < len(octopuses); x++ {
		for y := 0; y < len(octopuses[0]); y++ {
			brightness := uint8(octopuses[x][y]) * 25
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
	return img
}

func main() {
	inputPath := os.Args[1]
	lines := utils.AsInputList(utils.ReadInput(inputPath))
	octopuses := FilterLines(lines)
	total := 0
	first := -1
	size := len(octopuses) * len(octopuses[0])
	fmt.Println(size)
	fmt.Println(total, octopuses)
	images := make([]*image.Paletted, 0)
	delays := make([]int, 0)
	images = append(images, Dump(octopuses))
	delays = append(delays, 0)
	for i := 1; i < 500; i++ {
		flashes := OneStep(octopuses)
		images = append(images, Dump(octopuses))
		delays = append(delays, 0)
		total += flashes
		if flashes >= size {
			first = i
			break
		}
	}
	outputPath := fmt.Sprintf("%s/day11/octoposes.gif", os.Args[2])
	out, _ := os.Create(outputPath)
	defer out.Close()
	gif.EncodeAll(out, &gif.GIF{
		Image: images,
		Delay: delays,
	})
	fmt.Println(total, first, octopuses)
}
