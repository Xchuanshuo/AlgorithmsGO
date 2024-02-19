package lcw368

import (
	"math"
)

func minimumSum(nums []int) int {
	var n = len(nums)
	var rm = make([]int, n)
	rm[n-1] = n - 1
	for i := n - 2; i >= 0; i-- {
		if nums[i] < nums[rm[i+1]] {
			rm[i] = i
		} else {
			rm[i] = rm[i+1]
		}
	}
	var lm = 0
	var res = math.MaxInt32
	for i := 1; i < n-1; i++ {
		if nums[lm] < nums[i] && nums[i] > nums[rm[i+1]] {
			res = min(res, nums[lm]+nums[i]+nums[rm[i+1]])
		}
		if nums[i] < nums[lm] {
			lm = i
		}
	}
	if res == math.MaxInt32 {
		return -1
	}
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
