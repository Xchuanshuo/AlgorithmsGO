package test913


var (
	n int
	g [][]int
	dp [][][]int
)

func catMouseGame(graph [][]int) int {
	n = len(graph)
	g = graph
	dp = make([][][]int, 2*n+1)
	for i := 0;i < 2*n+1;i++ {
		dp[i] = make([][]int, n)
		for j := 0;j < n;j++ {
			dp[i][j] = make([]int, n)
			for k := 0; k < n; k++{
				dp[i][j][k] = -1
			}
		}
	}
	return dfs(0, 1, 2)
}

func dfs(k, m, c int) int {
	if c == m {
		dp[k][m][c] = 2
		return 2
	} else if m == 0 {
		dp[k][m][c] = 1
		return 1
	} else if k >= 2*n {
		dp[k][m][c] = 0
		return 0
	}
	if dp[k][m][c] != -1 { return dp[k][m][c] }
	if k%2 == 0 {
		isCatWin := true
		for _, next := range g[m] {
			res := dfs(k + 1, next, c)
			if res == 1 {
				dp[k][m][c] = 1
				return 1
			}
			if res != 2 { isCatWin = false }
		}
		if isCatWin {
			dp[k][m][c] = 2
			return dp[k][m][c]
		}
		dp[k][m][c] = 0
		return dp[k][m][c]
	} else {
		isMouseWin := true
		for _, next := range g[c] {
			if next == 0 { continue }
			res := dfs(k + 1, m, next)
			if res == 2 {
				dp[k][m][c] = 2
				return 2
			}
			if res != 1 { isMouseWin = false }
		}
		if isMouseWin {
			dp[k][m][c] = 1
			return 1
		}
		dp[k][m][c] = 0
		return dp[k][m][c]
	}
}