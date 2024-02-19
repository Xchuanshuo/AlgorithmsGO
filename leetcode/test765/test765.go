package test765

/**
 * @Description https://leetcode.cn/problems/couples-holding-hands
 * idea: 并查集.
			1.给每队情侣分配同一个id
			2.计算总共有多少对不同id在一个集合中
			3.交换次数为每个单独的集合大小-1, 也可以反过来计算 总对数 - 不需要交换的对数
 **/

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
