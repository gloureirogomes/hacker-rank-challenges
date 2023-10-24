package simplearraysum

func SumArray(ar []int32) int32 {
	var result int32 = 0
	for _, value := range ar {
		result += value
	}
	return result
}
