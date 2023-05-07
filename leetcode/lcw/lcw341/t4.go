package lcw341

import (
	"github.com/emirpasic/gods/queues/priorityqueue"
	"math"
)

func minimumTotalPrice(n int, edges [][]int, price []int, trips [][]int) int {
	var g = make([][]int, n)
	for _, e := range edges {
		g[e[0]] = append(g[e[0]], e[1])
		g[e[1]] = append(g[e[1]], e[0])
	}
	var getPath = func(s int, t int) []int {
		var pq = priorityqueue.NewWith(func(a, b interface{}) int {
			o1, o2 := a.([]int), b.([]int)
			return o1[1] - o2[1]
		})
		var visited = make([]bool, n)
		var dst = make([]int, n)
		var pre = make([]int, n)
		for i := 0; i < n; i++ {
			pre[i] = -1
			dst[i] = math.MaxInt32
		}
		pre[s] = s
		dst[s] = price[s]
		pq.Enqueue([]int{s, dst[s]})
		for !pq.Empty() {
			tv, _ := pq.Dequeue()
			node := tv.([]int)
			v := node[0]
			if visited[v] {
				continue
			}
			visited[v] = true
			for _, w := range g[v] {
				if !visited[w] && dst[w] > dst[v]+price[w] {
					dst[w] = dst[v] + price[w]
					pre[w] = v
					pq.Enqueue([]int{w, dst[w]})
				}
			}
		}
		var path = []int{t}
		var cur = t
		for cur != s {
			cur = pre[cur]
			path = append(path, cur)
		}
		return path
	}
	var cnt = make([]int, n)
	for _, trip := range trips {
		var path = getPath(trip[0], trip[1])
		for _, p := range path {
			cnt[p]++
		}
	}
	var dfs func(int, int) (int, int)
	dfs = func(x, fa int) (int, int) {
		noSub := price[x] * cnt[x]
		sub := noSub / 2
		for _, nxt := range g[x] {
			if nxt != fa {
				ns, s := dfs(nxt, x)
				noSub += min(ns, s)
				sub += ns
			}
		}
		return noSub, sub
	}
	nh, h := dfs(0, -1)
	return min(nh, h)
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
