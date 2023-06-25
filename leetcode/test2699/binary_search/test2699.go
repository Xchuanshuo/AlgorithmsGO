package binary_search

import (
	"github.com/emirpasic/gods/queues/priorityqueue"
)

/**
 * @Description https://leetcode.cn/problems/modify-graph-edge-weights/
 * idea: 由于对于最短路dst(s,e) 一条边权+1后，最短路要么+1，要么不变
		 所以考虑依次给边添加权重, 则相邻大小边权的方案满足【单调性】, 分情况讨论
 		 1.将负权边全置为1,此时最短路达到最小值 若dst[e]>target, 则不存在合法方案
 		 2.将负权边全置为target,此时最短路达到最大值 若dst[e]<target, 则不存在合法方案
		 3.此时肯定存在答案，通过二分枚举边权的变化量[1,target], 以确定的顺序依次添加权重，
		  直到找到解为止
		 整体时间复杂度 O(log(m*target)*m*logM), 其中m为边的数量
 **/

var INF = 2 * int(1e9)

func modifiedGraphEdges(n int, edges [][]int, s int, e int, target int) [][]int {
	var buildG = func(adds int64) map[int][][]int {
		var g = make(map[int][][]int)
		for _, edge := range edges {
			v1, v2, w := edge[0], edge[1], edge[2]-1
			if w < 0 {
				if adds <= 0 {
					w = 0
				} else {
					if adds >= int64(target-1) {
						w = target - 1
						adds -= int64(target - 1)
					} else {
						w = int(adds)
						adds = 0
					}
				}
			}
			g[v1] = append(g[v1], []int{v2, w + 1})
			g[v2] = append(g[v2], []int{v1, w + 1})
		}
		return g
	}
	var cnt = 0
	for _, e := range edges {
		if e[2] == -1 {
			cnt++
		}
	}
	var mx = int64(cnt*(target-1) + 1)
	if dijkstra(n, buildG(0), s, e) > target {
		return [][]int{}
	}
	if dijkstra(n, buildG(mx), s, e) < target {
		return [][]int{}
	}
	l, r := int64(0), mx
	var res = r
	for l <= r {
		var mid = l + (r-l)/2
		if dijkstra(n, buildG(mid), s, e) >= target {
			res = mid
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	for _, edge := range edges {
		var w = edge[2] - 1
		if w >= 0 {
			continue
		}
		if res <= 0 {
			w = 0
		} else {
			if res >= int64(target-1) {
				w = target - 1
				res -= int64(target - 1)
			} else {
				w = int(res)
				res = 0
			}
		}
		edge[2] = w + 1
	}
	return edges
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
