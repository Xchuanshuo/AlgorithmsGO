package lcwt111

import "sort"

func minimumOperations(nums []int) int {
	var n = len(nums)
	var calc = func(a []int) int {
		var n = len(a)
		var a1 = make([]int, n+1)
		var l = 1
		a1[0] = a[0]
		for i := 1; i < n; i++ {
			if a[i] >= a1[l-1] {
				a1[l] = a[i]
				l++
			} else {
				var pos = sort.Search(l, func(j int) bool {
					return a1[j] > a[i]
				})
				a1[pos] = a[i]
			}
		}
		return l
	}
	return n - calc(nums)
}
