package main

import (
	"github.com/emirpasic/gods/queues/priorityqueue"
	"github.com/emirpasic/gods/utils"
)

/**
 * @Description https://leetcode.cn/problems/p0NxJO
 * idea: 贪心+小顶堆，总血量>0则不移动，<=0, 则移动前面负血量最多的
 **/

func magicTower(nums []int) int {
	var sum = 0
	for _, v := range nums {
		sum += v
	}
	if sum < 0 {
		return -1
	}
	var pq = priorityqueue.NewWith(utils.IntComparator)
	hp, cnt := 1, 0
	for _, v := range nums {
		if v < 0 {
			pq.Enqueue(v)
		}
		hp += v
		if hp <= 0 {
			x, _ := pq.Dequeue()
			xv := x.(int)
			hp -= xv
			cnt++
		}
	}
	return cnt
}
