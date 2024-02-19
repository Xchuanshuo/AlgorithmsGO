package lcw369

import "math"

// dp[i]表示前i个数修改为增量数组所需代价
func minIncrementOperations(nums []int, k int) int64 {
	var n = len(nums)
	var dp = make([]int64, n)
	dp[0] = maxI64(0, int64(k-nums[0]))
	dp[1] = maxI64(0, int64(k-nums[1]))
	dp[2] = maxI64(0, int64(k-nums[2]))
	for i := 3; i < n; i++ {
		dp[i] = math.MaxInt64 / 2
		for j := i - 1; j >= i-3; j-- {
			dp[i] = minI64(dp[i], dp[j]+maxI64(0, int64(k-nums[i])))
		}
	}
	return minI64(dp[n-1], minI64(dp[n-2], dp[n-3]))
}

func minI64(a, b int64) int64 {
	if a < b {
		return a
	}
	return b
}
