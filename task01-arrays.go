package homework

func average(input [15]float32) (result float32) {
	var avg float32 = 0
	var inputLen int = len(input)

	for i := 0; i < inputLen; i++ {
		avg = avg + input[i]
	}

	return avg / float32(inputLen)
}
