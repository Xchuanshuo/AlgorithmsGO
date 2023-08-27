package test2076

// 并查集，对于每个请求，计算用户所属连通分量，再看限制好友的所属连通分量是否为当前好友请求的

func friendRequests(n int, restrictions [][]int, requests [][]int) []bool {
	var uf = NewDefUF()
	var res = make([]bool, len(requests))
	for i, req := range requests {
		p1, p2 := uf.find(req[0]), uf.find(req[1])
		var valid = true
		for _, r := range restrictions {
			c1, c2 := uf.find(r[0]), uf.find(r[1])
			if (c1 == p1 && c2 == p2) || (c1 == p2 && c2 == p1) {
				valid = false
				break
			}
		}
		if valid {
			uf.union(p1, p2)
			res[i] = true
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
