package test954

import (
	"math"
	"sort"
)

func canReorderDoubled(arr []int) bool {
	var mp = make(map[int]int)
	var a = make([]int, 0)
	for _, v := range arr {
		if _, exist := mp[v]; !exist {
			a = append(a, v)
		}
		mp[v]++
	}
	sort.Slice(a, func(i, j int) bool {
		return math.Abs(float64(a[i])) < math.Abs(float64(a[j]))
	})
	for _, v := range a {
		if mp[v] > mp[2*v] {
			return false
		}
		mp[2*v] -= mp[v]
	}
	return true
}
