package test1962

import (
	"github.com/emirpasic/gods/queues/priorityqueue"
)

func minStoneSum(piles []int, k int) int {
	var pq = priorityqueue.NewWith(func(a, b interface{}) int {
		va, vb := a.(int), b.(int)
		return vb - va
	})
	for _, p := range piles {
		pq.Enqueue(p)
	}
	for i := 0; i < k; i++ {
		tv, _ := pq.Dequeue()
		v := tv.(int)
		pq.Enqueue(v - v/2)
	}
	var res = 0
	for !pq.Empty() {
		tv, _ := pq.Dequeue()
		v := tv.(int)
		res += v
	}
	return res
}
