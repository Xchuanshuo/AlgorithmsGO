package test2902

/**
 * @Description https://leetcode.cn/problems/count-of-sub-multisets-with-bounded-sum
 * idea: 多重背包. 问题实际为凑成指定大小背包的方案数。需要解决以下几个问题
		1.如何去重? 同样大小的背包依次计算取数量1..cnt时的方案数即可，保证每类背包不同数量的组合只会出现一次。
		2.按照常规背包计算方法，复杂度为O(n^2)，考虑如何进行优化？从公式入手
		 背包大小为j时装大小为x的物品方案数为 dp[j] = dp[j] + dp[j-x] + dp[j-2*x] + ... + dp[j-cnt*x]
		 背包大小为j-x时装大小为x物品的方案  dp[j-x] = dp[j-x] + dp[j-2*x] + dp[j-3*x] + ... + dp[j-(cnt+1)*x]
		 两个式子相减得 dp[j]-dp[j-x] = dp[j] - dp[j-(cnt+1)*x],
				即     dp[j] = dp[j] + dp[j-x](当前) - dp[j-(cnt+1)*x]
 **/

var M = int(1e9) + 7

func countSubMultisets(nums []int, l int, r int) int {
	var dp0 = make([]int, r+1)
	var cnts = make(map[int]int)
	for _, v := range nums {
		cnts[v]++
	}
	// 滚动数组优化
	dp0[0] = cnts[0] + 1
	for x, cnt := range cnts {
		var dp = make([]int, r+1)
		for j := 0; j <= r; j++ {
			dp[j] = dp0[j]
			if j-x < 0 {
				continue
			}
			dp[j] = (dp[j] + dp[j-x]) % M
			if j-(cnt+1)*x >= 0 {
				dp[j] = (dp[j] - dp0[j-(cnt+1)*x] + M) % M
			}
		}
		dp0 = dp
	}
	var res = 0
	for i := l; i <= r; i++ {
		res = (res + dp0[i]) % M
	}
	return res
}
