package lcwt100

import "sort"

func maximizeGreatness(nums []int) int {
	sort.Ints(nums)
	var t = make([]int, len(nums))
	copy(t, nums)
	var res = 0
	var r = len(nums) - 1
	for i := len(nums) - 1; i >= 0; i-- {
		if nums[r] > nums[i] {
			res++
			r--
		}
	}
	return res
}
