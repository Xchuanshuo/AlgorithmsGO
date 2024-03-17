package test2369

func validPartition(nums []int) bool {
	var n = len(nums)
	var dp = make([]bool, n+1)
	dp[0] = true
	for i := 1; i <= n; i++ {
		var valid = false
		if i >= 2 {
			valid = dp[i-2] && nums[i-1] == nums[i-2]
		}
		if !valid && i >= 3 {
			valid = dp[i-3] &&
				(nums[i-1] == nums[i-2] && nums[i-2] == nums[i-3] ||
					nums[i-3]-nums[i-2] == -1 && nums[i-2]-nums[i-1] == -1)
		}
		dp[i] = valid
	}
	return dp[n]
}
