package diagonaldifference

import "math"

func DiagonalDifference(arr [][]int32) int32 {
	leftDiagonal := int32(0)
	rightDiagonal := int32(0)
	for i := 0; i < len(arr); i++ {
		leftDiagonal += arr[i][i]
		rightDiagonal += arr[i][len(arr)-1-i]
	}
	result := float64(leftDiagonal - rightDiagonal)
	return int32(math.Abs(result))
}
