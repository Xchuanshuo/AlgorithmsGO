package main

import (
	"fmt"
	"github.com/emirpasic/gods/queues/priorityqueue"
)

/**
 * @Description https://leetcode.cn/contest/weekly-contest-386/problems/earliest-second-to-mark-indices-ii/
 * idea: 二分 + 逆序反悔贪心;
		1.操作3可以任选i标记为0，所以可以将cnt【标记为0的剩余次数】进行保存
		2.操作2可以将nums[changeIndices[i]]设置0，所以若指定位置不为0，可以重置
		3.操作1可以任选nums[i]减1，待设置0的数可以放入此处

		贪心: 当前nums是最后一个位置时考虑优先进行操作2，操作后将操作的目标数放入最小堆；
			若后续不存在操作1剩余次数则从堆顶拿出nums次数最小的，比当前小则进行反悔贪心
 **/

func earliestSecondToMarkIndices(nums []int, changeIndices []int) int {
	var n = len(nums)
	var m = len(changeIndices)
	var cal = func(i int) bool {
		var pq = priorityqueue.NewWith(func(a, b interface{}) int {
			va, vb := a.(int), b.(int)
			return nums[va] - nums[vb]
		})
		var firstT = make([]int, n)
		for j := 0; j < n; j++ {
			firstT[j] = -1
		}
		for j := 0; j <= i; j++ {
			var p = changeIndices[j] - 1
			if firstT[p] < 0 {
				firstT[p] = j
			}
		}
		var in = make([]bool, n)
		var cnt = 0
		for j := i; j >= 0; j-- {
			var numP = changeIndices[j] - 1
			// 最后一次出现的位置
			if firstT[numP] == j && nums[numP] > 0 {
				if cnt == 0 && pq.Size() == 0 {
					cnt++
					continue
				}
				if cnt == 0 {
					tp, _ := pq.Peek()
					var top = tp.(int)
					if nums[numP] > nums[top] {
						cnt += 2
						pq.Dequeue()
						in[top] = false
					} else {
						cnt++
						continue
					}
				}
				cnt--
				in[numP] = true
				pq.Enqueue(numP)
			} else {
				cnt++
			}
		}
		for j, v := range in {
			if !v {
				cnt -= nums[j] + 1
			}
		}
		return cnt >= 0
	}
	l, r := 0, m-1
	for l <= r {
		var mid = l + (r-l)/2
		if !cal(mid) {
			l = mid + 1
		} else {
			if mid == 0 || !cal(mid-1) {
				return mid + 1
			}
			r = mid - 1
		}
	}
	return -1
}

func main() {
	var nums = []int{0, 3}
	var changes = []int{1, 1, 2, 1, 2, 2, 2, 2, 2}
	var res = earliestSecondToMarkIndices(nums, changes)
	fmt.Println(res)
}
