package test2009

import "sort"

func minOperations(nums []int) int {
	var n = len(nums)
	sort.Ints(nums)
	var set = make(map[int]bool)
	var a = make([]int, 0)
	for _, v := range nums {
		if !set[v] {
			set[v] = true
			a = append(a, v)
		}
	}
	var res = n
	for i, v := range a {
		var r = v + n - 1
		var j = sort.SearchInts(a, r+1)
		res = min(res, n-(j-i))
	}
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
