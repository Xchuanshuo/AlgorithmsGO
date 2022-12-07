package lcw315

/**
 * @Description https://leetcode.cn/problems/count-subarrays-with-fixed-bounds/
 * idea: 双指针，右端点 为minK, maxK 对应位置minI,maxI较小的, 这样能保证子数组[minI,maxI]包含minK,maxK
		 		左端点 为不在[minK,maxK]范围之类的最左边的数, 此时子数组数量为 min(minI,maxI)-l
				遍历过程中累加结果即可
 **/

func countSubarrays(nums []int, minK int, maxK int) int64 {
	var res int64 = 0
	l, maxI, minI := -1, -1, -1
	for i, num := range nums {
		if num == minK {
			minI = i
		}
		if num == maxK {
			maxI = i
		}
		if num < minK || num > maxK {
			l = i
		}
		res += int64(max(min(minI, maxI)-l, 0))
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
