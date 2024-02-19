package lcwt116

// dp[i]：和为i的子序列最大长度
func lengthOfLongestSubsequence(nums []int, target int) int {
	var dp = make([]int, target+1)
	for i := 1; i < len(dp); i++ {
		dp[i] = -1
	}
	for _, v := range nums {
		for j := target; j >= v; j-- {
			if dp[j-v] >= 0 {
				dp[j] = max(dp[j], dp[j-v]+1)
			}
		}
	}
	return dp[target]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
