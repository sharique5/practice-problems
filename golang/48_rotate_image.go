func rotate(matrix [][]int) {
	rows := len(matrix)
	col := len(matrix[0])
	for i := 0; i < rows; i++ {
		for j := i; j < col; j++ {
			temp := matrix[i][j]
			matrix[i][j] = matrix[j][i]
			matrix[j][i] = temp
		}
	}
	for i := 0; i < rows; i++ {
		for j := 0; j < col/2; j++ {
			temp := matrix[i][j]
			matrix[i][j] = matrix[i][col-j-1]
			matrix[i][col-j-1] = temp
		}
	}
}