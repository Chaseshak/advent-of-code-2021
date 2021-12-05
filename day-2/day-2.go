package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/Chaseshak/advent-of-code-2021/lib"
)

func main() {
	inputs := lib.ParseFile("input.txt")
	part1(inputs)
	part2(inputs)
}

func part1(inputs []string) {
	horizontal_pos, depth := 0, 0
	for _, input := range inputs {
		input := strings.Fields(input)
		instruction := input[0]
		x := input[1]

		num, err := strconv.Atoi(x)
		lib.Check(err)

		switch instruction {
		case "forward":
			horizontal_pos += num
		case "down":
			depth += num
		case "up":
			depth -= num
		}
	}

	product := depth * horizontal_pos
	fmt.Printf("-----\nPart 1\nHorizontal Position: %d\nDepth: %d\nDepth * Horizontal Position = %d\n", horizontal_pos, depth, product)
}

func part2(inputs []string) {
	horizontal_pos, depth, aim := 0, 0, 0
	for _, input := range inputs {
		input := strings.Fields(input)
		instruction := input[0]
		x := input[1]

		num, err := strconv.Atoi(x)
		lib.Check(err)

		switch instruction {
		case "forward":
			horizontal_pos += num
			depth += aim * num
		case "down":
			aim += num
		case "up":
			aim -= num
		}
	}

	product := depth * horizontal_pos
	fmt.Printf("-----\nPart 1\nHorizontal Position: %d\nDepth: %d\nAim: %d\nDepth * Horizontal Position = %d\n", horizontal_pos, depth, aim, product)
}
