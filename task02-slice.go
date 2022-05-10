package homework

func reverse(input []int64) (result []int64) {
	var length = len(input)

	result = make([]int64, length)

	for i := 0; i < length; i++ {
		result[length-1-i] = input[i]
	}

	return result
}
