package main

import (
	"adventofcode/utils"
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"
)

func GetOrder(input string) []int {
	raw := strings.Split(input, ",")
	order := make([]int, len(raw))
	for i := 0; i < len(raw); i++ {
		num, _ := strconv.Atoi(raw[i])
		order[i] = num
	}
	return order
}

func GetBoard(inputs []string) [][]int {
	r, _ := regexp.Compile("\\s*(\\d+)\\s*(\\d+)\\s*(\\d+)\\s*(\\d+)\\s*(\\d+)")
	board := make([][]int, 5)
	for i := 0; i < 5; i++ {
		input := inputs[i+1]
		match := r.FindStringSubmatch(input)
		board[i] = make([]int, 5)
		for j := 0; j < 5; j++ {
			num, _ := strconv.Atoi(match[j+1])
			board[i][j] = num
		}
	}
	return board
}

func GetBoards(inputs []string) [][][]int {
	numBoards := len(inputs) / 6
	boards := make([][][]int, 0, numBoards)
	for i := 0; i < len(inputs); i += 6 {
		boards = append(boards, GetBoard(inputs[i:i+6]))
	}
	return boards
}

func CalculateScore(board [][]int, matched [][]int, move int) int {
	score := 0
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if matched[i][j] == 0 {
				score += board[i][j]
			}
		}
	}
	return score * move
}

func DoesBoardWin(order []int, board [][]int) (int, int) {
	matched := make([][]int, 5)
	for i := 0; i < 5; i++ {
		matched[i] = make([]int, 5)
	}
	var k int
	for k = 0; k < len(order); k++ {
		num := order[k]
		for i := 0; i < 5; i++ {
			for j := 0; j < 5; j++ {
				if num == board[i][j] {
					matched[i][j] = 1
				}
			}
		}
		for i := 0; i < 5; i++ {
			if matched[0][i] == 1 &&
				matched[1][i] == 1 &&
				matched[2][i] == 1 &&
				matched[3][i] == 1 &&
				matched[4][i] == 1 {
				fmt.Println("win!", k, board, matched, num)
				return k, CalculateScore(board, matched, num)
			}
		}
		for j := 0; j < 5; j++ {
			if matched[j][0] == 1 &&
				matched[j][1] == 1 &&
				matched[j][2] == 1 &&
				matched[j][3] == 1 &&
				matched[j][4] == 1 {
				fmt.Println("win!", k, board, matched, num)
				return k, CalculateScore(board, matched, num)
			}
		}

	}
	return -1, -1
}

func PlayBingo(order []int, boards [][][]int) (int, int) {
	//fmt.Println(order)
	//fmt.Println(boards)
	moves := make([]int, len(boards))
	scores := make([]int, len(boards))
	for i := 0; i < len(boards); i++ {
		board := boards[i]
		moves[i], scores[i] = DoesBoardWin(order, board)
	}
	fmt.Println(moves, scores)
	winner := -1
	winnerScore := -1
	for i := 0; i < len(scores); i++ {
		move := moves[i]
		if move != -1 && move > winner {
			winner = move
			winnerScore = scores[i]
		}
	}
	return winner, winnerScore
}

type coord struct {
	x int
	y int
}

func ParseVent(input string) []coord {
	parsed := strings.Split(input, " ")
	fmt.Println(parsed)
	start := strings.Split(parsed[0], ",")
	startX, _ := strconv.Atoi(start[0])
	startY, _ := strconv.Atoi(start[1])
	end := strings.Split(parsed[2], ",")
	endX, _ := strconv.Atoi(end[0])
	endY, _ := strconv.Atoi(end[1])
	coords := make([]coord, 0)
	if startX != endX && startY != endY {
		xDir := 1
		yDir := 1
		if startX > endX {
			xDir = -1
		}
		if startY > endY {
			yDir = -1
		}
		currX := startX
		currY := startY
		for currX != endX && currY != endY {
			coords = append(coords, coord{x: currX, y: currY})
			currX += xDir
			currY += yDir
		}
		coords = append(coords, coord{x: currX, y: currY})
		return coords
	}
	for x := int(math.Min(float64(startX), float64(endX))); x <= int(math.Max(float64(startX), float64(endX))); x++ {
		for y := int(math.Min(float64(startY), float64(endY))); y <= int(math.Max(float64(startY), float64(endY))); y++ {
			coords = append(coords, coord{x: x, y: y})
		}
	}
	return coords
}

func GetCoords(inputs []string) []coord {
	coords := make([]coord, 0)
	for i := 0; i < len(inputs); i++ {
		input := inputs[i]
		coords = append(coords, ParseVent(input)...)
	}
	return coords
}

func BuildOcean(coords []coord) map[coord]int {
	ocean := make(map[coord]int)
	for i := 0; i < len(coords); i++ {
		_, present := ocean[coords[i]]
		if present {
			ocean[coords[i]]++
		} else {
			ocean[coords[i]] = 1
		}
	}
	return ocean
}

func main() {
	inputs := utils.AsInputList(utils.ReadInput("/Users/dignacio/Documents/adventofcode/day05/input_part_1"))
	coords := GetCoords(inputs)
	ocean := BuildOcean(coords)
	overlapped := 0
	for _, v := range ocean {
		if v > 1 {
			overlapped++
		}
	}
	//fmt.Println(coords, ocean)
	fmt.Println(overlapped)
}
