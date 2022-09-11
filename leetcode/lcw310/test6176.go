package main

import "sort"

func mostFrequentEven(nums []int) int {
	sort.Ints(nums)
	var cnt = make(map[int]int)
	mx, res := 0, -1
	for _, num := range nums {
		if num%2 == 0 {
			cnt[num]++
			if cnt[num] > mx {
				mx = cnt[num]
				res = num
			}
		}
	}
	return res
}
