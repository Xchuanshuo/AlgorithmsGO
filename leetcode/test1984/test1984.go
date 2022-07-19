package test1984

import "sort"

func minimumDifference(nums []int, k int) int {
	sort.Ints(nums)
	res := int(1e9)
	for i := k - 1; i < len(nums); i++ {
		l := i - k + 1
		if nums[i]-nums[l] < res {
			res = nums[i] - nums[l]
		}
	}
	return res
}
