package lcw340

/**
 * @Description https://leetcode.cn/problems/minimum-number-of-visited-cells-in-a-grid/
 * idea: bfs + 行列最大值优化 时间复杂度O(m*n*n)
 **/

func minimumVisitedCells1(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	// 记录每行最大的列 和 每列最大的行
	var rows = make([]int, m)
	var cols = make([]int, n)
	var q = [][]int{{0, 0}}
	var step = 0
	for len(q) > 0 {
		var sz = len(q)
		step++
		for i := 0; i < sz; i++ {
			x, y := q[0][0], q[0][1]
			q = q[1:]
			if x == m-1 && y == n-1 {
				return step
			}
			for nx := max(x+1, cols[y]); nx < m && nx <= x+grid[x][y]; nx++ {
				if nx == m-1 && y == n-1 {
					return step + 1
				}
				q = append(q, []int{nx, y})
			}
			for ny := max(y+1, rows[x]); ny < n && ny <= y+grid[x][y]; ny++ {
				if x == m-1 && ny == n-1 {
					return step + 1
				}
				q = append(q, []int{x, ny})
			}
			cols[y] = max(cols[y], x+grid[x][y])
			rows[x] = max(rows[x], y+grid[x][y])
		}
	}
	return -1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
