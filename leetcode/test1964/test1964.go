package test1964

import "sort"

func longestObstacleCourseAtEachPosition(os []int) []int {
	var n = len(os)
	var arr = make([]int, n+1)
	var res = make([]int, n)
	res[0], arr[0] = 1, os[0]
	var l = 1
	for i := 1; i < n; i++ {
		res[i] = 1
		if os[i] >= arr[l-1] {
			arr[l] = os[i]
			l++
			res[i] = l
		} else {
			var pos = sort.Search(l, func(i int) bool {
				return arr[i] >= os[i]
			})
			res[i] = pos + 1
			arr[pos] = os[i]
		}
	}
	return res
}
