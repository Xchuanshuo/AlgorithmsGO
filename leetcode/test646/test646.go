package test646

import "sort"

func findLongestChain(pairs [][]int) int {
	sort.Slice(pairs, func(i, j int) bool {
		if pairs[i][0] == pairs[j][0] {
			return pairs[i][1] > pairs[j][1]
		}
		return pairs[i][0] < pairs[j][0]
	})
	res, last := 0, 0
	for idx, p := range pairs {
		if idx == 0 || p[0] > last {
			res++
			last = p[1]
		} else {
			last = min(last, p[1])
		}
	}
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
