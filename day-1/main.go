package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	depths := ParseFile("depths.txt")

	fmt.Printf("Part 1: %d\n", calculate_increases(depths))
	fmt.Printf("Part 2: %d\n", calculate_window_increases(depths))
}

func ParseFile(path string) []int {
	f, err := os.Open("depths.txt")
	check(err)

	scanner := bufio.NewScanner(f)
	var depths []int
	for scanner.Scan() {
		x, err := strconv.Atoi(scanner.Text())
		if err == nil {
			depths = append(depths, x)
		}
	}

	return depths
}

func calculate_increases(depths []int) int {
	increases := 0

	for i, depth := range depths {
		if i == 0 {
			continue
		}

		if depth > depths[i-1] {
			increases += 1
		}
	}

	return increases
}

func calculate_window_increases(depths []int) int {
	windows := make([][]int, 0)
	for i := range depths {
		if i+3 > len(depths) {
			continue
		}
		windows = append(windows, depths[i:i+3])
	}

	increases := 0
	for i, window := range windows {
		if i == 0 {
			continue
		}

		curr_sum := Sum(window)
		prev_sum := Sum(windows[i-1])

		if curr_sum > prev_sum {
			increases += 1
		}
	}

	return increases
}

func Sum(to_sum []int) int {
	sum := 0
	for _, v := range to_sum {
		sum += v
	}

	return sum
}
