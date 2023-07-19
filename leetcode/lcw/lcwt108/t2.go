package lcwt108

import "sort"

func relocateMarbles(nums []int, moveFrom []int, moveTo []int) []int {
	var mp = make(map[int]int)
	for _, v := range nums {
		mp[v]++
	}
	for i, f := range moveFrom {
		mp[f]--
		delete(mp, f)
		mp[moveTo[i]]++
	}
	var res = make([]int, 0)
	for k := range mp {
		res = append(res, k)
	}
	sort.Ints(res)
	return res
}
