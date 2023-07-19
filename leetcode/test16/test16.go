package test16

import "sort"

func threeSumClosest(nums []int, target int) int {
	var n = len(nums)
	sort.Ints(nums)
	var res = nums[0] + nums[1] + nums[2]
	for i := 0; i < n; i++ {
		l, r := i+1, n-1
		for l < r {
			var s = nums[i] + nums[l] + nums[r]
			if s == target {
				return target
			}
			if s > target {
				r--
			} else {
				l++
			}
			if abs(s-target) < abs(res-target) {
				res = s
			}
		}
	}
	return res
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
