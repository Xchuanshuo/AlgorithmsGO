package lcwt102

import (
	"github.com/emirpasic/gods/queues/priorityqueue"
	"math"
)

type Graph struct {
	g [][]int
	w [][]int
	n int
}

func Constructor(n int, edges [][]int) Graph {
	var g = make([][]int, n)
	var w = make([][]int, n)
	for i := 0; i < n; i++ {
		w[i] = make([]int, n)
	}
	for _, e := range edges {
		g[e[0]] = append(g[e[0]], e[1])
		w[e[0]][e[1]] = e[2]
	}
	return Graph{g, w, n}
}

func (this *Graph) AddEdge(e []int) {
	this.g[e[0]] = append(this.g[e[0]], e[1])
	this.w[e[0]][e[1]] = e[2]
}

func (this *Graph) ShortestPath(node1 int, node2 int) int {
	return this.calc(node1, node2)
}

// dijkstra 最短路径
func (this *Graph) calc(s, t int) int {
	var pq = priorityqueue.NewWith(func(a, b interface{}) int {
		o1, o2 := a.([]int), b.([]int)
		return o1[1] - o2[1]
	})
	var dst = make([]int, this.n)
	for i := 0; i < this.n; i++ {
		dst[i] = math.MaxInt32
	}
	dst[s] = 0
	pq.Enqueue([]int{s, dst[s]})
	for !pq.Empty() {
		tv, _ := pq.Dequeue()
		node := tv.([]int)
		v := node[0]
		for _, w := range this.g[v] {
			if dst[w] > dst[v]+this.w[v][w] {
				dst[w] = dst[v] + this.w[v][w]
				pq.Enqueue([]int{w, dst[w]})
			}
		}
	}
	if dst[t] == math.MaxInt32 {
		return -1
	}
	return dst[t]
}
