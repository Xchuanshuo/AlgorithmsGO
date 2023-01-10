package test928

/**
 * @Description https://leetcode.cn/problems/minimize-malware-spread-ii/
 * idea: 枚举每个节点去除后, 其它病毒能感染的联通分量大小
 **/

import "sort"

func minMalwareSpread(graph [][]int, initial []int) int {
	var n = len(graph)
	sort.Ints(initial)
	min, res := int(1e9), initial[0]
	for _, init := range initial {
		var uf = NewDefUF()
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				if i == init || j == init {
					continue
				}
				if graph[i][j] == 1 {
					uf.union(i, j)
				}
			}
		}
		var set = make(map[int]bool)
		var cnt = 0
		for _, t := range initial {
			if t == init {
				continue
			}
			var p = uf.find(t)
			if !set[p] {
				cnt += uf.sz[p]
				set[p] = true
			}
		}
		if cnt < min {
			min = cnt
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
