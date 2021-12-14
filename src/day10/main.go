package main

import (
	"adventofcode/utils"
	"fmt"
	"os"
	"sort"
	"strings"
)

func FilterLines(lines []string) ([]string, int, []int) {
	filtered := make([]string, 0)
	completions := make([]int, 0)
	score := 0
	for i := 0; i < len(lines); i++ {
		line := lines[i]
		valid, cost, completion := IsValid(line)
		if valid {
			fmt.Println(line, valid, cost, completion)
			filtered = append(filtered, line)
			completions = append(completions, completion)
		} else {
			score += cost
		}
	}
	return filtered, score, completions
}

var points = map[string]int{
	")": 3,
	"]": 57,
	"}": 1197,
	">": 25137,
}

func IsValid(chunk string) (bool, int, int) {
	stack := make([]string, 0)
	for i := 0; i < len(chunk); i++ {
		char := chunk[i : i+1]
		if IsOpen(char) {
			stack = append(stack, char)
		} else if IsClose(char) {
			if len(stack) == 0 {
				return false, -1, CompletionScore(stack)
			}
			if Closes(char, stack[len(stack)-1]) {
				stack = stack[:len(stack)-1]
			} else {
				return false, points[char], CompletionScore(stack)
			}
		}
	}
	return true, len(stack), CompletionScore(stack)
}

var completionPoints = map[string]int{
	"(": 1,
	"[": 2,
	"{": 3,
	"<": 4,
}

func CompletionScore(left []string) int {
	score := 0
	for i := len(left) - 1; i >= 0; i-- {
		score *= 5
		score += completionPoints[left[i]]
	}
	fmt.Println(len(left), left, score)
	return score
}

// lol naming
func IsOpen(char string) bool {
	return strings.EqualFold(char, "[") ||
		strings.EqualFold(char, "(") ||
		strings.EqualFold(char, "{") ||
		strings.EqualFold(char, "<")
}

func IsClose(char string) bool {
	return strings.EqualFold(char, "]") ||
		strings.EqualFold(char, ")") ||
		strings.EqualFold(char, "}") ||
		strings.EqualFold(char, ">")
}

var openers = map[string]string{
	"]": "[",
	")": "(",
	"}": "{",
	">": "<",
}

func Closes(char string, opener string) bool {
	if IsOpen(char) || !IsClose(char) {
		return false
	}
	return strings.EqualFold(opener, openers[char])
}

func main() {
	inputPath := os.Args[1]
	lines := utils.AsInputList(utils.ReadInput(inputPath))
	_, score, completion := FilterLines(lines)
	sort.Ints(completion)
	fmt.Println(score, completion)
	fmt.Println(completion[len(completion)/2])
}
