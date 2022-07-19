package test1200

import "sort"

func minimumAbsDifference(arr []int) [][]int {
	sort.Ints(arr)
	var mp = make(map[int][][]int)
	var min int = 1e9
	for i := 1; i < len(arr); i++ {
		var v = abs(arr[i] - arr[i-1])
		if _, exist := mp[v]; !exist {
			mp[v] = make([][]int, 0)
		}
		mp[v] = append(mp[v], []int{arr[i-1], arr[i]})
		if v < min {
			min = v
		}
	}
	return mp[min]
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
