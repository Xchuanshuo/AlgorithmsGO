package main

import (
	"fmt"
	"sort"
)

func goodIndices(nums []int, k int) []int {
	var n = len(nums)
	var inc = make([]bool, n)
	last, cnt := nums[0], 0
	for i, v := range nums {
		if v >= last {
			cnt++
		} else {
			cnt = 1
		}
		if cnt >= k && i-k > 0 {
			inc[i-k] = true
		}
		last = v
	}
	last, cnt = nums[n-1], 0
	var res = make([]int, 0)
	for i := n - 1; i >= 0; i-- {
		var v = nums[i]
		if v >= last {
			cnt++
		} else {
			cnt = 1
		}
		if cnt >= k && i+k < n && inc[i+k] {
			res = append(res, i+k)
		}
		last = v
	}
	sort.Ints(res)
	return res
}

func main() {
	var arr = []int{2, 1, 1, 2}
	res := goodIndices(arr, 2)
	fmt.Println(res)
}
