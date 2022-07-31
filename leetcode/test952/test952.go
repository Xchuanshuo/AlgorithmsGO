package main

import (
	"fmt"
	"math"
)

func largestComponentSize(nums []int) int {
	var uf = NewUF(nums)
	var resMap = make(map[int]int)
	for _, a := range nums {
		var root = uf.find(a)
		resMap[root]++
	}
	var max = 0
	for _, v := range resMap {
		if v > max {
			max = v
		}
	}
	return max
}

type UF struct {
	parent map[int]int
	sz     map[int]int
}

func NewUF(arr []int) *UF {
	var parent = make(map[int]int)
	var sz = make(map[int]int)
	var set = make(map[int]bool)
	for _, a := range arr {
		set[a] = true
		parent[a] = a
		sz[a] = 1
	}
	var uf = &UF{parent: parent, sz: sz}
	for _, a := range arr {
		var cols = getCollection(a)
		for _, c := range cols {
			uf.union(a, c)
		}
	}
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
	if uf.parent[p] == 0 {
		return p
	}
	var parent = uf.parent
	for parent[p] != p {
		parent[p] = parent[parent[p]]
		p = parent[p]
	}
	return p
}

func getCollection(a int) []int {
	var sq = int(math.Sqrt(float64(a)))
	// a = b * c
	// a = sqrt(a)*sqrt(a)
	// 20 = 4.xx 10
	var res = make([]int, 0)
	for i := 2; i <= sq; i++ {
		if a%i == 0 {
			res = append(res, i)
			res = append(res, a/i)
		}
	}
	return res
}

func main() {
	var arr = []int{2, 3, 6, 7, 4, 12, 21, 39}
	var res = largestComponentSize(arr)
	fmt.Println(res)
}
