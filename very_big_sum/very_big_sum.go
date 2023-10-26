package verybigsum

func VeryBigSum() int64 {
	input := []int64{1000000001, 1000000002, 1000000003, 1000000004, 1000000005}

	result := int64(0)
	for _, value := range input {
		result += value
	}

	return result
}
