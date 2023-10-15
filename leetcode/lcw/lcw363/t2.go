package lcw363

import "sort"

func countWays(nums []int) int {
	var n = len(nums)
	sort.Ints(nums)
	var res = 1      // 全选
	if nums[0] > 0 { // 全不选
		res++
	}
	nums = append(nums, -1)
	for i := 0; i < n-1; i++ {
		var x = i + 1
		if x > nums[i] && x < nums[i+1] {
			res++
		}
	}
	return res
}
