package test1403

import "sort"

func minSubsequence(nums []int) []int {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] > nums[j]
	})
	var res = make([]int, 0)
	sum, total := 0, 0
	for _, num := range nums {
		sum += num
	}
	for _, num := range nums {
		if total > sum-total {
			return res
		}
		res = append(res, num)
		total += num
	}
	return res
}
