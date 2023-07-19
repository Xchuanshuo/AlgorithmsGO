package test2713

import "sort"

func maxIncreasingCells(mat [][]int) int {
	m, n := len(mat), len(mat[0])
	var mp = make(map[int][][]int, 0)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			mp[mat[i][j]] = append(mp[mat[i][j]], []int{i, j})
		}
	}
	var arr = make([]int, 0)
	for k := range mp {
		arr = append(arr, k)
	}
	sort.Ints(arr)
	var res = 0
	var rowMx = make([]int, m)
	var colMx = make([]int, n)
	for _, key := range arr {
		var t = make([]int, len(mp[key]))
		for i, v := range mp[key] {
			t[i] = max(rowMx[v[0]], colMx[v[1]]) + 1
			res = max(res, t[i])
		}
		for i, v := range mp[key] {
			rowMx[v[0]] = max(rowMx[v[0]], t[i])
			colMx[v[1]] = max(colMx[v[1]], t[i])
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
