package lcwt112

import "sort"

/**
 * @Description https://leetcode.cn/contest/biweekly-contest-112/problems/count-k-subsequences-of-a-string-with-maximum-beauty/
 * idea: 贪心 + 背包dp 求最大美丽值: 1.字符串计数排序取后k个和 记为target
			问题转换，n个物品选k个, 且值为target的方案数, 设dp[j][k]表示取k个物品，价值为j的方案数
			dp[j][k] = (dp[j][k] + dp[j-v][k-1]*v) % M
 **/

var M = int(1e9) + 7

func countKSubsequencesWithMaxBeauty(s string, K int) int {
	var cnts = make([]int, 26)
	for _, v := range s {
		cnts[int(v-'a')]++
	}
	sort.Ints(cnts)
	var l = -1
	for i := 0; i < 26; i++ {
		if cnts[i] != 0 {
			break
		}
		l = i
	}
	cnts = cnts[l+1:]
	if len(cnts) < K {
		return 0
	}
	var n = len(cnts)
	var target = 0
	for i := len(cnts) - 1; i >= n-K; i-- {
		target += cnts[i]
	}
	// n个物品选k个, 且值为target的方案数
	var dp = make([][]int, target+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, K+1)
	}
	dp[0][0] = 1
	for _, v := range cnts {
		for j := target; j >= v; j-- {
			for k := 1; k <= K; k++ {
				dp[j][k] = (dp[j][k] + dp[j-v][k-1]*v) % M
			}
		}
	}
	return dp[target][K]
}
