package test807

func maxIncreaseKeepingSkyline(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	rows, cols := make([]int, m), make([]int, n)
	for i := 0;i < m;i++ {
		for j := 0; j < n; j++{
		    cols[j] = max(grid[i][j], cols[j])
			rows[i] = max(grid[i][j], rows[i])
		}
	}
	res := 0
	for i := 0;i < m;i++ {
		for j := 0; j < n; j++{
			res += min(rows[i], cols[j]) - grid[i][j]
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
	if a > b {
		return b
	}
	return a
}


