package main

func longestSubarray(nums []int) int {
	var max = 1
	for _, v := range nums {
		if v > max {
			max = v
		}
	}
	cnt, res := 0, 1
	for i := 0; i < len(nums); i++ {
		if nums[i] == max {
			cnt++
		} else {
			cnt = 0
		}
		if cnt > res {
			res = cnt
		}
	}
	return res
}
