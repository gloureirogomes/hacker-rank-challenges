package diagonaldifference

import "math"

func DiagonalDifference() int32 {
	input := [][]int32{
		{11, 2, 4},
		{4, 5, 6},
		{10, 8, -12},
	}

	leftDiagonal := int32(0)
	rightDiagonal := int32(0)
	for i := 0; i < len(input); i++ {
		leftDiagonal += input[i][i]
		rightDiagonal += input[i][len(input)-1-i]
	}

	result := float64(leftDiagonal - rightDiagonal)
	return int32(math.Abs(result))
}
