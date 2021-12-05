package lib

import (
	"bufio"
	"os"
	"strconv"
)

func ParseFile(path string) []int {
	f, err := os.Open("depths.txt")
	if err != nil {
		panic(err)
	}

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
