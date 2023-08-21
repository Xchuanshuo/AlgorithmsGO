package lcwt110

import (
	"sort"
)

/**
 * @Description https://leetcode.cn/problems/minimum-time-to-make-array-sum-at-most-x/
 * idea: 排序 + 背包dp
 **/

func minimumTime(nums1 []int, nums2 []int, x int) int {
	var n = len(nums1)
	var a = make([][]int, n)
	s1, s2 := 0, 0
	for i, v := range nums1 {
		a[i] = []int{v, nums2[i]}
		s1 += v
		s2 += nums2[i]
	}
	sort.Slice(a, func(i, j int) bool {
		return a[i][1] < a[j][1]
	})
	var dp = make([]int, n+1)
	for i := 1; i <= n; i++ {
		for j := i; j >= 1; j-- {
			dp[j] = max(dp[j], dp[j-1]+a[i-1][0]+j*a[i-1][1])
		}
	}
	for i := 0; i <= n; i++ {
		var s = s1 + s2*i
		if s-dp[i] <= x {
			return i
		}
	}

	return -1
}
