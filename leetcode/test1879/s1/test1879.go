package s1

import "math"

func minimumXORSum(nums1 []int, nums2 []int) int {
	var n = len(nums1)
	var l = 1 << n
	var dp = make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, l)
		for j := 0; j < len(dp[i]); j++ {
			dp[i][j] = -1
		}
	}
	var dfs func(l, vis int) int
	dfs = func(l, vis int) int {
		if l == n {
			return 0
		}
		if dp[l][vis] != -1 {
			return dp[l][vis]
		}
		var res = math.MaxInt32
		for i := 0; i < n; i++ {
			if vis&(1<<i) != 0 {
				continue
			}
			res = min(res, dfs(l+1, vis|(1<<i))+(nums1[l]^nums2[i]))
		}
		dp[l][vis] = res
		return res
	}
	return dfs(0, 0)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
