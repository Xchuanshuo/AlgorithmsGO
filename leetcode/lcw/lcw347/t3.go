package lcw347

import "math"

func minimumCost(s string) int64 {
	var n = len(s)
	var left = make([][]int64, n+1)
	var right = make([][]int64, n+1)
	for i := 0; i < len(left); i++ {
		left[i] = make([]int64, 2)
		right[i] = make([]int64, 2)
	}
	for i := n - 1; i >= 0; i-- {
		if s[i] == '0' {
			right[i][0] = min(right[i+1][0], right[i+1][1]+int64(n-i-1))
			right[i][1] = int64(n-i) + min(right[i+1][0], right[i+1][1]+int64(n-i-1))
		} else {
			right[i][0] = int64(n-i) + min(right[i+1][1], right[i+1][0]+int64(n-i-1))
			right[i][1] = min(right[i+1][1], right[i+1][0]+int64(n-i-1))
		}
	}
	var res int64 = math.MaxInt64
	for i := 1; i <= n; i++ {
		if s[i-1] == '0' {
			left[i][0] = min(left[i-1][0], left[i-1][1]+int64(i-1))
			left[i][1] = int64(i) + min(left[i-1][0], left[i-1][1]+int64(i-1))
		} else {
			left[i][0] = int64(i) + min(left[i-1][1], left[i-1][0]+int64(i-1))
			left[i][1] = min(left[i-1][1], left[i-1][0]+int64(i-1))
		}
		res = min(res, left[i][0]+right[i][0])
		res = min(res, left[i][1]+right[i][1])
	}
	return res
}

func min(a, b int64) int64 {
	if a < b {
		return a
	}
	return b
}
