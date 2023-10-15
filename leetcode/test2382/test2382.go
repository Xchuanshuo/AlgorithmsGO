package test2382

/**
 * @Description https://leetcode.cn/problems/maximum-segment-sum-after-removals
 * idea: 倒序添加 + 并查集
 **/

func maximumSegmentSum(nums []int, removeQueries []int) []int64 {
	var n = len(nums)
	var uf = NewDefUF()
	var res = make([]int64, n)
	var sum = make([]int64, n)
	var maxS = int64(0)
	for i := n - 1; i >= 0; i-- {
		var q = removeQueries[i]
		res[i] = maxS
		var cur = int64(nums[q])
		if q-1 >= 0 {
			var pl = uf.find(q - 1)
			if sum[pl] > 0 {
				uf.union(q-1, q)
				cur += sum[pl]
			}
		}
		if q+1 < n {
			var pr = uf.find(q + 1)
			if sum[pr] > 0 {
				uf.union(q+1, q)
				cur += sum[pr]
			}
		}
		sum[uf.find(q)] = cur
		maxS = max(maxS, cur)
	}
	return res
}

func max(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
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
