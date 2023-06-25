package dfs

import "math"

func connectTwoGroups(cost [][]int) int {
	m, n := len(cost), len(cost[0])
	var mem = make([][]int, m)
	for i := 0; i < len(mem); i++ {
		mem[i] = make([]int, 1<<n)
		for j := 0; j < len(mem[i]); j++ {
			mem[i][j] = -1
		}
	}
	var minCost = make([]int, n)
	for i := 0; i < n; i++ {
		var cur = cost[0][i]
		for j := 0; j < m; j++ {
			cur = min(cur, cost[j][i])
		}
		minCost[i] = cur
	}
	var dfs func(i, j int) int
	dfs = func(i, j int) int {
		if i < 0 {
			var sum = 0
			for k := 0; k < n; k++ {
				if j&(1<<k) == 0 {
					continue
				}
				sum += minCost[k]
			}
			return sum
		}
		if mem[i][j] != -1 {
			return mem[i][j]
		}
		var res = math.MaxInt32 / 2
		for k := 0; k < n; k++ {
			res = min(res, dfs(i-1, j&^(1<<k))+cost[i][k])
		}
		mem[i][j] = res
		return res
	}
	// 递归入口 A集合全部未选, B集合全部未选
	return dfs(m-1, 1<<n-1)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
