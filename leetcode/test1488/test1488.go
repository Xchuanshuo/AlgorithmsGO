package test1488

/**
 * @Description https://leetcode.cn/problems/avoid-flood-in-the-city/
 * idea: 优先队列 每次抽干湖泊时考虑先抽空后面【先下雨】的湖泊
		 1.要知道哪个湖泊先下雨，可以预先从后往前处理，w保存当前湖泊出现的下标，取该下标即为最先出现的。
         2.小顶堆，对所有湖泊按当前位置的w[i]进行排序.
		 3.入堆，r>0,即当前下雨的湖泊;出堆，r==0, 即选择一个后续最新出现的湖泊进行抽干
 **/

import q "github.com/emirpasic/gods/queues/priorityqueue"

func avoidFlood(rains []int) []int {
	var n = len(rains)
	var mp = make(map[int]int)
	var w = make([]int, n)
	for i := n - 1; i >= 0; i-- {
		if rains[i] == 0 {
			continue
		}
		if v, exist := mp[rains[i]]; exist {
			w[i] = v
		} else {
			w[i] = n
		}
		mp[rains[i]] = i
	}
	var pq = q.NewWith(func(a, b interface{}) int {
		ia, ib := a.(int), b.(int)
		return w[ia] - w[ib]
	})
	var res = make([]int, n)
	mp = make(map[int]int)
	for i, r := range rains {
		res[i] = 1
		if r > 0 {
			if mp[r] > 0 { // 发生洪水
				return []int{}
			}
			mp[r]++
			pq.Enqueue(i)
			res[i] = -1
		} else {
			if !pq.Empty() { // 选择最近的湖泊抽干
				tv, _ := pq.Dequeue()
				v := rains[tv.(int)]
				mp[v]--
				res[i] = v
			}
		}
	}
	return res
}
