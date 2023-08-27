package test1579

// 贪心 + 并查集， 1.选择双向边一定优于选择单向边
//				2.两个子图不连通时，若有双向边，优先选择双向边。
//	 			若只能选择单向边时，一定不存在其它双向边使图连通

func maxNumEdgesToRemove(n int, edges [][]int) int {
	var ufa = NewDefUF()
	var ufb = NewDefUF()
	var res = 0
	for _, e := range edges {
		if e[0] != 3 {
			continue
		}
		var valid = false
		if !ufa.isConnected(e[1], e[2]) {
			ufa.union(e[1], e[2])
			valid = true
		}
		if !ufb.isConnected(e[1], e[2]) {
			ufb.union(e[1], e[2])
			valid = true
		}
		if valid {
			res++
		}
	}
	for _, e := range edges {
		if e[0] == 1 && !ufa.isConnected(e[1], e[2]) {
			ufa.union(e[1], e[2])
			res++
		} else if e[0] == 2 && !ufb.isConnected(e[1], e[2]) {
			ufb.union(e[1], e[2])
			res++
		}
	}
	if ufa.getCnt(1) != n || ufb.getCnt(1) != n {
		return -1
	}
	return len(edges) - res
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
