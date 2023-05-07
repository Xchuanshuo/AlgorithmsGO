package lcwt100

import (
	"github.com/emirpasic/gods/trees/redblacktree"
)

func findScore(nums []int) int64 {
	var tree = redblacktree.NewWith(func(a, b interface{}) int {
		a1, b1 := a.([]int), b.([]int)
		if a1[0] == b1[0] {
			return a1[1] - b1[1]
		}
		return a1[0] - b1[0]
	})
	for i, v := range nums {
		tree.Put([]int{v, i}, struct{}{})
	}
	var res int64 = 0
	for !tree.Empty() {
		var cur = tree.Left().Key.([]int)
		tree.Remove(cur)
		l, r := cur[1]-1, cur[1]+1
		if l >= 0 {
			tree.Remove([]int{nums[l], l})
		}
		if r < len(nums) {
			tree.Remove([]int{nums[r], r})
		}
		res += int64(cur[0])
	}
	return res
}
