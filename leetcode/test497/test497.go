package test497

import "math/rand"

type Solution struct {
	rects [][]int
}

func Constructor(rects [][]int) Solution {
	return Solution{
		rects: rects,
	}
}

func (this *Solution) Pick() []int {
	cur, curSum, index := 0, 0, 0
	for idx, r := range this.rects {
		x1, y1 := r[0], r[1]
		x2, y2 := r[2], r[3]
		cur = (x2 - x1 + 1) * (y2 - y1 + 1)
		curSum += cur
		if rand.Int()%curSum < cur {
			index = idx
		}
	}
	var p = this.rects[index]
	x1, y1, x2, y2 := p[0], p[1], p[2], p[3]
	return []int{x1 + rand.Intn(x2-x1+1), y1 + rand.Intn(y2-y1+1)}
}
