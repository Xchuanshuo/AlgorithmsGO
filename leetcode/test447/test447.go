package test447

import "math"

func numberOfBoomerangs(points [][]int) int {
	var getDis = func(i, j int) float64 {
		x1, y1 := points[i][0], points[i][1]
		x2, y2 := points[j][0], points[j][1]
		var v = (x1-x2)*(x1-x2) + (y1-y2)*(y1-y2)
		return math.Sqrt(float64(v))
	}
	var res = 0
	for i := range points {
		var mp = make(map[float64]int)
		for j := range points {
			if i == j {
				continue
			}
			var d = getDis(i, j)
			mp[d]++
		}
		for _, v := range mp {
			if v >= 2 {
				res += v * (v - 1)
			}
		}
	}
	return res
}
