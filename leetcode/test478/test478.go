package test478

import "math/rand"

type Solution struct {
	cx, cy float64
	r      float64
}

func Constructor(radius float64, x_center float64, y_center float64) Solution {
	return Solution{
		cx: x_center, cy: y_center,
		r: radius,
	}
}

var bound = [][]float64{{1, 1}, {1, -1}, {-1, 1}, {-1, -1}}

func (this *Solution) RandPoint() []float64 {
	for {
		var v1 = rand.Float64() * this.r
		var v2 = rand.Float64() * this.r
		var b = bound[rand.Intn(len(bound))]
		v1, v2 = v1*b[0], v2*b[1]
		if v1*v1+v2*v2 <= this.r*this.r {
			return []float64{v1 + this.cx, v2 + this.cy}
		}
	}
}
