package main

import "github.com/emirpasic/gods/trees/redblacktree"

// 滑动窗口 + treeMap, 无需考虑右边界，用有序集合存储元素后，保证窗口内的最大与最小差值<=2即可

func continuousSubarrays(nums []int) int64 {
	var n = len(nums)
	var tree = redblacktree.NewWithIntComparator()
	var res int64 = 0
	l, r := 0, 0
	for r < n {
		var rv = 0
		v, ok := tree.Get(nums[r])
		if ok {
			rv = v.(int)
		}
		tree.Put(nums[r], rv+1)
		mxv := tree.Right().Key.(int)
		mnv := tree.Left().Key.(int)
		for l < r && abs(mxv-mnv) > 2 {
			lt, _ := tree.Get(nums[l])
			lv := lt.(int)
			if lv == 1 {
				tree.Remove(nums[l])
			} else {
				tree.Put(nums[l], lv-1)
			}
			mxv = tree.Right().Key.(int)
			mnv = tree.Left().Key.(int)
			l++
		}
		res += int64(r - l + 1)
		r++
	}
	return res
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
