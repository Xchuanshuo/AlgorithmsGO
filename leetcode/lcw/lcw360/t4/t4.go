package t4

/**
 * @Description https://leetcode.cn/contest/weekly-contest-360/problems/maximize-value-of-function-in-a-ball-passing-game/
 * idea: 树上倍增 求出每个节点的第2^0, 2^1...,2^31个父节点及节点编号和
 **/

func getMaxFunctionValue(receiver []int, k int64) int64 {
	var n = len(receiver)
	var dp = make([][35]int, n)
	var sum = make([][35]int64, n)
	for i := 0; i < n; i++ {
		dp[i][0] = receiver[i]
		sum[i][0] = int64(receiver[i])
	}
	for j := 1; j < 35; j++ {
		for i := 0; i < n; i++ {
			dp[i][j] = dp[dp[i][j-1]][j-1]
			sum[i][j] = sum[i][j-1] + sum[dp[i][j-1]][j-1]
		}
	}
	var res = int64(0)
	for i := 0; i < n; i++ {
		x, s := i, int64(i)
		for j := 0; j < 35; j++ {
			if k&(1<<j) == 0 {
				continue
			}
			s += sum[x][j]
			x = dp[x][j]
		}
		res = max(res, s)
	}
	return res
}

func max(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}
