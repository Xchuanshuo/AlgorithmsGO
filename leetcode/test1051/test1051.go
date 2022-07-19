package test1051

import "sort"

func heightChecker(heights []int) int {
	var copyH = make([]int, 0)
	for _, h := range heights {
		copyH = append(copyH, h)
	}
	sort.SliceStable(heights, func(i, j int) bool {
		return heights[i] < heights[j]
	})
	var res = 0
	for idx, v := range heights {
		if v != copyH[idx] {
			res++
		}
	}
	return res
}
