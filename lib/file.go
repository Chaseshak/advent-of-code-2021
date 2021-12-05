package lib

import (
	"bufio"
	"os"
)

func ParseFile(path string) []string {
	f, err := os.Open(path)
	Check(err)

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
