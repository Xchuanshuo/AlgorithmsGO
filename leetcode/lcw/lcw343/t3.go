package lcw343

import (
	"github.com/emirpasic/gods/queues/priorityqueue"
	"math"
)

func minimumCost(s []int, t []int, sp [][]int) int {
	var pq = priorityqueue.NewWith(func(a, b interface{}) int {
		o1, o2 := a.([]int), b.([]int)
		return o1[1] - o2[1]
	})
	_, n := t[0]+1, t[1]+1
	var dst = make(map[int]int)
	for _, p := range sp {
		var d1 = p[0]*n + p[1]
		dst[d1] = math.MaxInt32
		var d2 = p[2]*n + p[3]
		dst[d2] = math.MaxInt32
		if abs(p[0]-p[2])+abs(p[1]-p[3]) < p[4] {
			p[4] = abs(p[0]-p[2]) + abs(p[1]-p[3])
		}
	}
	var e = t[0]*n + t[1]
	dst[e] = abs(s[0]-t[0]) + abs(s[1]-t[1])
	dst[s[0]*n+s[1]] = 0
	pq.Enqueue([]int{s[0]*n + s[1], 0})
	for !pq.Empty() {
		tv, _ := pq.Dequeue()
		node := tv.([]int)
		v := node[0]
		x, y := v/n, v%n
		for _, p := range sp {
			var d1 = p[0]*n + p[1]
			var d2 = p[2]*n + p[3]
			if v == d1 && dst[d2] > dst[d1]+p[4] {
				dst[d2] = dst[d1] + p[4]
				pq.Enqueue([]int{d2, dst[d2]})
			}
			if v != d1 && dst[d1] > dst[v]+abs(x-p[0])+abs(y-p[1]) {
				dst[d1] = dst[v] + abs(x-p[0]) + abs(y-p[1])
				pq.Enqueue([]int{d1, dst[d1]})
			}
		}
		if dst[e] > dst[v]+abs(x-t[0])+abs(y-t[1]) {
			dst[e] = dst[v] + abs(x-t[0]) + abs(y-t[1])
		}
	}
	return dst[e]
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
