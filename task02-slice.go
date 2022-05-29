package homework

func reverse(input []int64) (result []int64) {
	idx := len(input) - 1
	for i := 0; i < len(input)/2; i++ {
		input[i], input[idx] = input[idx], input[i]
		idx = idx - 1
	}

	return input
}
