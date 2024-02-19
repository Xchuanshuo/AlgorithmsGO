package test1690

/**
 * @Description https://leetcode.cn/problems/stone-game-vii/
 * idea: 区间dp，对于A和B而言 都是最大化 当前 - 对方 (对A,最大化A-B;对,最小化A-B，即最大化B-A)
 **/

func stoneGameVII(stones []int) int {
	var n = len(stones)
	var sum = make([]int, n+1)
	var dp = make([][]int, n)
	for i := 0; i < n; i++ {
		sum[i+1] = sum[i] + stones[i]
		dp[i] = make([]int, n)
	}
	for l := 2; l <= n; l++ {
		for i := 0;i <= n-l;i++ {
			var j = i + l - 1
			if l == 2 {
				dp[i][j] = max(stones[i], stones[j])
			} else {
				dp[i][j] = max(sum[j+1] - sum[i+1] - dp[i+1][j], sum[j] - sum[i] - dp[i][j-1])
			}
		}
	}
	return dp[0][n-1]
}
