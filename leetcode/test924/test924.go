package test924

import "sort"

func minMalwareSpread(graph [][]int, initial []int) int {
	var n = len(graph)
	var uf = NewDefUF()
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if graph[i][j] == 1 {
				uf.union(i, j)
			}
		}
	}
	var repeat = make(map[int]int)
	sort.Ints(initial)
	for _, init := range initial {
		repeat[uf.find(init)]++
	}
	mx, res := 0, initial[0]
	for _, init := range initial {
		if repeat[uf.find(init)] != 1 {
			continue
		}
		var cnt = uf.getCnt(init)
		if cnt > mx {
			mx = cnt
			res = init
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

func (uf *UF) getCnt(p int) int {
	return uf.sz[uf.find(p)]
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
