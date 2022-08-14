package test768

import "sort"

/**
 * @Description 与769类似，重复先用map记录，然后每次取靠前的索引即可
 **/

func maxChunksToSorted(arr []int) int {
	var mp = make(map[int][]int)
	var clone = make([]int, len(arr))
	copy(clone, arr)
	sort.Ints(clone)
	for i := 0; i < len(clone); i++ {
		var v = clone[i]
		if _, exist := mp[v]; !exist {
			mp[v] = make([]int, 0)
		}
		mp[v] = append(mp[v], i)
	}
	m, res := 0, 0
	for i := 0; i < len(arr); i++ {
		var v = arr[i]
		m = max(m, mp[v][0])
		mp[v] = mp[v][1:]
		if i == m {
			res++
		}
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
