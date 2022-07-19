package test719

import (
	"math"
	"sort"
)

func smallestDistancePair(nums []int, k int) int {
	var n = len(nums)
	sort.Ints(nums)
	l, r := 0, math.MaxInt32
	var isValid func(m int) bool
	isValid = func(m int) bool {
		i, j := 0, 0
		var cnt = 0
		for j < n {
			for nums[j]-nums[i] > m {
				i++
			}
			cnt += j - i
			j++
		}
		return cnt >= k
	}
	for l <= r {
		mid := l + (r-l)/2
		if !isValid(mid) {
			l = mid + 1
		} else {
			if mid == l || !isValid(mid-1) {
				return mid
			}
			r = mid - 1
		}
	}
	return -1
}
