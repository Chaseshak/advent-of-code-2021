package main

import (
	"fmt"
	"strconv"

	"github.com/Chaseshak/advent-of-code-2021/lib"
)

func main() {
	inputs := lib.ParseFile("input.txt")
	part1(inputs)
	fmt.Println("-----")
	part2(inputs)
}

func part1(inputs []string) {
	gamma_bits, epsilon_bits := "", ""
	ones_bits := onesBits(inputs)

	num_binary_nums := len(inputs)
	for _, num_ones := range ones_bits {
		if num_ones > (num_binary_nums / 2) {
			gamma_bits += "1"
			epsilon_bits += "0"
		} else {
			gamma_bits += "0"
			epsilon_bits += "1"
		}
	}

	gamma_int, err := strconv.ParseInt(gamma_bits, 2, 64)
	lib.Check(err)
	epsilon_int, err := strconv.ParseInt(epsilon_bits, 2, 64)
	lib.Check(err)
	power_consumption := gamma_int * epsilon_int

	fmt.Printf("-----\nPart 1\nGamma Bits: %v\nGamma Int: %d\nEpsilon Bits: %v\nEpsilon Int: %d\nPower Consumption: %d\n", gamma_bits, gamma_int, epsilon_bits, epsilon_int, power_consumption)
}

func onesBits(inputs []string) []int {
	ones_bits := make([]int, len(inputs[0]))

	for _, bits := range inputs {
		for i, bit := range bits {
			if bit == '1' {
				ones_bits[i] += 1
			}
		}
	}

	return ones_bits
}

func part2(inputs []string) {
	fmt.Println(calculateCo2ScrubberRating(inputs, 0))
}

func calculateCo2ScrubberRating(values []string, index int) int {
	if len(values) == 1 {
		rating, _ := strconv.Atoi(values[0])
		return rating
	}

	grouped_bits := [][]string{
		{}, {}, {}, {}, {}, {}, {}, {}, {}, {}, {}, {},
	}
	for _, binary_num := range values {
		for y, bit := range binary_num {
			grouped_bits[y] = append(grouped_bits[y], string(bit))
		}
	}
	zeroes, ones := countBits(grouped_bits[index])

	key_value := "0"
	if zeroes == ones || zeroes < ones {
		key_value = "1"
	}

	var next_values []string
	for _, binary_num := range values {
		if string(binary_num[index]) == key_value {
			next_values = append(next_values, binary_num)
		}
	}

	return calculateCo2ScrubberRating(next_values, index+1)
}

func calculateOxygenScrubberRating(values []string, index int) int {
	if len(values) == 1 {
		rating, _ := strconv.Atoi(values[0])
		return rating
	}

	grouped_bits := [][]string{
		{}, {}, {}, {}, {}, {}, {}, {}, {}, {}, {}, {},
	}
	for _, binary_num := range values {
		for y, bit := range binary_num {
			grouped_bits[y] = append(grouped_bits[y], string(bit))
		}
	}
	zeroes, ones := countBits(grouped_bits[index])

	key_value := "1"
	if zeroes == ones || zeroes < ones {
		key_value = "1"
	}

	var next_values []string
	for _, binary_num := range values {
		if string(binary_num[index]) == key_value {
			next_values = append(next_values, binary_num)
		}
	}

	return calculateCo2ScrubberRating(next_values, index+1)
}

func countBits(bits []string) (int, int) {
	zeroes, ones := 0, 0

	for _, bit := range bits {
		if bit == "0" {
			zeroes += 1
		}
		if bit == "1" {
			ones += 1
		}
	}
	return zeroes, ones
}
