package test933

import "sort"

type RecentCounter struct {
	arr []int
}

func Constructor() RecentCounter {
	return RecentCounter{arr: make([]int, 0)}
}

func (this *RecentCounter) Ping(t int) int {
	this.arr = append(this.arr, t)
	var left = t - 3000
	var n = len(this.arr)
	var idx = sort.Search(n, func(i int) bool {
		return this.arr[i] >= left
	})
	return n - idx
}
