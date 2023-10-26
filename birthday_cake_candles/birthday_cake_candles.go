package birthdaycakecandles

import (
	"sort"
)

func BirthdayCakeCandles() int32 {
	input := []int32{3, 2, 1, 3}

	sort.Slice(input, func(i, j int) bool {
		return input[i] < input[j]
	})

	tallCandles := int32(0)
	for i := 0; i < len(input); i++ {
		if input[i] == input[len(input)-1] {
			tallCandles++
			continue
		}
	}

	return tallCandles
}
