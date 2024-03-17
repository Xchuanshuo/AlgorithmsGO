package main

import "sort"

func unmarkedSumArray(nums []int, queries [][]int) []int64 {
	var sum = int64(0)
	var ids = make([]int, 0)
	for i, v := range nums {
		sum += int64(v)
		ids = append(ids, i)
	}
	sort.Slice(ids, func(i, j int) bool {
		if nums[ids[i]] == nums[ids[j]] {
			return ids[i] < ids[j]
		}
		return nums[ids[i]] < nums[ids[j]]
	})
	var m = len(queries)
	var res = make([]int64, m)
	var mark = make([]bool, len(nums))
	var j = 0
	for k, q := range queries {
		if !mark[q[0]] {
			sum -= int64(nums[q[0]])
		}
		mark[q[0]] = true
		var i = 0
		for ; i < q[1] && j < len(ids); j++ {
			if mark[ids[j]] {
				continue
			}
			mark[ids[j]] = true
			sum -= int64(nums[ids[j]])
			i++
		}
		res[k] = sum
	}
	return res
}