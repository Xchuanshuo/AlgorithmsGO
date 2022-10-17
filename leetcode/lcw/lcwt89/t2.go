package main

import "math"

func productQueries(n int, queries [][]int) []int {
	var powers = make([]int, 0)
	for i := 0; i < 32; i++ {
		if (1<<i)&n > 0 {
			powers = append(powers, int(math.Pow(2.0, float64(i))))
		}
	}
	var res = make([]int, 0)
	var M = int(1e9) + 7
	for _, q := range queries {
		var cur = 1
		for i := q[0]; i <= q[1]; i++ {
			cur = (cur * powers[i]) % M
		}
		res = append(res, cur)
	}
	return res
}
