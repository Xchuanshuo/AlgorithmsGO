package test1574

/**
 * @Description https://leetcode.cn/problems/shortest-subarray-to-be-removed-to-make-array-sorted/
 * idea: 双指正 分别从左到右找到不满足非递减条件的临界点
		易错点: 1.删除可能完全删除左边或者右边，两种情况都需要考虑
			   2.删除中间时，直接使用类似归并的逻辑 能保证较小的元素能先加入结果集
 **/

func findLengthOfShortestSubarray(arr []int) int {
	var n = len(arr)
	if n == 1 {
		return 0
	}
	l, r := 0, n-1
	for l < n-1 && arr[l] <= arr[l+1] {
		l++
	}
	for r > 0 && arr[r-1] <= arr[r] {
		r--
	}
	if l > r {
		return 0
	}
	var res = min(r, n-l-1)
	i, j := 0, r
	for i <= l && j < n {
		if arr[j] >= arr[i] {
			res = min(res, j-i-1)
			i++
		} else {
			j++
		}
	}
	return max(0, res)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
