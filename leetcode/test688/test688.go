package test688

var dirs = [][]int{{1, 2}, {-1, 2}, {1, -2}, {-1, -2}, {2, 1}, {2, -1}, {-2, 1}, {-2, -1}}

func knightProbability(n int, K int, row int, column int) float64 {
	var dp = make([][][]float64, n)
	for i := 0; i < n; i++ {
		dp[i] = make([][]float64, n)
		for j := 0; j < n; j++ {
			dp[i][j] = make([]float64, K+1)
			dp[i][j][0] = 1.0
		}
	}
	dp[row][column][0] = 1.0
	for k := 1; k <= K; k++ {
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				for d := 0; d < len(dirs); d++ {
					nx, ny := i+dirs[d][0], j+dirs[d][1]
					if nx < 0 || nx >= n || ny < 0 || ny >= n {
						continue
					}
					dp[nx][ny][k] += dp[i][j][k-1] * 0.125
				}
			}
		}
	}
	return dp[row][column][K]
}
