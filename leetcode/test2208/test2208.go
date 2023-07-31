package test2208

import (
	"github.com/emirpasic/gods/queues/priorityqueue"
)

func halveArray(nums []int) int {
	var pq = priorityqueue.NewWith(func(a, b interface{}) int {
		ia, ib := a.(float64), b.(float64)
		if ib == ia {
			return 0
		}
		if ib < ia {
			return -1
		}
		return 1
	})
	var sum = 0.0
	for _, v := range nums {
		sum += float64(v) / 2.0
		pq.Enqueue(float64(v))
	}
	var cnt = 0
	for sum > 0 {
		t, _ := pq.Dequeue()
		v := t.(float64)
		sum -= float64(v) / 2.0
		pq.Enqueue(float64(v) / 2.0)
		cnt++
	}
	return cnt
}
