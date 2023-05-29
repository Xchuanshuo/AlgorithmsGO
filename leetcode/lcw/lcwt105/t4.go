package lcwt105

import (
	"math"
)

/**
 * idea: 1.分解质因数, 记录质因数对应的元素位置
		 2.将所有有相同质因数的元素进行合并，如果任意两个元素都能走到，
		  则最终并查集一定为元素个数
 **/

func canTraverseAllPairs(nums []int) bool {
	var n = len(nums)
	var mp = make(map[int]int)
	var uf = NewUF(nums)
	for i := 0; i < n; i++ {
		for _, v := range f(nums[i]) {
			if v == 1 {
				continue
			}
			if _, exist := mp[v]; !exist {
				mp[v] = i
			} else {
				uf.union(i, mp[v])
			}
		}
	}
	return uf.getCnt(0) == n
}

// 分解质因数
var f = func(n int) []int {
	var facs = make([]int, 0)
	for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
		if n%i == 0 {
			for n%i == 0 {
				n /= i
				facs = append(facs, i)
			}
		}
	}
	if n > 1 {
		facs = append(facs, n)
	}
	return facs
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

func (uf *UF) getCnt(p int) int {
	return uf.sz[uf.find(p)]
}
