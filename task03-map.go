package homework

import (
	"sort"
)

func sortMapValues(input map[int]string) (result []string) {
	var length = len(input)
	var keys = make([]int, 0, length)

	for k := range input {
		keys = append(keys, k)
	}

	sort.Ints(keys)

	result = make([]string, 0, length)

	for _, k := range keys {
		result = append(result, input[k])
	}

	return result
}
