package main

import (
	"adventofcode/utils"
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"strings"
	"time"
)

func BuildRules(lines []string) map[byte]map[byte]byte {
	rules := make(map[byte]map[byte]byte)
	for _, rule := range lines {
		split := strings.Split(rule, " ")
		head := split[0][0]
		tail := split[0][1]
		val, present := rules[head]
		if !present {
			val = make(map[byte]byte, 0)
		}
		val[tail] = split[2][0]
		rules[head] = val
	}
	return rules
}

func ApplyRules(inStep int, outStep int, rules map[byte]map[byte]byte, outputRoot string) {
	in, _ := os.Open(fmt.Sprintf("%s/day14/output_step_%d", outputRoot, inStep))
	reader := bufio.NewReader(in)
	out, _ := os.Create(fmt.Sprintf("%s/day14/output_step_%d", outputRoot, outStep))
	writer := bufio.NewWriter(out)
	current := make([]byte, 1)
	next := make([]byte, 1)
	reader.Read(current)
	for {
		_, err := reader.Read(next)
		if err == io.EOF {
			break
		}
		writer.Write(current)
		tail := rules[current[0]]
		toInsert, present := tail[next[0]]
		if present {
			writer.WriteString(string(toInsert))
		}
		copy(current, next)
	}
	writer.Write(current)
	in.Close()
	writer.Flush()
	out.Close()
}

func ApplyRule(head byte, tail byte, depth int, max int, rules map[byte]map[byte]byte, counts map[byte]int) {
	if depth == max {
		_, present := counts[tail]
		if present {
			counts[tail]++
		} else {
			counts[tail] = 1
		}
		ComputeSize(counts)
		return
	}
	rule := rules[head]
	toInsert, present := rule[tail]
	if present {
		ApplyRule(head, toInsert, depth + 1, max, rules, counts)
		ApplyRule(toInsert, tail, depth + 1, max, rules, counts)
	}
}

func ComputeSize(counts map[byte]int) {
	sum := 0
	max := math.MinInt
	min := math.MaxInt
	for _, val := range counts {
		sum += val
		if val > max {
			max = val
		}
		if val < min {
			min = val
		}
	}
	if sum % (100 * 1024 * 1024) == 0 {
		fmt.Println(time.Now(), sum, counts, max, min, max-min)
	}
}

func CountPairs(outputRoot string, step int) map[string]int {
	count := make(map[string]int)
	in, _ := os.Open(fmt.Sprintf("%s/day14/output_step_%d", outputRoot, step))
	reader := bufio.NewReader(in)
	current := make([]byte, 1)
	for {
		_, err := reader.Read(current)
		if err == io.EOF {
			break
		}
		count[string(current)]++
	}
	in.Close()
	return count
}

func DumpStep(polymer string, step int, outputRoot string) {
	path := fmt.Sprintf("%s/day14/output_step_%d", outputRoot, step)
	output, _ := os.Create(path)
	output.WriteString(polymer)
	output.Close()
}

func main() {
	inputPath := os.Args[1]
	lines := utils.AsInputList(utils.ReadInput(inputPath))
	polymer := []byte(lines[0])
	rules := BuildRules(lines[2:])
	numSteps := 40
	current := polymer[0]
	counts := make(map[byte]int)
	counts[current] = 1
	for _, next := range polymer[1:] {
		fmt.Println(string(current), string(next), counts, time.Now())
		ApplyRule(current, next, 0, numSteps, rules, counts)
		current = next
	}
	fmt.Println(counts)
}
