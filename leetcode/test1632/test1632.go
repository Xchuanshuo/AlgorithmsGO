package test1632

import "sort"

/**
 * @Description https://leetcode.cn/problems/rank-transform-of-a-matrix/
 * idea:
	// 1.矩阵每个格子的rank被所有行列的rank共享
	// 2.需要先处理value较小的格子
	// 3.使用并查集将value相同的格子合并
	// 4.对相同value的格子进行统一处理，rank值取当前最大rank值+1
	// 5.保存每一个行列当前最大rank值即可
 **/
func matrixRankTransform(matrix [][]int) [][]int {
	m, n := len(matrix), len(matrix[0])
	var mp = make(map[int][][]int)
	var vArr = make([]int, 0)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			var v = matrix[i][j]
			if _, exist := mp[v]; !exist {
				mp[v] = make([][]int, 0)
				vArr = append(vArr, v)
			}
			mp[v] = append(mp[v], []int{i, j})
		}
	}
	sort.Ints(vArr)
	rows, cols := make([]int, m), make([]int, n)
	var rank = make([]int, m+n)
	var res = make([][]int, m)
	for i := 0; i < len(res); i++ {
		res[i] = make([]int, n)
	}
	for _, val := range vArr {
		var values = mp[val]
		var uf = NewUF(m + n)
		for _, v := range values {
			uf.union(v[0], v[1]+m)
		}
		for _, v := range values {
			x, y := v[0], v[1]
			rank[uf.find(x)] = max(rank[uf.find(x)], max(rows[x], cols[y]))
		}
		for _, v := range values {
			x, y := v[0], v[1]
			res[x][y] = 1 + rank[uf.find(x)]
			rows[x] = res[x][y]
			cols[y] = res[x][y]
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

type UF struct {
	parent []int
	cnt    int
}

func NewUF(n int) *UF {
	var parent = make([]int, n)
	for i := 0; i < n; i++ {
		parent[i] = i
	}
	return &UF{parent, n}
}

func (uf *UF) find(id int) int {
	for id != uf.parent[id] {
		id = uf.parent[id]
	}
	return uf.parent[id]
}

func (uf *UF) union(p, q int) {
	var pRoot = uf.find(p)
	var qRoot = uf.find(q)
	if pRoot == qRoot {
		return
	}
	uf.parent[pRoot] = qRoot
	uf.cnt = uf.cnt - 1
}
