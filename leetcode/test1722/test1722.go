package test1722

// 求交换后最小不相等的数对数量 按下标进行分组，每个组之间可以任意交换, 使用并查集或dfs即可
func minimumHammingDistance(source []int, target []int, allowedSwaps [][]int) int {
	var n = len(source)
	var uf = NewUF(make([]int, n))
	for _, s := range allowedSwaps {
		uf.union(s[0], s[1])
	}
	var mp = make(map[int][]int)
	for i := 0; i < n; i++ {
		var p = uf.find(i)
		mp[p] = append(mp[p], i)
	}
	var res = 0
	for _, ids := range mp {
		var t = make(map[int]int)
		for _, id := range ids {
			t[source[id]]++
		}
		for _, id := range ids {
			if t[target[id]] == 0 {
				res++
			} else {
				t[target[id]]--
			}
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
