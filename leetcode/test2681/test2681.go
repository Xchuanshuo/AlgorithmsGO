package test2681

import "sort"

var M = int(1e9) + 7

func sumOfPower(nums []int) int {
	sort.Ints(nums)
	res, pre := 0, 0
	for _, v := range nums {
		res = (res + (v*v)%M*(pre+v)%M) % M
		pre = (pre*2 + v) % M
	}
	return res
}
