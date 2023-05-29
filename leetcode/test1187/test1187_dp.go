package test1187

import (
	"math"
	"sort"
)

func makeArrayIncreasing(a []int, arr2 []int) int {
	sort.Ints(arr2)
	var b = make([]int, 0)
	for _, v := range arr2 {
		if len(b) == 0 || b[len(b)-1] != v {
			b = append(b, v)
		}
	}
	var INF = math.MaxInt32 / 2
	a = append([]int{-INF}, a...)
	a = append(a, INF)
	var n = len(a)
	var dp = make([]int, n)
	for i := 1; i < n; i++ {
		dp[i] = INF
		if a[i] > a[i-1] {
			// 不替换
			dp[i] = dp[i-1]
		}
		var j = sort.SearchInts(b, a[i])
		// 替换, 枚举替换个数k
		for k := 1; k <= min(i-1, j); k++ {
			if b[j-k] > a[i-1-k] {
				dp[i] = min(dp[i], dp[i-1-k]+k)
			}
		}
	}
	if dp[n-1] >= INF {
		return -1
	}
	return dp[n-1]
}
