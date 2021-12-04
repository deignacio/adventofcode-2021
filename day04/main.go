package main

import (
	"adventofcode/utils"
	"fmt"
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



func main() {
	inputs := utils.AsInputList(utils.ReadInput("/Users/dignacio/Documents/adventofcode/day04/input_part_1"))
	order := GetOrder(inputs[0])
	boards := GetBoards(inputs[1:])
	winner, score := PlayBingo(order, boards)
	fmt.Println(winner, score)
}
