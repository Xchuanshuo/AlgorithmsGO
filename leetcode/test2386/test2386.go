package test2386

import (
	"github.com/emirpasic/gods/queues/priorityqueue"
	"sort"
)

/**
 * @Description https://leetcode.cn/problems/find-the-k-sum-of-an-array/
 * idea: 排序 + 贪心 + 大顶堆
		1.最大和一定是所有正整数相加
		2.第二大的和为 max(最大和-最小正整数, 最大和+最大负整数)
		 实际可以转换为 最大和 - |最小整数|(绝对值)
		3.依次类推，将每个上一次操作当作当前最大和，次大和一定还是按照这个规律进行计算
		4.使用堆护【当前最大和】
 **/

func kSum(nums []int, k int) int64 {
	var sum = int64(0)
	var n = int64(len(nums))
	for i, v := range nums {
		if v < 0 {
			nums[i] = -nums[i]
		} else {
			sum += int64(nums[i])
		}
	}
	sort.Ints(nums)
	var pq = priorityqueue.NewWith(func(a, b interface{}) int {
		av, bv := a.([]int64), b.([]int64)
		return int(bv[0] - av[0])
	})
	pq.Enqueue([]int64{sum, 0})
	for i := 1; i < k; i++ {
		t, _ := pq.Dequeue()
		var top = t.([]int64)
		var cur = top[1]
		if cur < n {
			// 直接减去当前
			pq.Enqueue([]int64{top[0] - int64(nums[cur]), cur + 1})
			// 替换前一个减去的数为
			if cur >= 1 {
				pq.Enqueue([]int64{top[0] - int64(nums[cur]-nums[cur-1]), cur + 1})
			}
		}
	}
	t, _ := pq.Peek()
	var top = t.([]int64)
	return top[0]
}