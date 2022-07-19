package test587

import "sort"

// 凸包求解 Andrew算法
func outerTrees(trees [][]int) [][]int {
	sort.Slice(trees, func(i, j int) bool {
		if trees[i][0] == trees[j][0] {
			return trees[i][1] < trees[j][1]
		}
		return trees[i][0] < trees[j][0]
	})
	var n = len(trees)
	var st = make([]int, n+10)
	var visited = make([]bool, n+10)
	var p = 1
	for i := 1; i < n; i++ {
		var c = trees[i]
		for p >= 2 {
			a, b := trees[st[p-1]], trees[st[p]]
			if getV(a, b, c) >= 0 {
				break
			}
			visited[st[p]] = false
			p--
		}
		p++
		st[p] = i
		visited[i] = true
	}
	var size = p
	for i := n - 1; i >= 0; i-- {
		if visited[i] {
			continue
		}
		var c = trees[i]
		for p > size {
			a, b := trees[st[p-1]], trees[st[p]]
			if getV(a, b, c) >= 0 {
				break
			}
			p--
		}
		p++
		st[p] = i
		visited[i] = true
	}
	var res = make([][]int, 0)
	for i := 1; i < p; i++ {
		res = append(res, trees[st[i]])
	}
	return res
}

func sub(a, b []int) []int {
	return []int{a[0] - b[0], a[1] - b[1]}
}

// 向量的叉乘
// =0:三个点在一条直线, >0:在逆时针方向, <0:在顺时针方向
func getV(a, b, c []int) int {
	return cross(sub(b, a), sub(c, a))
}

func cross(a, b []int) int {
	return a[0]*b[1] - a[1]*b[0]
}
