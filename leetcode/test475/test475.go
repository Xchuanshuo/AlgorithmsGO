package test475

import "sort"

func findRadius(houses []int, heaters []int) int {
	sort.Ints(houses)
	sort.Ints(heaters)
	l, r := 0, int(1e9)
	for l <= r {
		mid := l + (r-l)/2
		if !check(houses, heaters, mid) {
			l = mid + 1
		} else {
			if mid == l || !check(houses, heaters, mid-1) {
				return mid
			}
			r = mid - 1
		}
	}
	return -1
}

func check(houses, heaters[]int, d int) bool {
	m, idx := len(houses), 0
	for _, v := range heaters {
		l, r := v - d, v + d
		for idx < m && houses[idx] >= l && houses[idx] <= r {
			idx++
		}
		if idx == m { return true }
	}
	return false
}