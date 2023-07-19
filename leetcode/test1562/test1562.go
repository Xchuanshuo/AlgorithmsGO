package test1562

/**
 * @Description https://leetcode.cn/problems/find-latest-group-of-size-m/
 * idea: 并查集，每个步骤合并左右连续的1时存在数量为m的记录为最新答案
 **/

func findLatestStep(arr []int, m int) int {
	var n = len(arr)
	if m == n {
		return n
	}
	var uf = NewDefUF()
	var res = -1
	for i := 0; i < n; i++ {
		uf.union(arr[i], arr[i])
		pre, nxt := arr[i]-1, arr[i]+1
		if uf.sz[pre] != 0 {
			if uf.getCnt(pre) == m {
				res = i
			}
			uf.union(pre, arr[i])
		}
		if uf.sz[nxt] != 0 {
			if uf.getCnt(nxt) == m {
				res = i
			}
			uf.union(nxt, arr[i])
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
