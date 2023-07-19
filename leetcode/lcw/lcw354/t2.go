package lcw354

import (
	"sort"
)

func maximumBeauty(nums []int, k int) int {
	sort.Ints(nums)
	mn, mx := nums[0], nums[len(nums)-1]
	var res = 0
	for i := mn; i <= mx; i++ {
		var l = sort.SearchInts(nums, i-k)
		var r = sort.SearchInts(nums, i+k+1)
		res = max(res, r-l)
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
