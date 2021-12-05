package lib

import (
	"bufio"
	"os"
)

func ParseFile(path string) []string {
	f, err := os.Open("depths.txt")
	if err != nil {
		panic(err)
	}

	var data []string
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}

		data = append(data, scanner.Text())
	}

	return data
}
