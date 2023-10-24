package verybigsum

func VeryBigSum(ar []int64) int64 {
	result := int64(0)
	for _, value := range ar {
		result += value
	}
	return result
}
