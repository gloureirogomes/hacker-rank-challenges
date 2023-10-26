package simplearraysum

func SumArray() int32 {
	input := []int32{1, 2, 3, 4, 10, 11}

	var result int32 = 0
	for _, value := range input {
		result += value
	}

	return result
}
