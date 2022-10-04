package main

import (
	"github.com/emirpasic/gods/trees/redblacktree"
)

func explorationSupply(station []int, pos []int) []int {
	var mp = make(map[int]int)
	tree := redblacktree.NewWithIntComparator()
	for i, v := range station {
		mp[v] = i
		tree.Put(v, struct{}{})
	}
	var res = make([]int, 0)
	for _, p := range pos {
		l, ok1 := tree.Floor(p)
		r, ok2 := tree.Ceiling(p)
		var cur = 0
		if ok1 && ok2 {
			if abs(l.Key.(int)-p) <= abs(r.Key.(int)-p) {
				cur = mp[l.Key.(int)]
			} else {
				cur = mp[r.Key.(int)]
			}
		} else if ok1 {
			cur = mp[l.Key.(int)]
		} else {
			cur = mp[r.Key.(int)]
		}
		res = append(res, cur)
	}
	return res
}
func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

//func main() {
//	var stas = []int{2, 5, 8, 14, 17}
//	var pos = []int{1, 14, 11, 2}
//	var res = explorationSupply(stas, pos)
//	fmt.Println(res)
//}
