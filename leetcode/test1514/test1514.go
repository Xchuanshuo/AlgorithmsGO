package test1514

import "github.com/emirpasic/gods/queues/priorityqueue"

/**
 * @Description https://leetcode.cn/problems/path-with-maximum-probability
 * idea: 常规dijkstra变种题，由于路径的权值在[0,1]之间，概率为乘法运算，所以增加新的边后不会使概率变大，
		优先不考虑添加新的边，与dijkstra一致，添加多余的边不会使权值变小。本题满足该性质后，使用dijkstra
		进行计算，优先队列保存当前概率最大的路径，比较大小时也是看添加新边后概率是否变大。
 **/

func maxProbability(n int, edges [][]int, probs []float64, start int, end int) float64 {
	var g = make(map[int][]T)
	for i, e := range edges {
		g[e[0]] = append(g[e[0]], T{e[1], probs[i]})
		g[e[1]] = append(g[e[1]], T{e[0], probs[i]})
	}
	return dijkstra(n, g, start, end)
}

func dijkstra(n int, g map[int][]T, s, t int) float64 {
	var pq = priorityqueue.NewWith(func(a, b interface{}) int {
		o1, o2 := a.(T), b.(T)
		if o1.p < o2.p {
			return 1
		}
		return -1
	})
	var dst = make([]float64, n)
	dst[s] = 1.0
	pq.Enqueue(T{s, dst[s]})
	for !pq.Empty() {
		tv, _ := pq.Dequeue()
		node := tv.(T)
		v := node.v
		for _, cur := range g[v] {
			nxt, w := cur.v, cur.p
			if dst[nxt] < dst[v]*w {
				dst[nxt] = dst[v] * w
				pq.Enqueue(T{nxt, dst[nxt]})
			}
		}
	}
	return dst[t]
}

type T struct {
	v int
	p float64
}
