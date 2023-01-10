package test1631

/**
 * @Description https://leetcode.cn/problems/path-with-minimum-effort
 * idea: 并查集 本题按相邻点的绝对值作为体力值，所以考虑枚举所有边全 使用并查集判断连通性
		1.将所有边按边权【从小到大】排序 (右下)
		2.【从小到大】枚举权值，使用并查集依次合并小于等于指定权值的边
		3.判定起点、终点是否联通
 **/

import (
	"sort"
)

var dirs = [][]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}

func minimumEffortPath(h [][]int) int {
	m, n := len(h), len(h[0])
	var q = make([][]int, 0)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			var pos = i*n + j
			if i+1 < m {
				q = append(q, []int{abs(h[i][j] - h[i+1][j]), pos, (i+1)*n + j})
			}
			if j+1 < n {
				q = append(q, []int{abs(h[i][j] - h[i][j+1]), pos, i*n + j + 1})
			}
		}
	}
	sort.Slice(q, func(i, j int) bool {
		return q[i][0] < q[j][0]
	})
	var uf = NewDefUF()
	for j := 0; j < len(q); j++ {
		uf.union(q[j][1], q[j][2])
		if uf.find(0) == uf.find(m*n-1) {
			return q[j][0]
		}
	}
	return 0
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

type UF struct {
	parent map[int]int
	sz     map[int]int
}

func NewUF(arr []int) *UF {
	var parent = make(map[int]int)
	var sz = make(map[int]int)
	for i := 0; i < len(arr); i++ {
		parent[i] = i
		sz[i] = 1
	}
	var uf = &UF{parent: parent, sz: sz}
	return uf
}

func NewDefUF() *UF {
	return NewUF([]int{})
}

func (uf *UF) union(p, q int) {
	var pRoot = uf.find(p)
	var qRoot = uf.find(q)
	var sz = uf.sz
	if pRoot == qRoot {
		return
	}
	if sz[pRoot] > sz[qRoot] {
		sz[pRoot] += sz[qRoot]
		uf.parent[qRoot] = pRoot
	} else {
		sz[qRoot] += sz[pRoot]
		uf.parent[pRoot] = qRoot
	}
}

func (uf *UF) isConnected(p, q int) bool {
	return uf.find(p) == uf.find(q)
}

func (uf *UF) find(p int) int {
	if _, exist := uf.parent[p]; !exist {
		uf.parent[p] = p
		uf.sz[p] = 1
	}
	var parent = uf.parent
	for parent[p] != p {
		parent[p] = parent[parent[p]]
		p = parent[p]
	}
	return p
}
