package test462

import (
	"sort"
)

func minMoves2(nums []int) int {
	sort.Ints(nums)
	var n = len(nums)
	var num = nums[n/2]
	var res = 0
	for i := 0; i < n; i++ {
		res += abs(nums[i] - num)
	}
	return res
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
