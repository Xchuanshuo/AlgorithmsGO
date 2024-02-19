package test2141

import (
	"sort"
)

/**
 * @Description https://leetcode.cn/problems/maximum-running-time-of-n-computers
 * idea: 二分 + 贪心。
		1.若运行时间最大为t, 则至少可以运行[0,t-1]的时间。 显然，运行的时间是满足单调性的，所以可以考虑二分答案。
		2.由于电脑需要同时运行，所以即使当前电池供电时间超过t时，也没法把多余的电量分给其它电脑；而当前供电时间不足
		  t时需要借用其它电池，从当前未选的电池借t-cur的电。此时相当于借用的电池x运行到x-(t-cur)时给当前电脑用，依次类推。
		 所以排序后贪心的电量从大到小处理即可
 **/

var INF = int64(1e14)

func maxRunTime(n int, bs []int) int64 {
	var m = len(bs)
	sort.Ints(bs)
	var check = func(t int64) bool {
		var preSub = int64(0)
		for i := m - 1; i >= m-n; i-- {
			var cur = int64(bs[i]) + preSub
			if cur >= t {
				continue
			}
			preSub = cur - t
		}
		for i := 0; i < m-n; i++ {
			preSub += int64(bs[i])
		}
		// 不满足条件时选剩余的电池，直到给所有n台供电，最后剩余电量不为负即可
		return preSub >= 0
	}
	l, r := int64(1), INF
	for l <= r {
		var mid = l + (r-l)/2
		if !check(mid) {
			r = mid - 1
		} else {
			if mid == INF || !check(mid+1) {
				return mid
			}
			l = mid + 1
		}
	}
	return 0
}
