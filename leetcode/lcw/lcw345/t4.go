package lcw345

func countCompleteComponents(n int, edges [][]int) int {
	var uf = NewUF(make([]int, n))
	var t = make([][]bool, n)
	for i := 0; i < n; i++ {
		t[i] = make([]bool, n)
	}
	for _, e := range edges {
		uf.union(e[0], e[1])
		t[e[0]][e[1]] = true
		t[e[1]][e[0]] = true
	}
	var mp = make(map[int][]int)
	for i := 0; i < n; i++ {
		var p = uf.find(i)
		mp[p] = append(mp[p], i)
	}
	var res = 0
	for _, vs := range mp {
		var valid = true
		for _, v1 := range vs {
			for _, v2 := range vs {
				if v1 != v2 && !t[v1][v2] {
					valid = false
					break
				}
			}
			if !valid {
				break
			}
		}
		if valid {
			res++
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
	var parent = uf.parent
	for parent[p] != p {
		parent[p] = parent[parent[p]]
		p = parent[p]
	}
	return p
}
