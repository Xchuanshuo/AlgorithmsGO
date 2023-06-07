package test2699

import (
	"github.com/emirpasic/gods/queues/priorityqueue"
)

/**
 * @Description https://leetcode.cn/problems/modify-graph-edge-weights/
 * idea: 1.首先不考虑负权边，跑一遍最短路
		   若dst[e]<target，没法添加新的边使得最短路变大，说明不存在修改方案;
		   若dst[e]=target, 直接将负权边修改为最大值即为答案
	     2.若dst[e]>target，则依次修改每条边，若当前边修改为1后，使得dst[e]<=target,
		  说明当前边为最短路经过的边，直接将当前边修改为target-dst[e], 其它边修改为最大即可
		 3.依次进行2后，最终最短路dst[e]>target，也说明无解
		 整体时间复杂度 O(m*m*logM), 其中m为边的数量
 **/

var INF = 2 * int(1e9)

func modifiedGraphEdges(n int, edges [][]int, s int, e int, target int) [][]int {
	var updates = make([]int, 0)
	var g = make(map[int][][]int)
	for i, e := range edges {
		if e[2] == -1 {
			updates = append(updates, i)
		} else {
			g[e[0]] = append(g[e[0]], []int{e[1], e[2]})
			g[e[1]] = append(g[e[1]], []int{e[0], e[2]})
		}
	}
	dst := dijkstra(n, g, s, e)
	if dst < target {
		return [][]int{}
	}
	if dst == target {
		for _, idx := range updates {
			if edges[idx][2] == -1 {
				edges[idx][2] = INF
			}
		}
		return edges
	}
	var valid = false
	for _, idx := range updates {
		var edge = edges[idx]
		edge[2] = 1
		g[edge[0]] = append(g[edge[0]], []int{edge[1], edge[2]})
		g[edge[1]] = append(g[edge[1]], []int{edge[0], edge[2]})
		dst = dijkstra(n, g, s, e)
		if dst <= target {
			edge[2] += target - dst
			valid = true
			break
		}
	}
	if valid {
		for _, idx := range updates {
			if edges[idx][2] == -1 {
				edges[idx][2] = INF
			}
		}
		return edges
	}
	return [][]int{}
}

func dijkstra(n int, g map[int][][]int, s, t int) int {
	var pq = priorityqueue.NewWith(func(a, b interface{}) int {
		o1, o2 := a.([]int), b.([]int)
		return o1[1] - o2[1]
	})
	var dst = make([]int, n)
	for i := 0; i < n; i++ {
		dst[i] = INF
	}
	dst[s] = 0
	pq.Enqueue([]int{s, dst[s]})
	for !pq.Empty() {
		tv, _ := pq.Dequeue()
		node := tv.([]int)
		v := node[0]
		for _, cur := range g[v] {
			nxt, w := cur[0], cur[1]
			if dst[nxt] > dst[v]+w {
				dst[nxt] = dst[v] + w
				pq.Enqueue([]int{nxt, dst[nxt]})
			}
		}
	}
	return dst[t]
}
