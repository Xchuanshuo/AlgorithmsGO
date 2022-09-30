package test765

func minSwapsCouples(row []int) int {
	var n = len(row)
	var uf = NewUF(n / 2)
	for i := 0; i < n; i += 2 {
		uf.union(row[i]/2, row[i+1]/2)
	}
	return n/2 - uf.cnt
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
