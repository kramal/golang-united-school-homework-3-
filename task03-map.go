package homework

func sortMapValues(input map[int]string) (result []string) {
	var keys []int

	for k := range input {
		keys = append(keys, k)
	}

	for i := 0; i < len(input); i++ {
		for j := i; j < len(input); j++ {
			if keys[i] > keys[j] {
				keys[i], keys[j] = keys[j], keys[i]
			}
		}
	}

	for k := range keys {
		result = append(result, input[keys[k]])
	}

	return result
}
