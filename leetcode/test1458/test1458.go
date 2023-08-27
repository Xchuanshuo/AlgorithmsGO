package test1458

/**
 * @Description https://leetcode.cn/problems/max-dot-product-of-two-subsequences
 * idea: 注意点：由于【必选】一对元素，所以初始值赋较小的值. 选取数对后，dp计算max时,
			使用max(0, dp[i-1][j-1])，防止前面的负数将最大值减小
 **/

func maxDotProduct(nums1 []int, nums2 []int) int {
	m, n := len(nums1), len(nums2)
	var dp = make([][]int, m+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, n+1)
		for j := 0; j < len(dp[i]); j++ {
			dp[i][j] = -int(1e9)
		}
	}
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			dp[i][j] = max(dp[i][j-1], max(dp[i-1][j], max(0, dp[i-1][j-1])+nums1[i-1]*nums2[j-1]))
		}
	}
	return dp[m][n]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
