package test1697

import "sort"

func distanceLimitedPathsExist(n int, es [][]int, queries [][]int) []bool {
	sort.Slice(es, func(i, j int) bool {
		return es[i][2] < es[j][2]
	})
	var k = len(queries)
	var ids = make([]int, k)
	for i := 0; i < k; i++ {
		ids[i] = i
	}
	sort.Slice(ids, func(i, j int) bool {
		return queries[ids[i]][2] < queries[ids[j]][2]
	})
	var uf = NewUF(make([]int, n))
	var res = make([]bool, k)
	var j = 0
	for _, id := range ids {
		var qv = queries[id][2]
		for ; j < len(es) && es[j][2] < qv; j++ {
			uf.union(es[j][0], es[j][1])
		}
		res[id] = uf.isConnected(queries[id][0], queries[id][1])
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
	var parent = uf.parent
	for parent[p] != p {
		parent[p] = parent[parent[p]]
		p = parent[p]
	}
	return p
}
