package homework

func average(input [15]float32) (result float32) {
	result = 0;	

	if len(input) > 0 {
		for i := range input {
			result += input[i];
		}

		result = result / float32(len(input));
	}

	return result;
}
