package test2092

import "sort"

func findAllPeople(n int, meetings [][]int, firstPerson int) []int {
	var mp = make(map[int][]int)
	var arr = make([]int, 0)
	for i, m := range meetings {
		if len(mp[m[2]]) == 0 {
			arr = append(arr, m[2])
		}
		mp[m[2]] = append(mp[m[2]], i)
	}
	sort.Ints(arr)
	var set = make(map[int]bool)
	set[firstPerson] = true
	set[0] = true
	for _, a := range arr {
		var uf = NewDefUF()
		for _, v := range mp[a] {
			x, y := meetings[v][0], meetings[v][1]
			uf.union(x, y)
		}
		var g = make(map[int]bool)
		for _, v := range mp[a] {
			x, y := meetings[v][0], meetings[v][1]
			if set[x] || set[y] {
				g[uf.find(x)] = true
			}
		}
		for _, v := range mp[a] {
			x, y := meetings[v][0], meetings[v][1]
			if g[uf.find(x)] {
				set[x] = true
				set[y] = true
			}
		}
	}
	var res = make([]int, 0)
	for i := 0; i < n; i++ {
		if set[i] {
			res = append(res, i)
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
