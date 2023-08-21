package test2354

import (
	"math/bits"
	"sort"
)

// bit1(n1|n2) + bit1(n1&n2) => bit1(n2) + bit1(n2)
// 所以去重后，按bit1排序，二分查找k-v下界bit1即可

func countExcellentPairs(nums []int, k int) int64 {
	var res = int64(0)
	var mp = make(map[int]bool)
	for _, v := range nums {
		mp[v] = true
	}
	var a = make([]int, 0)
	for v := range mp {
		a = append(a, bits.OnesCount(uint(v)))
	}
	sort.Ints(a)
	for _, v := range a {
		var t = k - v
		var j = sort.SearchInts(a, t)
		res += int64(len(a) - j)
	}
	return res
}
