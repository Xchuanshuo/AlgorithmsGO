package main

func maxFrequencyElements(nums []int) int {
	var cnt = make(map[int]int)
	var maxV = 0
	var res = 0
	for _, v := range nums {
		cnt[v]++
		maxV = max(maxV, cnt[v])
	}
	for _, v := range cnt {
		if v == maxV {
			res += v
		}
	}
	return res
}
