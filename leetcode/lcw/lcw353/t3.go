package lcw353

func maxNonDecreasingLength(nums1 []int, nums2 []int) int {
	var n = len(nums1)
	var a = make([][]int, 0)
	for i := 0; i < n; i++ {
		a = append(a, []int{min(nums1[i], nums2[i]), max(nums1[i], nums2[i])})
	}
	var dp = make([][2]int, n)
	dp[0][0] = 1
	dp[0][1] = 1
	var res = 1
	for i := 1; i < n; i++ {
		dp[i][0], dp[i][1] = 1, 1
		if a[i][0] >= a[i-1][1] {
			dp[i][0] = dp[i-1][1] + 1
		}
		if a[i][0] >= a[i-1][0] {
			dp[i][0] = max(dp[i][0], dp[i-1][0]+1)
		}
		if a[i][1] >= a[i-1][0] {
			dp[i][1] = max(dp[i][1], dp[i-1][0]+1)
		}
		if a[i][1] >= a[i-1][1] {
			dp[i][1] = max(dp[i][1], dp[i-1][1]+1)
		}
		res = max(res, max(dp[i][0], dp[i][1]))
	}
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
