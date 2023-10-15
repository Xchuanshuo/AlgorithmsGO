package lcw364

/**
 * @Description https://leetcode.cn/problems/beautiful-towers-ii/
 * idea: 单调栈 + 前缀和 + 贪心 + 枚举
		1.枚举每个数作为山顶。靠近山顶位置i左边的元素实际需要修改为min(heights[i-1], heights[i])。
		 修改操作需要应用到与山顶相邻且大于山顶的元素。所以可以使用单调栈进行维护递减顺序的元素位置。对于当前元素，
		 需要找到栈顶第一个<=山顶的元素j. 则需要修改的区间为[j,i]，修改后当前位置的值为left[j-1] + (i-j)*heights[i]
		2.分别计算左侧右侧后，枚举每个位置计算答案即可
 **/

func maximumSumOfHeights(heights []int) int64 {
	var n = len(heights)
	var left = make([]int64, n+1)
	var right = make([]int64, n+1)
	var st = make([]int, 0)
	for i := 0; i < n; i++ {
		var cur = heights[i]
		var cnt = 0
		for len(st) > 0 && heights[st[len(st)-1]] > cur {
			st = st[0 : len(st)-1]
		}
		if len(st) == 0 {
			cnt = i + 1
		} else {
			cnt = i - st[len(st)-1]
			left[i+1] = left[i+1-cnt]
		}
		left[i+1] += int64(cnt * cur)
		st = append(st, i)
	}
	var res = int64(0)
	st = make([]int, 0)
	for i := n; i >= 1; i-- {
		var cur = heights[i-1]
		var cnt = 0
		for len(st) > 0 && heights[st[len(st)-1]] > cur {
			st = st[0 : len(st)-1]
		}
		if len(st) == 0 {
			cnt = n - i + 1
		} else {
			cnt = st[len(st)-1] - i + 1
			right[i] = right[i+cnt]
		}
		right[i] += int64(cnt * cur)
		st = append(st, i-1)
		res = maxI64(res, left[i-1]+right[i])
	}
	return res
}

func maxI64(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}
