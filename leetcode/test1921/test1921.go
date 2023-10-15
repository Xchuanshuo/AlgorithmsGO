package test1921

import "sort"

func eliminateMaximum(dist []int, speed []int) int {
	var mp = make(map[int]int)
	for i, d := range dist {
		var t = (d + speed[i] - 1) / speed[i]
		mp[t]++
	}
	var a = make([][]int, 0)
	for k, v := range mp {
		a = append(a, []int{k, v})
	}
	sort.Slice(a, func(i, j int) bool {
		return a[i][0] < a[j][0]
	})
	var res = 0
	for _, v := range a {
		if v[0] < res+v[1] {
			res += v[0] - res
			break
		}
		res += v[1]
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
