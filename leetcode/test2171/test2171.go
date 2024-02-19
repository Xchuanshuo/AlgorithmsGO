package test2171

import "sort"

func minimumRemoval(beans []int) int64 {
	sort.Ints(beans)
	var n = len(beans)
	var sums = make([]int64, n+1)
	for i, v := range beans {
		sums[i+1] = sums[i] + int64(v)
	}
	var res = sums[n]
	for i := 1; i <= n; i++ {
		var t = int64(beans[i-1] * (n - i))
		var cur = sums[i-1] + (sums[n] - sums[i] - t)
		res = min(res, cur)
	}
	return res
}
