package test1439

import (
	"github.com/emirpasic/gods/queues/priorityqueue"
)

func kthSmallest(g [][]int, k int) int {
	var pre = []int{0}
	for _, row := range g {
		pre = kthArr(row, pre, k)
	}
	return pre[k-1]
}

// 给定两个有序数组，求最小的k个数对值
func kthArr(a []int, b []int, k int) []int {
	var pq = priorityqueue.NewWith(func(a, b interface{}) int {
		o1, o2 := a.([]int), b.([]int)
		return o1[0] - o2[0]
	})
	m, n := len(a), len(b)
	var visited = make([][]bool, m)
	for i := 0; i < m; i++ {
		visited[i] = make([]bool, n)
	}
	pq.Enqueue([]int{a[0] + b[0], 0, 0})
	visited[0][0] = true
	var res = make([]int, 0)
	for k >= 0 && !pq.Empty() {
		t, _ := pq.Dequeue()
		cur := t.([]int)
		res = append(res, cur[0])
		if cur[1]+1 < m && !visited[cur[1]+1][cur[2]] {
			pq.Enqueue([]int{a[cur[1]+1] + b[cur[2]], cur[1] + 1, cur[2]})
			visited[cur[1]+1][cur[2]] = true
		}
		if cur[2]+1 < n && !visited[cur[1]][cur[2]+1] {
			pq.Enqueue([]int{a[cur[1]] + b[cur[2]+1], cur[1], cur[2] + 1})
			visited[cur[1]][cur[2]+1] = true
		}
		k--
	}
	return res
}
