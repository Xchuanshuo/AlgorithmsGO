package lcw362

// n个空格子与带石头的格子之间互相匹配

var INF = int(1e9) + 7

func minimumMoves(g [][]int) int {
	var ls = make([][]int, 0)
	var rs = make([][]int, 0)
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if g[i][j] == 0 {
				ls = append(ls, []int{i, j})
			} else if g[i][j] > 1 {
				for k := g[i][j]; k > 1; k-- {
					rs = append(rs, []int{i, j})
				}
			}
		}
	}
	var cal = func(i, j int) int {
		x1, y1 := ls[i][0], ls[i][1]
		x2, y2 := rs[j][0], rs[j][1]
		return abs(x1-x2) + abs(y1-y2)
	}
	var n = len(ls)
	var l = 1 << n
	var dp = make([]int, l)
	for i := 1; i < l; i++ {
		dp[i] = INF
	}
	for i := 0; i < n; i++ {
		for s := l - 1; s >= 0; s-- {
			for j := 0; j < n; j++ {
				if s&(1<<j) != 0 {
					dp[s] = min(dp[s], dp[s^(1<<j)]+cal(i, j))
				}
			}
		}
	}
	return dp[l-1]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
