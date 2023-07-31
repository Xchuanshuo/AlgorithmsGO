package test1353

import "sort"
import "github.com/emirpasic/gods/queues/priorityqueue"

/**
 * @Description https://leetcode.cn/problems/maximum-number-of-events-that-can-be-attended/
 * idea: 贪心 1.按结束时间排序 优先处理结束时间靠后的
			 2.连续多个结束时间相同，则统一入堆后优先处理开始时间较晚的
 **/

func maxEvents(events [][]int) int {
	sort.Slice(events, func(i, j int) bool {
		return events[i][1] < events[j][1]
	})
	var pq = priorityqueue.NewWith(func(a, b interface{}) int {
		a1, b1 := a.([]int), b.([]int)
		return b1[0] - a1[0]
	})
	var n = len(events)
	cur, res := events[n-1][1], 0
	var i = n - 1
	for i >= 0 || !pq.Empty() {
		if i >= 0 && events[i][1] < cur {
			cur = events[i][1]
		}
		for i >= 0 && events[i][1] == cur {
			pq.Enqueue(events[i])
			i--
		}
		for !pq.Empty() {
			t, _ := pq.Peek()
			tv := t.([]int)
			if i >= 0 && events[i][1] == cur {
				break
			}
			if tv[0] <= cur && cur <= tv[1] {
				res++
				cur--
			}
			pq.Dequeue()
		}
	}
	return res
}
