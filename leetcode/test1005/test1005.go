package test1005

import (
	"math"
	"sort"
)

func largestSumAfterKNegations(nums []int, k int) int {
	sort.Ints(nums)
	min, sum := math.MaxInt,0
	for _, v := range nums {
		if v < 0 && k > 0 {
			v = -v
			k--
		}
		sum += v
		min = minI(min, v)
	}
	if k&1 == 1 {
		sum -= 2 * min
	}
	return sum
}

func minI(x, y int) int {
	if x < y {
		return x
	}
	return y
}
