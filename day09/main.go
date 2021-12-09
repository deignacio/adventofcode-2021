package main

import (
	"adventofcode/utils"
	"fmt"
	"image"
	"image/png"
	"os"
	"sort"
	"strconv"
	"strings"
)

func ParseInputs(inputs []string) [][]int {
	values := make([][]int, 0)
	for i := 0; i < len(inputs); i++ {
		input := strings.Split(inputs[i], "")
		value := make([]int, 0)
		for j := 0; j < len(input); j++ {
			val, _ := strconv.Atoi(input[j])
			if val == 9 {
				value = append(value, -1)
			} else {
				value = append(value, 0)
			}
		}
		values = append(values, value)
	}
	return values
}

func PaintBucket(depths [][]int, x int, y int, color int) int {
	if x < 0 {
		return 0
	}
	if x >= len(depths) {
		return 0
	}
	if y < 0 {
		return 0
	}
	if y >= len(depths[x]) {
		return 0
	}
	if depths[x][y] != 0 {
		return 0
	}

	depths[x][y] = color

	return 1 +
		PaintUp(depths, x - 1, y, color) +
		PaintLeft(depths, x, y - 1, color) +
		PaintRight(depths, x, y + 1, color) +
		PaintDown(depths, x + 1, y, color)
}

func PaintUp(depths [][]int, x int, y int, color int) int {
	if x < 0 {
		return 0
	}
	if x >= len(depths) {
		return 0
	}
	if y < 0 {
		return 0
	}
	if y > len(depths[x]) {
		return 0
	}
	if depths[x][y] != 0 {
		return 0
	}

	depths[x][y] = color

	return 1 +
		PaintUp(depths, x - 1, y, color) +
		PaintLeft(depths, x, y - 1, color) +
		PaintRight(depths, x, y + 1, color)
}

func PaintLeft(depths [][]int, x int, y int, color int) int {
	if x < 0 {
		return 0
	}
	if x >= len(depths) {
		return 0
	}
	if y < 0 {
		return 0
	}
	if y > len(depths[x]) {
		return 0
	}
	if depths[x][y] != 0 {
		return 0
	}

	depths[x][y] = color

	return 1 +
		PaintUp(depths, x - 1, y, color) +
		PaintLeft(depths, x, y - 1, color) +
		PaintDown(depths, x + 1, y, color)
}

func PaintRight(depths [][]int, x int, y int, color int) int {
	if x < 0 {
		return 0
	}
	if x >= len(depths) {
		return 0
	}
	if y < 0 {
		return 0
	}
	if y >= len(depths[x]) {
		return 0
	}
	if depths[x][y] != 0 {
		return 0
	}

	depths[x][y] = color

	return 1 +
		PaintUp(depths, x - 1, y, color) +
		PaintRight(depths, x, y + 1, color) +
		PaintDown(depths, x + 1, y, color)
}

func PaintDown(depths [][]int, x int, y int, color int) int {
	if x < 0 {
		return 0
	}
	if x >= len(depths) {
		return 0
	}
	if y < 0 {
		return 0
	}
	if y >= len(depths[x]) {
		return 0
	}
	if depths[x][y] != 0 {
		return 0
	}

	depths[x][y] = color

	return 1 +
		PaintLeft(depths, x, y - 1, color) +
		PaintRight(depths, x, y + 1, color) +
		PaintDown(depths, x + 1, y, color)
}

func main() {
	inputs := utils.AsInputList(utils.ReadInput("/Users/dignacio/Documents/adventofcode/day09/input_part_1"))
	depths := ParseInputs(inputs)
	color := 1
	basins := make([]int, 0)
	for i := 0; i < len(depths); i++ {
		row := depths[i]
		for j := 0; j < len(row); j++ {
			if depths[i][j] != 0 {
				continue
			}
			size := PaintBucket(depths, i, j, color)
			basins = append(basins, size)
			color++
		}
	}
	sort.Ints(basins)
	fmt.Println(basins)
	basinCount := len(basins)
	fmt.Println(basins[basinCount - 1] * basins[basinCount - 2] * basins[basinCount - 3])
	img := image.NewRGBA(image.Rect(0, 0, len(depths), len(depths[0])))
	i := 0
	for x := 0; x < len(depths); x++ {
		for y := 0; y < len(depths[0]); y++ {
			if depths[x][y] == -1 {
				img.Pix[i] = 0
				img.Pix[i+1] = 0
				img.Pix[i+2] = 0
				img.Pix[i+3] = 255
			} else {
				depth := uint8(depths[x][y])
				img.Pix[i] = depth
				img.Pix[i+1] = depth
				img.Pix[i+2] = depth
				img.Pix[i+3] = 255
			}
			i += 4
		}
	}
	output, _ := os.Create("/Users/dignacio/Documents/adventofcode/day09/input_part_1.png")
	png.Encode(output, img)
	output.Close()
}
