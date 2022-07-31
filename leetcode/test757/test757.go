package test757

import "sort"

func intersectionSizeTwo(invs [][]int) int {
	sort.Slice(invs, func(i, j int) bool {
		a, b := invs[i], invs[j]
		if a[1] == b[1] {
			return b[0] < a[0]
		}
		return a[1] < b[1]
	})
	var res = 0
	m1, m2 := -1, -1
	for _, a := range invs {
		if a[0] > m1 {
			res += 2
			m1 = a[1]
			m2 = a[1] - 1
		} else if a[0] > m2 {
			res += 1
			m2 = m1
			m1 = a[1]
		}
	}
	return res
}
