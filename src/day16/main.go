package main

import (
	"adventofcode/utils"
	"fmt"
	"math"
	"os"
)

type packet struct {
	version    int
	typeId     int
	lengthId   int
	literal    int
	length     int
	subPackets []packet
}

// Don't judge me
var toBits = map[string][]int{
	"0": {0, 0, 0, 0},
	"1": {0, 0, 0, 1},
	"2": {0, 0, 1, 0},
	"3": {0, 0, 1, 1},
	"4": {0, 1, 0, 0},
	"5": {0, 1, 0, 1},
	"6": {0, 1, 1, 0},
	"7": {0, 1, 1, 1},
	"8": {1, 0, 0, 0},
	"9": {1, 0, 0, 1},
	"A": {1, 0, 1, 0},
	"B": {1, 0, 1, 1},
	"C": {1, 1, 0, 0},
	"D": {1, 1, 0, 1},
	"E": {1, 1, 1, 0},
	"F": {1, 1, 1, 1},
}

func ToInt(bits []int) int {
	size := len(bits)
	sum := 0
	for i := 0; i < size; i++ {
		sum += bits[i] * int(math.Exp2(float64(size-1-i)))
	}
	return sum
}

func BuildInput(raw string) []int {
	input := make([]int, 0)
	for _, c := range raw {
		input = append(input, toBits[string(c)]...)
	}
	return input
}

func BuildPacket(input []int) (packet, int) {
	version := ToInt(input[0:3])
	typeId := ToInt(input[3:6])
	consumed := 6
	if typeId == 4 {
		literal, literalLength := ParseLiteral(input[6:])
		return packet{version: version, typeId: typeId, literal: literal}, consumed + literalLength
	}
	lengthId := input[6]
	length := 0
	subPackets := make([]packet, 0)
	if lengthId == 0 {
		length = ToInt(input[7:22])
		consumed = 22
		subPacketConsumed := 0
		for subPacketConsumed < length {
			sub, subC := BuildPacket(input[consumed+subPacketConsumed:])
			subPackets = append(subPackets, sub)
			subPacketConsumed += subC
		}
		consumed += subPacketConsumed
	} else {
		length = ToInt(input[7:18])
		consumed = 18
		subPacketConsumed := 0
		for i := 0; i < length; i++ {
			sub, subC := BuildPacket(input[consumed+subPacketConsumed:])
			subPackets = append(subPackets, sub)
			subPacketConsumed += subC
		}
		consumed += subPacketConsumed
	}
	return packet{version: version, typeId: typeId, lengthId: lengthId, length: length, subPackets: subPackets}, consumed
}

func ParseLiteral(value []int) (int, int) {
	literal := make([]int, 0)
	i := 0
	for {
		head := value[i]
		literal = append(literal, value[i+1], value[i+2], value[i+3], value[i+4])
		i += 5
		if head == 0 || i >= len(value) {
			break
		}
	}
	return ToInt(literal), i
}

func Evaluate(root packet) int {
	if root.typeId == 0 {
		val := 0
		for _, p := range root.subPackets {
			val += Evaluate(p)
		}
		return val
	} else if root.typeId == 1 {
		val := 1
		for _, p := range root.subPackets {
			val *= Evaluate(p)
		}
		return val
	} else if root.typeId == 2 {
		val := math.MaxInt
		for _, p := range root.subPackets {
			v := Evaluate(p)
			if v < val {
				val = v
			}
		}
		return val
	} else if root.typeId == 3 {
		val := math.MinInt
		for _, p := range root.subPackets {
			v := Evaluate(p)
			if v > val {
				val = v
			}
		}
		return val
	} else if root.typeId == 4 {
		return root.literal
	} else if root.typeId == 5 {
		first := Evaluate(root.subPackets[0])
		second := Evaluate(root.subPackets[1])
		if first > second {
			return 1
		}
		return 0
	} else if root.typeId == 6 {
		first := Evaluate(root.subPackets[0])
		second := Evaluate(root.subPackets[1])
		if first < second {
			return 1
		}
		return 0
	} else if root.typeId == 7 {
		first := Evaluate(root.subPackets[0])
		second := Evaluate(root.subPackets[1])
		if first == second {
			return 1
		}
		return 0
	}
	return 0
}

func SumPackets(root packet) int {
	sum := 0
	for _, p := range root.subPackets {
		sum += Evaluate(p)
	}
	return sum
}

func main() {
	inputPath := os.Args[1]
	lines := utils.AsInputList(utils.ReadInput(inputPath))
	raw := lines[0]
	input := BuildInput(raw)
	root, consumed := BuildPacket(input)
	fmt.Println(root, consumed)
	val := Evaluate(root)
	fmt.Println(val)

}
