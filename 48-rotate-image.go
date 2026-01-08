package LeetCode

func rotate(matrix [][]int) {
	matrixLen := len(matrix)
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < matrixLen/2; j++ {
			temp := matrix[i][j]
			matrix[i][j] = matrix[i][matrixLen-j-1]
			matrix[i][matrixLen-j-1] = temp
		}
	}

	i := 0
	j := matrixLen - 1
	count := matrixLen - 1

	for k := 0; k < matrixLen; k++ {
		for m := 1; m <= count; m++ {
			temp := matrix[i][j-m]
			matrix[i][j-m] = matrix[i+m][j]
			matrix[i+m][j] = temp
		}
		count--
		i++
		j--
	}
}
