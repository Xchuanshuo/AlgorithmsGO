package test741

import "math"

func cherryPickup(grid [][]int) int {
	var n = len(grid)
	var mem = make([][][]int, n)
	for i := 0; i < n; i++ {
		mem[i] = make([][]int, n)
		for j := 0; j < n; j++ {
			mem[i][j] = make([]int, n)
			for k := 0; k < n; k++ {
				mem[i][j][k] = math.MinInt32
			}
		}
	}
	var dp func(int, int, int) int
	dp = func(r1 int, c1 int, c2 int) int {
		var r2 = r1 + c1 - c2
		if r1 == n || c1 == n || r2 == n || c2 == n ||
			grid[r1][c1] == -1 || grid[r2][c2] == -1 {
			return -99999
		} else if r1 == n-1 && c1 == n-1 { // 一条路径到了终点，无需继续向下计算
			return grid[r1][c1]
		} else if mem[r1][c1][c2] != math.MinInt32 {
			return mem[r1][c1][c2]
		} else {
			var val = grid[r1][c1]
			if c1 != c2 {
				val += grid[r2][c2]
			}
			val += max(max(dp(r1+1, c1, c2), dp(r1, c1+1, c2+1)),
				max(dp(r1+1, c1, c2+1), dp(r1, c1+1, c2)))
			mem[r1][c1][c2] = val
			return val
		}
	}
	return max(0, dp(0, 0, 0))
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
