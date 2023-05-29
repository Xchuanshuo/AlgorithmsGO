package test1439

import "sort"

// 暴力 只保留k个最小的数参与下一轮计算即可

func kthSmallest1(g [][]int, k int) int {
	var pre = []int{0}
	for _, row := range g {
		var t = make([]int, 0)
		for _, p := range pre {
			for _, r := range row {
				t = append(t, p+r)
			}
		}
		sort.Ints(t)
		if len(t) > k {
			t = t[0:k]
		}
		pre = t
	}
	return pre[k-1]
}
