package test812

import "math"

func largestTriangleArea(points [][]int) float64 {
	var n = len(points)
	var max float64 = 0
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			for k := j + 1; k < n; k++ {
				x1, y1 := float64(points[i][0]), float64(points[i][1])
				x2, y2 := float64(points[j][0]), float64(points[j][1])
				x3, y3 := float64(points[k][0]), float64(points[k][1])
				var area = 0.5 * math.Abs(x1*(y2-y3)+x2*(y3-y1)+x3*(y1-y2))
				if area > max {
					max = area
				}
			}
		}
	}
	return max
}
