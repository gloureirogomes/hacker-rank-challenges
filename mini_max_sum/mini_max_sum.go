package minimaxsum

import (
	"fmt"
	"sort"
)

func MiniMaxSum() {
	input := []int32{1, 2, 3, 4, 5}

	sort.Slice(input, func(i, j int) bool {
		return input[i] < input[j]
	})

	var minSum, maxSum int64
	for i := 0; i < len(input)-1; i++ {
		minSum += int64(input[i])
		maxSum += int64(input[i+1])
	}

	fmt.Printf("%d %d\n", minSum, maxSum)
}
