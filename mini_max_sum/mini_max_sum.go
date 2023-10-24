package minimaxsum

import (
	"fmt"
	"sort"
)

func MiniMaxSum(arr []int32) {
	sort.Slice(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	})

	var minSum, maxSum int64
	for i := 0; i < len(arr)-1; i++ {
		minSum += int64(arr[i])
		maxSum += int64(arr[i+1])
	}

	fmt.Printf("%d %d\n", minSum, maxSum)
}
