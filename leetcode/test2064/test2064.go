package test2064

import "math"

func minimizedMaximum(n int, quantities []int) int {
	var check = func(x int) bool {
		var cnt = 0
		for _, q := range quantities {
			cnt += (q + x - 1) / x
		}
		return cnt <= n
	}
	l, r := 1, math.MaxInt32
	for l <= r {
		var mid = l + (r-l)/2
		if !check(mid) {
			l = mid + 1
		} else {
			if mid == 1 || !check(mid-1) {
				return mid
			}
			r = mid - 1
		}
	}
	return 0
}
