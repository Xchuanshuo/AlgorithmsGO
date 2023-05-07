package test1187

import (
	"math"
	"sort"
)

var INF = math.MaxInt32 / 2

func makeArrayIncreasing(a1 []int, a2 []int) int {
	var n = len(a1)
	sort.Ints(a2)
	var mem = make([]map[int]int, n)
	for i := 0; i < len(mem); i++ {
		mem[i] = make(map[int]int)
	}
	var dfs func(i, limit int) int
	dfs = func(i, limit int) int {
		if i < 0 {
			return 0
		}
		if v, exist := mem[i][limit]; exist {
			return v
		}
		var cur = INF
		// 当前不调整
		if a1[i] < limit {
			cur = dfs(i-1, a1[i])
		}
		// 调整, 找到小于limit的最大值
		var j = sort.SearchInts(a2, limit) - 1
		if j >= 0 {
			cur = min(cur, dfs(i-1, a2[j])+1)
		}
		mem[i][limit] = cur
		return cur
	}
	var res = dfs(n-1, INF)
	if res < INF {
		return res
	}
	return -1
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
