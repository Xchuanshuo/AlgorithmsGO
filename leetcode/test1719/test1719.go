package test1719

import "math"

func checkWays(pairs [][]int) int {
	var g = make(map[int]map[int]bool)
	for _, p := range pairs {
		x, y := p[0], p[1]
		if g[x] == nil {
			g[x] = make(map[int]bool)
		}
		if g[y] == nil {
			g[y] = make(map[int]bool)
		}
		g[x][y] = true
		g[y][x] = true
	}
	var root = -1
	for k, v := range g {
		if len(v) == len(g)-1 {
			root = k
			break
		}
	}
	if root == -1 {
		return 0
	}
	var res = 1
	for cur, adj := range g {
		if cur == root {
			continue
		}
		curDegree := len(adj)
		parent, parentDegree := -1, math.MaxInt32
		for next, _ := range g[cur] {
			if len(g[next]) >= curDegree && len(g[next]) < parentDegree {
				parent = next
				parentDegree = len(g[next])
			}
		}
		if parent == -1 {
			return 0
		}
		for v, _ := range adj {
			if v != parent && !g[parent][v] {
				return 0
			}
		}
		if curDegree == parentDegree {
			res = 2
		}
	}
	return res
}
