package test1388

/**
 * @Description https://leetcode.cn/problems/pizza-with-3n-slices/description/
 * idea: 3n披萨中任意挑n块 => 长度为3n的序列中任选n个不相邻的数，最大和为多少
 **/

func maxSizeSlices(slices []int) int {
	var v1 = calc(slices[0 : len(slices)-1])
	var v2 = calc(slices[1:])
	return max(v1, v2)
}

func calc(slices []int) int {
	var n = len(slices)
	var c = (n + 1) / 3
	var dp = make([][]int, n+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, c+1)
	}
	dp[1][1] = slices[0]
	for i := 2; i <= n; i++ {
		for j := 1; j <= c && j <= i; j++ {
			dp[i][j] = max(dp[i-1][j], dp[i-2][j-1]+slices[i-1])
		}
	}
	return dp[n][c]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
