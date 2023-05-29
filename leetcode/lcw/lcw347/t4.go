package lcw347

import "sort"

func maxIncreasingCells(g [][]int) int {
	m, n := len(g), len(g[0])
	var mp = make(map[int][][]int)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			mp[g[i][j]] = append(mp[g[i][j]], []int{i, j})
		}
	}
	var a = make([]int, 0)
	for k := range mp {
		a = append(a, k)
	}
	sort.Ints(a)
	var rowMax = make([]int, m)
	var colMax = make([]int, n)
	var res = 0
	for _, key := range a {
		var vs = mp[key]
		var mx = make([]int, len(vs))
		// 根据上一轮结果计算答案
		for i, v := range vs {
			mx[i] = max(rowMax[v[0]], colMax[v[1]]) + 1
			res = max(res, mx[i])
		}
		// 更新最新结果
		for i, v := range vs {
			rowMax[v[0]] = max(rowMax[v[0]], mx[i])
			colMax[v[1]] = max(colMax[v[1]], mx[i])
		}
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
