package lcw353

func maximumJumps(nums []int, target int) int {
	var n = len(nums)
	var dp = make([]int, n)
	for i := 1; i < n; i++ {
		dp[i] = -1
		for j := i - 1; j >= 0; j-- {
			if abs(nums[i]-nums[j]) <= target && dp[j] != -1 {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
	}
	return dp[n-1]
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
