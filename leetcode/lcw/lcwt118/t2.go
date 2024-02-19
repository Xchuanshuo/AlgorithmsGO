package lcwt118

import "sort"

func maximizeSquareHoleArea(n int, m int, hBars []int, vBars []int) int {
	var cal = func(a []int) int {
		sort.Ints(a)
		var m = len(a)
		var i = 0
		var w = 1
		for i < m {
			var j = i + 1
			for ; j < m; j++ {
				if a[j]-a[j-1] != 1 {
					break
				}
			}
			w = max(w, j-i+1)
			i = j
		}
		return w
	}
	hBars = append(hBars)
	vBars = append(vBars)
	w, h := cal(hBars), cal(vBars)
	return min(w, h) * min(w, h)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
