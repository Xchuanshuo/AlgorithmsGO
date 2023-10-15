package test910

import (
	"math"
	"sort"
)

/**
 * @Description https://leetcode.cn/problems/smallest-range-ii/
 * idea:
		1.排序 2.贪心，最右边值一定为-k, 左边一直+k
		2.枚举分界点i。分界点左边为+k, 右边为-k。注意边界情况
		max(nums[i]+k, nums[n-1]-k), min(nums[0]+k, nums[i+1]-k)
 **/

func smallestRangeII(nums []int, k int) int {
	var n = len(nums)
	nums = append(nums, int(1e9))
	sort.Ints(nums)
	var res = math.MaxInt32
	for i := 0; i < n; i++ {
		mx := max(nums[i]+k, nums[n-1]-k)
		mn := min(nums[0]+k, nums[i+1]-k)
		res = min(res, mx-mn)
	}
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
