package main

import "fmt"

func missingTwo(nums []int) []int {
	var n = len(nums)
	nums = append(nums, 0, 0)
	for i, v := range nums {
		if v == 0 || i >= n {
			continue
		}
		var cur = v
		if cur < 0 {
			cur = -v
		}
		if nums[cur-1] == 0 {
			nums[cur-1] = -1
		} else {
			nums[cur-1] *= -1
		}
	}
	var res = make([]int, 0)
	for i, v := range nums {
		if v >= 0 {
			res = append(res, i+1)
		}
	}
	return res
}

func main() {
	var arr = []int{2}
	var res = missingTwo(arr)
	fmt.Println(res)
}
