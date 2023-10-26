package plusminus

import (
	"fmt"
)

func PlusMinus() {
	input := []int32{-4, 3, -9, 0, 4, 1}

	positive := float64(0)
	negative := float64(0)
	zero := float64(0)

	for _, value := range input {
		if value > 0 {
			positive++
			continue
		}

		if value < 0 {
			negative++
			continue
		}

		zero++
	}

	arraySize := float64(len(input))
	positiveRatio := positive / arraySize
	negativeRatio := negative / arraySize
	zeroRatio := zero / arraySize

	fmt.Printf("%.6f\n", positiveRatio)
	fmt.Printf("%.6f\n", negativeRatio)
	fmt.Printf("%.6f\n", zeroRatio)
}
