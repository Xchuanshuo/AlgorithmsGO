package test1499

import "math"

// 单调队列
func findMaxValueOfEquation(points [][]int, k int) int {
	var q = make([][]int, 0)
	var res = math.MinInt32
	for _, p := range points {
		x, y := p[0], p[1]
		for len(q) > 0 && abs(x-q[0][1]) > k {
			q = q[1:]
		}
		if len(q) > 0 {
			res = max(res, x+y+q[0][0])
		}
		for len(q) > 0 && y-x >= q[len(q)-1][0] {
			q = q[0 : len(q)-1]
		}
		q = append(q, []int{y - x, x})
	}
	return res
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
