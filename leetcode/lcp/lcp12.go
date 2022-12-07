package main

/**
 * @Description https://leetcode.cn/problems/xiao-zhang-shua-ti-ji-hua/
 * idea: 1.二分 2.贪心，过程【记录花费时间最多】的题，让小杨做
 **/

func minTime(time []int, m int) int {
	var check = func(target int) bool {
		cur, days := 0, 1
		mx, cnt := 0, 1
		for j := 0; j < len(time); j++ {
			if mx < time[j] {
				mx = time[j]
			}
			cur += time[j]
			if cur <= target {
				continue
			}
			if cnt == 1 {
				cur -= mx
				cnt--
			} else {
				days++
				cnt, mx, cur = 1, 0, 0
				j--
			}
		}
		return days <= m
	}
	var sum = 0
	for i := 0; i < len(time); i++ {
		sum += time[i]
	}
	l, r := 0, sum
	for l <= r {
		mid := l + (r-l)/2
		if !check(mid) {
			l = mid + 1
		} else {
			if mid == 0 || !check(mid-1) {
				return mid
			}
			r = mid - 1
		}
	}
	return sum
}
