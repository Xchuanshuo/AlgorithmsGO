package test2242

import "sort"

/**
 * @Description
 * idea:
 *	    1.长度为4, 以边为中间点 枚举所有边 a-x-y-b
 *		2.由于总共选取4个点，要保证最大的分数对应点不会出现重复，
 *		 对于每个顶多需要保存三个最大的邻点
 **/

func maximumScore(scores []int, edges [][]int) int {
	var graph = make(map[int][]int)
	for _, e := range edges {
		graph[e[0]] = append(graph[e[0]], e[1])
		graph[e[1]] = append(graph[e[1]], e[0])
	}
	for k, g := range graph {
		sort.Slice(g, func(i, j int) bool {
			return scores[g[i]] > scores[g[j]]
		})
		if len(g) > 3 {
			graph[k] = g[0:3]
		}
	}
	var res = -1
	for _, e := range edges {
		x, y := e[0], e[1]
		for _, a := range graph[x] {
			for _, b := range graph[y] {
				if a != b && a != y && b != x {
					res = max(res, scores[a]+scores[b]+scores[x]+scores[y])
				}
			}
		}
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
