package test2251

import "sort"

func fullBloomFlowers(flowers [][]int, people []int) []int {
	var a = make([][]int, 0)
	for _, f := range flowers {
		a = append(a, []int{f[0], -2})
		a = append(a, []int{f[1] + 1, -1})
	}
	for i, p := range people {
		a = append(a, []int{p, i})
	}
	sort.Slice(a, func(i, j int) bool {
		if a[i][0] == a[j][0] {
			return a[i][1] < a[j][1]
		}
		return a[i][0] < a[j][0]
	})
	var res = make([]int, len(people))
	var cnt = 0
	for _, v := range a {
		if v[1] == -2 {
			cnt++
		} else if v[1] == -1 {
			cnt--
		} else {
			res[v[1]] = cnt
		}
	}
	return res
}
