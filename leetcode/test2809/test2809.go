package test2809

import "sort"

/**
 * @Description https://leetcode.cn/problems/minimum-time-to-make-array-sum-at-most-x
 * idea: 背包dp + 排序贪心;
		1.对一个位置多次置0，相当于前面的重置都无效，所以每个位置【最多置0一次】，整个数组最多重置len(nums1)次
		2.设s1、s2分别为nums1、nums2的和。假设经过了t秒，当前总和为 s1 + s2*t，需要考虑【如何分配t次操作】，使元素和最小
		3.正向计算比较麻烦，考虑计算到第j秒减少的最大值，设dp[j]为减少的最大值，则最小总和为 s1+s2*j - dp[j]
		4.考察每元素位置i，当前j秒可以选择是否重置，不重置则减少最大值不变，重置则为dp[j-1] + nums1[i] + nums2[i]*j
		即				dp[j] = max(dp[j]，dp[j-1] + nums1[i] + nums2[i]*j)
		5.元素顺序，nums2[i]较大的元素放在较后考虑实际能减少更多的值，所以需要排序后，元素值较大的后处理
 **/

func minimumTime(nums1 []int, nums2 []int, x int) int {
	var n = len(nums1)
	s1, s2 := 0, 0
	var pos = make([]int, 0)
	for i := 0; i < n; i++ {
		pos = append(pos, i)
		s1 += nums1[i]
		s2 += nums2[i]
	}
	sort.Slice(pos, func(i, j int) bool {
		return nums2[pos[i]] < nums2[pos[j]]
	})
	var dp = make([]int, n+1)
	for i := 0; i < n; i++ {
		for j := n; j >= 1; j-- {
			var p = pos[i]
			dp[j] = max(dp[j], dp[j-1]+nums1[p]+j*nums2[p])
		}
	}
	for j := 0; j <= n; j++ {
		var total = s1 + s2*j
		if total-dp[j] <= x {
			return j
		}
	}
	return -1
}
