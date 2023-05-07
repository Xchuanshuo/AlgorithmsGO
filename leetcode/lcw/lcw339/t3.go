package lcw339

import (
	"sort"
)

func miceAndCheese(r1 []int, r2 []int, k int) int {
	var n = len(r1)
	var a = make([][]int, 0)
	for i := 0; i < n; i++ {
		a = append(a, []int{r1[i], r2[i]})
	}
	sort.Slice(a, func(i, j int) bool {
		if abs(a[i][0]-a[i][1]) == abs(a[j][0]-a[j][1]) {
			return a[i][0]+a[i][1] < a[j][0]+a[j][1]
		}
		return abs(a[i][0]-a[i][1]) < abs(a[j][0]-a[j][1])
	})
	var res = 0
	for i := n - 1; i >= 0; i-- {
		if i >= k && (k == 0 || a[i][1] > a[i][0]) {
			res += a[i][1]
		} else {
			res += a[i][0]
			k--
		}
	}
	return res
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
