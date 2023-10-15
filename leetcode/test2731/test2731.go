package test2731

import "sort"

var M = int(1e9) + 7

func sumDistance(nums []int, s string, d int) int {
	for i := range nums {
		if s[i] == 'L' {
			nums[i] -= d
		} else {
			nums[i] += d
		}
	}
	sort.Ints(nums)
	var res = 0
	var sum = 0
	for i, x := range nums {
		res = (res + i*x - sum) % M
		sum += x
	}
	return res
}
