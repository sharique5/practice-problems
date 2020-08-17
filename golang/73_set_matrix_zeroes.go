func setZeroes(matrix [][]int) {
	row0, col0 := false, false
	m := len(matrix)
	n := len(matrix[0])
	for i := 0; i < m; i++ {
		if matrix[i][0] == 0 {
			col0 = true
		}
	}
	for i := 0; i < n; i++ {
		if matrix[0][i] == 0 {
			row0 = true
		}
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if matrix[i][j] == 0 {
				matrix[i][0] = 0
				matrix[0][j] = 0
			}
		}
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if matrix[i][0] == 0 || matrix[0][j] == 0 {
				matrix[i][j] = 0
			}
		}
	}
	if row0 == true {
		for i := 0; i < n; i++ {
			matrix[0][i] = 0
		}
	}
	if col0 == true {
		for i := 0; i < m; i++ {
			matrix[i][0] = 0
		}
	}
}