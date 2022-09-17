package test1608

import "sort"

func specialArray(nums []int) int {
	sort.Ints(nums)
	for i := 0; i <= 1000; i++ {
		var idx = sort.SearchInts(nums, i)
		if i == len(nums)-idx {
			return i
		}
	}
	return -1
}
