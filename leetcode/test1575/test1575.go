package test1575

/**
 * @Description
 * idea: dp[i][j] 表示到达城市i，花费汽油为j的方案数
   		dp[i][j] = dp[0..i-1][j-gas(i,j)] + dp[i+1,n][j-gas(i,j)]
		dp[start][0] = 1
 **/

const M = int(1e9) + 7

func countRoutes(locations []int, start int, finish int, fuel int) int {
	//if start > finish { start, finish = finish, start }
	var n = len(locations)
	var dp = make([][]int, n)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, fuel+1)
	}
	dp[start][0] = 1
	for j := 1; j <= fuel; j++ {
		for i := 0; i < n; i++ {
			for k := 0; k < n; k++ {
				if k == i {
					continue
				}
				var cost = abs(locations[i] - locations[k])
				if cost > j {
					continue
				}
				dp[i][j] = (dp[i][j] + dp[k][j-cost]) % M
			}
		}
	}
	var res = 0
	for i := 0; i <= fuel; i++ {
		res = (res + dp[finish][i]) % M
	}
	return res
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
