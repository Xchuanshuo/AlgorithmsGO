package test684

/**
 * @Description https://leetcode.cn/problems/redundant-connection-ii/
 * idea:  根节点无父节点 非根节点只有1个父节点
		  入度判断 1.无环图  2.有环图
		  删除的边: 1.度为2的边中的一条  2.同t684
 **/

func findRedundantDirectedConnection(edges [][]int) []int {
	var n = len(edges)
	var uf = NewUF(make([]int, n+1))
	var inDegree = make(map[int]int)
	for _, e := range edges {
		inDegree[e[1]]++
	}
	var delEdges = make([][]int, 0)
	for _, e := range edges {
		if inDegree[e[1]] == 2 {
			delEdges = append(delEdges, e)
		}
	}
	if len(delEdges) == 0 {
		for _, e := range edges {
			p, q := uf.find(e[0]), uf.find(e[1])
			if p == q {
				return e
			}
			uf.union(p, q)
		}
	} else {
		var t = delEdges[1]
		var correct = true
		for _, e := range edges {
			if e[0] == t[0] && e[1] == t[1] {
				continue
			}
			p, q := uf.find(e[0]), uf.find(e[1])
			if p == q {
				correct = false
				break
			}
			uf.union(p, q)
		}
		if correct {
			return delEdges[1]
		}
		return delEdges[0]
	}
	return []int{}
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
