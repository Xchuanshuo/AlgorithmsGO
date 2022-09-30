package test1779

import "sort"

func nearestValidPoint(x int, y int, points [][]int) int {
	var pairs = make([]*Pair, 0)
	for i, p := range points {
		pairs = append(pairs, &Pair{p, i})
	}
	sort.Slice(pairs, func(i, j int) bool {
		a, b := pairs[i].point, pairs[j].point
		var da = abs(x-a[0]) + abs(y-a[1])
		var db = abs(x-b[0]) + abs(y-b[1])
		if da == db {
			return pairs[i].idx < pairs[j].idx
		}
		return da < db
	})
	for _, pair := range pairs {
		var p = pair.point
		if p[0] == x || p[1] == y {
			return pair.idx
		}
	}
	return -1
}

type Pair struct {
	point []int
	idx   int
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
