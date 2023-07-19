package test2111

import "sort"

/**
 * @Description https://leetcode.cn/problems/minimum-operations-to-make-the-array-k-increasing/
 * idea: 使k递增最少操作次数 -> 按k为步长进行分组，求每组的LIS，答案为最后总长度-所有组的LIS和
 **/

func kIncreasing(arr []int, k int) int {
	var n = len(arr)
	var calc = func(a []int) int {
		var n = len(a)
		var a1 = make([]int, n+1)
		var l = 1
		a1[0] = a[0]
		for i := 1; i < n; i++ {
			if a[i] >= a1[l-1] {
				a1[l] = a[i]
				l++
			} else {
				var pos = sort.Search(l, func(j int) bool {
					return a1[j] > a[i]
				})
				a1[pos] = a[i]
			}
		}
		return l
	}
	var mx = 0
	for l := 1; l <= k; l++ {
		var a = make([]int, 0)
		for j := l - 1; j < n; j += k {
			a = append(a, arr[j])
		}
		mx += calc(a)
	}
	return n - mx
}
