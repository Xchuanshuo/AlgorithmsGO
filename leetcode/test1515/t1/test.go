package t1

import "math"

var dirs = [][]float64{{0, -1.0}, {-1.0, 0}, {1.0, 0}, {0, 1.0}}

func getMinDistSum(ps [][]int) float64 {
	var n = len(ps)
	var calDis = func(tx, ty float64) float64 {
		var total = 0.0
		for _, p := range ps {
			total += math.Sqrt((float64(p[0])-tx)*(float64(p[0])-tx) +
				(float64(p[1])-ty)*(float64(p[1])-ty))
		}
		return total
	}
	var tx, ty float64
	for _, p := range ps {
		tx += float64(p[0])
		ty += float64(p[0])
	}
	tx, ty = tx/float64(n), ty/float64(n)
	var ep = 1e-7
	var step = 1.0
	for step > ep {
		var found = false
		for _, d := range dirs {
			nx, ny := tx+step*d[0], ty+step*d[1]
			if calDis(nx, ny) < calDis(tx, ty) {
				found = true
				tx, ty = nx, ny
			}
		}
		if !found {
			step *= 0.5
		}
	}
	return calDis(tx, ty)
}
