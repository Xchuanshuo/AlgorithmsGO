package lcwt113

import (
	"github.com/emirpasic/gods/queues/priorityqueue"
)

func minLengthAfterRemovals1(nums []int) int {
	var mp = make(map[int]int)
	for _, v := range nums {
		mp[v]++
	}
	var pq = priorityqueue.NewWith(func(a, b interface{}) int {
		return b.(int) - a.(int)
	})
	for _, v := range mp {
		pq.Enqueue(v)
	}
	for pq.Size() >= 2 {
		t1, _ := pq.Dequeue()
		t2, _ := pq.Dequeue()
		v1, v2 := t1.(int), t2.(int)
		v1--
		v2--
		if v1 > 0 {
			pq.Enqueue(v1)
		}
		if v2 > 0 {
			pq.Enqueue(v2)
		}
	}
	var res = 0
	for !pq.Empty() {
		t1, _ := pq.Dequeue()
		res += t1.(int)
	}
	return res
}
