package lib

import "strconv"

func StringArrayToIntArray(to_convert []string) []int {
	converted := make([]int, len(to_convert))
	for i, value := range to_convert {
		x, err := strconv.Atoi(value)
		if err == nil {
			converted[i] = x
		}

	}

	return converted
}
