package test1380

func luckyNumbers(matrix [][]int) []int {
	m, n := len(matrix), len(matrix[0])
	var rows = make([]int, m)
	for i := 0; i < m; i++ {
		rows[i] = 1e9
	}
	var cols = make([]int, n)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			cols[j] = max(cols[j], matrix[i][j])
			rows[i] = min(rows[i], matrix[i][j])
		}
	}
	var res = make([]int, 0)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if rows[i] == matrix[i][j] && cols[j] == matrix[i][j] {
				res = append(res, matrix[i][j])
			}
		}
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}