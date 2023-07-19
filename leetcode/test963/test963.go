package test963

import (
	"math"
	"sort"
)

/**
 * @Description https://leetcode.cn/problems/minimum-area-rectangle-ii/
 * idea: 可以围成矩形 -> 1.两对坐标中点相等 2.且两点的直线距离相等
		 这样的坐标对有多个，需要两两比较 算最小面积即可.
		 整体时间复杂度O(n^2)
 **/

func minAreaFreeRect(points [][]int) float64 {
	var n = len(points)
	var set = make(map[T][][]int)
	var res = math.MaxFloat64
	for i := 0; i < n; i++ {
		x1, y1 := points[i][0], points[i][1]
		for j := i + 1; j < n; j++ {
			x2, y2 := points[j][0], points[j][1]
			var x = x1 + x2
			var y = y1 + y2
			var d = int64((x1-x2)*(x1-x2) + (y1-y2)*(y1-y2))
			set[T{x, y, d}] = append(set[T{x, y, d}], []int{i, j})
		}
	}
	for _, ps := range set {
		if len(ps) < 2 {
			continue
		}
		for i := 0; i < len(ps); i++ {
			p1, p2 := ps[i][0], ps[i][1]
			for j := i + 1; j < len(ps); j++ {
				p3, p4 := ps[j][0], ps[j][1]
				res = min(res, calcArea([][]int{points[p1], points[p2], points[p3], points[p4]}))
			}
		}
	}
	if res == math.MaxFloat64 {
		return 0
	}
	return res
}

func calcArea(points [][]int) float64 {
	var ds = make([]float64, 0)
	x1, y1 := points[0][0], points[0][1]
	for j := 1; j < len(points); j++ {
		x2, y2 := points[j][0], points[j][1]
		ds = append(ds, calD(x1, y1, x2, y2))
	}
	sort.Float64s(ds)
	return ds[0] * ds[1]
}

func calD(x1, y1, x2, y2 int) float64 {
	return math.Sqrt(float64((x1-x2)*(x1-x2) + (y1-y2)*(y1-y2)))
}

type T struct {
	x, y int
	d    int64
}

func min(a, b float64) float64 {
	if a < b {
		return a
	}
	return b
}
