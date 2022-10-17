package test801

// dp[i][0]表示交换位置i, dp[i][1]表示不交换位置i
// 1.a[i]>a[i-1], b[i]>b[i-1], dp[i][0]=dp[i-1][0], dp[i][1] = dp[i-1][1]+1
// 2.a[i]>b[i-1], b[i]>a[i-1], dp[i][0]=dp[i-1[1], dp[i][1] = dp[i-1][0]+1
func minSwap(a []int, b []int) int {
	var n = len(a)
	var dp = make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, 2)
		for j := 0; j < 2; j++ {
			dp[i][j] = n
		}
	}
	dp[0][0], dp[0][1] = 0, 1
	for i := 1; i < n; i++ {
		if a[i] > a[i-1] && b[i] > b[i-1] {
			dp[i][0] = dp[i-1][0]
			dp[i][1] = dp[i-1][1] + 1
		}
		if a[i] > b[i-1] && b[i] > a[i-1] {
			dp[i][0] = min(dp[i][0], dp[i-1][1])
			dp[i][1] = min(dp[i][1], dp[i-1][0]+1)
		}
	}
	return min(dp[n-1][0], dp[n-1][1])
}

// 滚动数组优化
func minSwap1(a []int, b []int) int {
	var n = len(a)
	dp0, dp1 := 0, 1
	for i := 1; i < n; i++ {
		curDp0, preDp1 := n, dp1
		if a[i] > a[i-1] && b[i] > b[i-1] {
			curDp0 = min(curDp0, dp0)
			dp1 = dp1 + 1
		} else {
			dp1 = n
		}
		if a[i] > b[i-1] && b[i] > a[i-1] {
			dp1 = min(dp1, dp0+1)
			dp0 = min(curDp0, preDp1)
		}
	}
	return min(dp0, dp1)
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
