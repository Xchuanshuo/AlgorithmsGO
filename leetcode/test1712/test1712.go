package test1712

/**
 * @Description https://leetcode.cn/problems/ways-to-split-array-into-three-subarrays/
 * idea: 前缀和，直接枚举分割位置i,j，时间复杂度不满足要求。重新考虑条件, 假设第一段和为s1, 实际第
		二段和开始位置至少需要满足l>=s1, 而数组元素都为正数，所以可以直接对前缀和数组进行【二分查找】；
		对于第二段和右边界，需要满足r<=s1+(sum-s1)/2, 即保证第三段和大于等于第二段和。同理二分查找即可
 **/

import "sort"

var M = int(1e9) + 7

func waysToSplit(nums []int) int {
	var n = len(nums)
	var sum = make([]int, n+1)
	for i := 0; i < n; i++ {
		sum[i+1] = sum[i] + nums[i]
	}
	var res = 0
	for i := 1; i <= n; i++ {
		var s1 = sum[i]
		var l = i + 1 + sort.SearchInts(sum[i+1:], 2*s1)
		var r = sort.Search(len(sum), func(i int) bool {
			return sum[i] > (sum[n]-s1)/2+s1
		})
		if sum[n] == 0 {
			r = n
		}
		if l < r && (r != len(sum)) {
			res = (res + (r - l)) % M
		}
	}
	return res
}
