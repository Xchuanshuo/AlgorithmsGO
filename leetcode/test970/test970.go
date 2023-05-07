package test970

import "math"

func powerfulIntegers(x int, y int, bound int) []int {
	var set = make(map[int]bool)
	for i := 0; i < 21; i++ {
		for j := 0; j < 21; j++ {
			var v = int(math.Pow(float64(x), float64(i)) + math.Pow(float64(y), float64(j)))
			if v > 0 && v <= bound {
				set[v] = true
			}
		}
	}
	var res = make([]int, 0)
	for v := range set {
		res = append(res, v)
	}
	return res
}
