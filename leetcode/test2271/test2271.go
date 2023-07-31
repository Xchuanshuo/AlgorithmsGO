package test2271

/**
 * @Description https://leetcode.cn/problems/maximum-white-tiles-covered-by-a-carpet/
 * idea: 排序 + 滑动窗口
		   1.区间从小到大排序
		   2.维护大小在carpetLen内的窗口，若剩余覆盖区间大小小于最新区间，则将窗口整体左移
 **/

import "sort"

func maximumWhiteTiles(tiles [][]int, carpetLen int) int {
	sort.Slice(tiles, func(i, j int) bool {
		return tiles[i][0] < tiles[j][0]
	})
	var cur = carpetLen
	res, cnt := 0, 0
	var l = 0
	for r, t := range tiles {
		var size = t[1] - t[0] + 1
		if r > 0 {
			cur -= t[0] - tiles[r-1][1] - 1
		}
		for cur < size && l < r {
			res = max(res, cnt+cur)
			cur += tiles[l+1][0] - tiles[l][0]
			cnt -= tiles[l][1] - tiles[l][0] + 1
			l++
		}
		cnt += min(size, cur)
		cur -= size
		res = max(res, cnt)
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
