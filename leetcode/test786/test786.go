package test786

import "github.com/emirpasic/gods/queues/priorityqueue"

func kthSmallestPrimeFraction(arr []int, k int) []int {
	var n = len(arr)
	var pq = priorityqueue.NewWith(func(a, b interface{}) int {
		o1, o2 := a.(T), b.(T)
		if o1.val < o2.val {
			return -1
		}
		return 1
	})
	var visited = make([][]bool, n)
	for i := 0; i < n; i++ {
		visited[i] = make([]bool, n)
	}
	pq.Enqueue(T{float64(arr[0]) / float64(arr[n-1]), 0, n - 1})
	visited[0][n-1] = true
	var res T
	for k > 0 && !pq.Empty() {
		t, _ := pq.Dequeue()
		cur := t.(T)
		res = cur
		if cur.i+1 < n && cur.i+1 < cur.j && !visited[cur.i+1][cur.j] {
			pq.Enqueue(T{float64(arr[cur.i+1]) / float64(arr[cur.j]), cur.i + 1, cur.j})
			visited[cur.i+1][cur.j] = true
		}
		if cur.j-1 >= 0 && cur.i < cur.j-1 && !visited[cur.i][cur.j-1] {
			pq.Enqueue(T{float64(arr[cur.i]) / float64(arr[cur.j-1]), cur.i, cur.j - 1})
			visited[cur.i][cur.j-1] = true
		}
		k--
	}
	return []int{arr[res.i], arr[res.j]}
}

type T struct {
	val  float64
	i, j int
}
