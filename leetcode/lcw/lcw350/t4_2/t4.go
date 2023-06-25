package t4_1

import "math"

// 0-1背包问题
// 记忆化搜索 dfs(i,j)表示前i面墙，当前总共付费时间为j时的最小开销
//	        付费刷第i面墙, dfs(i-1, j+time[i])+cost[i]
//			免费刷第i面墙,需要1单位付费时间, dfs(i-1, j-1),
func paintWalls(cost []int, time []int) int {
	var n = len(cost)
	var mem = make([]map[int]int, n)
	for i := 0; i < len(mem); i++ {
		mem[i] = make(map[int]int)
	}
	var dfs func(i, j int) int
	dfs = func(i, j int) int {
		if j > i {
			return 0
		}
		if i < 0 {
			return math.MaxInt32 / 2
		}
		if v, exist := mem[i][j]; exist {
			return v
		}
		var res = min(dfs(i-1, j-1), dfs(i-1, j+time[i])+cost[i])
		mem[i][j] = res
		return res
	}
	return dfs(n-1, 0)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
