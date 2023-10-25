package birthdaycakecandles

import (
	"sort"
)

func BirthdayCakeCandles(candles []int32) int32 {
	sort.Slice(candles, func(i, j int) bool {
		return candles[i] < candles[j]
	})

	tallCandles := int32(0)
	for i := 0; i < len(candles); i++ {
		if candles[i] == candles[len(candles)-1] {
			tallCandles++
			continue
		}
	}
	return tallCandles
}
