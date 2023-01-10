package lcw323

import "sort"

/**
 * @Description https://leetcode.cn/problems/maximum-number-of-points-from-grid-queries/
 * idea: 离线查询 + 并查集
		 1.将矩阵所有点按大小排序, 将所有查询值按大小排序
		 2.使用并查集 从小到大依次合并所有小于当前查询值qv的点
		 3.若左上角开始点也在并差集中 说明本轮合并的所有点都满足条件 计入答案即可
    同 test1697
 **/

var dirs = [][]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}

func maxPoints(grid [][]int, queries []int) []int {
	var k = len(queries)
	var ids = make([]int, k)
	for i := 0; i < k; i++ {
		ids[i] = i
	}
	sort.Slice(ids, func(i, j int) bool {
		return queries[ids[i]] < queries[ids[j]]
	})
	m, n := len(grid), len(grid[0])
	var q = make([][]int, 0)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			q = append(q, []int{grid[i][j], i, j})
		}
	}
	sort.Slice(q, func(i, j int) bool {
		return q[i][0] < q[j][0]
	})
	var uf = NewUF(make([]int, m*n))
	var res = make([]int, k)
	var j = 0
	for _, id := range ids {
		var qv = queries[id]
		for ; j < m*n && q[j][0] < qv; j++ {
			x, y := q[j][1], q[j][2]
			for i := 0; i < len(dirs); i++ {
				nx, ny := x+dirs[i][0], y+dirs[i][1]
				if nx < 0 || nx >= m || ny < 0 || ny >= n || grid[nx][ny] >= qv {
					continue
				}
				uf.union(x*n+y, nx*n+ny)
			}
		}
		if qv > grid[0][0] {
			res[id] = uf.sz[uf.find(0)]
		}
	}
	return res
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
	if uf.parent[p] == 0 {
		return p
	}
	var parent = uf.parent
	for parent[p] != p {
		parent[p] = parent[parent[p]]
		p = parent[p]
	}
	return p
}
