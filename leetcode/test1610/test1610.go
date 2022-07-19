package test1610

import (
	"math"
	"sort"
)

func visiblePoints(points [][]int, angle int, location []int) int {
	x, y := location[0], location[1]
	t := float64(angle) * math.Pi / 180
	list := make([]float64, 0, len(points))
	same := 0
	for _, p := range points {
		x1, y1 := p[0], p[1]
		if x1 == x && y1 == y {
			same++
		} else {
			list = append(list, math.Atan2(float64(y1-y), float64(x1-x)))
		}
	}
	sort.Float64s(list)
	max, n := 0, len(list)
	for i := 0;i < n;i++ {
		list = append(list, list[i] + 2 * math.Pi)
	}
	for i, j := 0, 0; j < 2*n;j++ {
		for i < j && list[j] - list[i] > t {
			i++
		}
		if max < j - i + 1 {
			max = j - i + 1
		}
	}
	return max + same
}
