package test875

import "math"

func minEatingSpeed(piles []int, h int) int {
	l, r := 1, math.MaxInt32
	for l <= r {
		var mid = l + (r-l)/2
		if !canFinish(piles, h, mid) {
			l = mid + 1
		} else {
			if mid == l || !canFinish(piles, h, mid-1) {
				return mid
			}
			r = mid - 1
		}
	}
	return 0
}

func canFinish(piles []int, h, k int) bool {
	var cnt = 0
	for _, v := range piles {
		cnt += (v + k - 1) / k
	}
	return cnt <= h
}
