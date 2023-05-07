package lcw335

/**
 * @Description https://leetcode.cn/contest/weekly-contest-335/problems/number-of-ways-to-earn-points/
 * idea: 分组背包模板题
 **/

func waysToReachTarget(target int, types [][]int) int {
	var n = len(types)
	var M = int(1e9) + 7
	var dp = make([]int, target+1)
	dp[0] = 1
	for i := 0; i < n; i++ {
		for j := target; j >= types[i][1]; j-- {
			for k := 1; k <= types[i][0]; k++ {
				if j < k*types[i][1] {
					break
				}
				dp[j] = (dp[j] + dp[j-k*types[i][1]]) % M
			}
		}
	}
	return dp[target]
}
