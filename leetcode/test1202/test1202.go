package test1202

import "sort"

/**
 * @Description https://leetcode.cn/problems/smallest-string-with-swaps/
 * idea: 由于对于pairs中的索引能任意交换, 所以使用并查集将能交换的索引联通起来
		1.使用并查集联通所有能交换的位置
		2.将每个连通分量的所有索引按在字符串中对应字符大小 从小到大排序
		3.将每个连通分量的索引列表使用排序后的替换，即可得出最小字符串
 **/

func smallestStringWithSwaps(s string, pairs [][]int) string {
	var uf = NewDefUF()
	for _, p := range pairs {
		uf.union(p[0], p[1])
	}
	var idMap = make(map[int][]int)
	for k := range uf.parent {
		var p = uf.find(k)
		idMap[p] = append(idMap[p], k)
	}
	var bytes = []byte(s)
	for _, ids := range idMap {
		var t = make([]int, len(ids))
		copy(t, ids)
		sort.Ints(t)
		sort.Slice(ids, func(i, j int) bool {
			return s[ids[i]] < s[ids[j]]
		})
		for i, idx := range t {
			bytes[idx] = s[ids[i]]
		}
	}
	return string(bytes)
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
