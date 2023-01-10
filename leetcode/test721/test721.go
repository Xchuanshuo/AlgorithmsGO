package test721

import "sort"

func accountsMerge(accounts [][]string) [][]string {
	var emailMap = make(map[string]int)
	var n = len(accounts)
	var uf = NewUF(make([]int, n))
	for k, account := range accounts {
		for i := 1; i < len(account); i++ {
			var email = account[i]
			if v, exist := emailMap[email]; exist {
				uf.union(v, k)
			} else {
				emailMap[email] = k
			}
		}
	}
	var userMap = make(map[int][]int)
	for i := 0; i < n; i++ {
		var p = uf.find(i)
		userMap[p] = append(userMap[p], i)
	}
	var res = make([][]string, 0)
	for k, vs := range userMap {
		var cur = []string{accounts[k][0]}
		var t = make([]string, 0)
		var visited = make(map[string]bool)
		for _, v := range vs {
			for i := 1; i < len(accounts[v]); i++ {
				var email = accounts[v][i]
				if visited[email] {
					continue
				}
				visited[email] = true
				t = append(t, email)
			}
		}
		sort.Strings(t)
		cur = append(cur, t...)
		res = append(res, cur)
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
