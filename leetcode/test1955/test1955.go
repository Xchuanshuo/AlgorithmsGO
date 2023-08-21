package test1955

// dp[i][j]表示前i个数以j结尾的方案数，分情况讨论。
//	若当前元素为0, 则dp[i][0]=2*dp[i-1][0]+1, 即前面i-1个数中以0结尾的 序列+当前位置，以及当前单个元素作为结尾的序列, 再累加上之前0结尾的总数即可;
//  若当前元素为1, 则dp[i][1]=(dp[i-1][0]+2*dp[i-1][1])，即前面i-1个数以0结尾的序列+前面i-1个数以1结尾的序列 + 当前元素；
//  若当前元素为2，同1，最终答案为dp[n][2]

var M = int(1e9) + 7

func countSpecialSubsequences(nums []int) int {
	var n = len(nums)
	var dp = make([][3]int, n+1)
	for i := 1; i <= n; i++ {
		var x = nums[i-1]
		if x == 0 {
			dp[i][0] = (2*dp[i-1][0] + 1) % M
			dp[i][1] = dp[i-1][1]
			dp[i][2] = dp[i-1][2]
		} else if x == 1 {
			dp[i][0] = dp[i-1][0]
			dp[i][1] = (dp[i-1][0] + 2*dp[i-1][1]) % M
			dp[i][2] = dp[i-1][2]
		} else if x == 2 {
			dp[i][0] = dp[i-1][0]
			dp[i][1] = dp[i-1][1]
			dp[i][2] = (dp[i-1][1] + 2*dp[i-1][2]) % M
		}
	}
	return dp[n][2]
}
