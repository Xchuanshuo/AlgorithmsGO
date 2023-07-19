package lcw351

var M = int(1e9) + 7
func numberOfGoodSubarraySplits(nums []int) int {
	var n = len(nums)
	var dp = make([]int, n+1)
	var cnt0 = 0
	for i := 1;i <= n;i++ {
		if nums[i-1] == 0 {
			dp[i] = dp[i-1]
			cnt0++
		} else {
			if dp[i-1] == 0 {
				dp[i] = 1
			} else {
				dp[i] = (dp[i-1] * (cnt0+1)) %M
			}
			cnt0 = 0
		}
	}
	return dp[n]
}