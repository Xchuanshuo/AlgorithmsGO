package lcw357

func canSplitArray(nums []int, m int) bool {
	var n = len(nums)
	var dp = make([][]bool, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]bool, n)
	}
	var s = make([]int, n+1)
	for i := 0; i < n; i++ {
		s[i+1] = s[i] + nums[i]
	}
	for l := 1; l <= n; l++ {
		for i := 0; i <= n-l; i++ {
			var j = i + l - 1
			if l <= 2 {
				dp[i][j] = true
			} else {
				dp[i][j] = (dp[i+1][j] && s[j+1]-s[i+1] >= m) || (dp[i][j-1] && s[j]-s[i] >= m)
			}
		}
	}
	return dp[0][n-1]
}
