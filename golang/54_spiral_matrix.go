func spiralOrder(matrix [][]int) []int {
	var res []int
	m := len(matrix)
	if m == 0 {
		return res
	}
	n := len(matrix[0])
	k, l := 0, 0
	for k < m && l < n {
		for i := l; i < n; i++ {
			res = append(res, matrix[k][i])
		}
		k++
		for i := k; i < m; i++ {
			res = append(res, matrix[i][n-1])
		}
		n--
		if k < m {
			for i := n - 1; i > l-1; i-- {
				res = append(res, matrix[m-1][i])
			}
			m--
		}
		if l < n {
			for i := m - 1; i > k-1; i-- {
				res = append(res, matrix[i][l])
			}
			l++
		}
	}
	return res
}