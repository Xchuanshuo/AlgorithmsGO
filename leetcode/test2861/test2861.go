package test2861

/**
 * @Description https://leetcode.cn/problems/maximum-number-of-alloys/
 * idea: 二分
 **/

func maxNumberOfAlloys(n int, k int, budget int, composition [][]int, stock []int, cost []int) int {
	var check = func(id, cnt int) bool {
		var total = budget
		var t = composition[id]
		for i := 0; i < len(t); i++ {
			var target = t[i] * cnt
			if stock[i] >= target {
				continue
			}
			total -= (target - stock[i]) * cost[i]
		}
		return total >= 0
	}
	var cal = func(id int) int {
		var INF = int(1e9)
		l, r := 0, INF
		for l <= r {
			var mid = l + (r-l)/2
			if !check(id, mid) {
				r = mid - 1
			} else {
				if mid == INF || !check(id, mid+1) {
					return mid
				}
				l = mid + 1
			}
		}
		return 0
	}
	var res = 0
	for i := 0; i < k; i++ {
		res = max(res, cal(i))
	}
	return res
}
