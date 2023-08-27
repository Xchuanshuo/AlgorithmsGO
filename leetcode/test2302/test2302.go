package test2302

/**
 * @Description https://leetcode.cn/problems/count-subarrays-with-score-less-than-k/
 * idea: 滑动窗口 求len*sum的子数组，看加入当前元素后是否满足条件，满足则以当前结尾的所有子数组都计入答案
				不满足则移动左窗口，找到满足条件的最大区间为止
 **/

func countSubarrays(nums []int, k int64) int64 {
	var n = len(nums)
	l, r := 0, 0
	var sum = int64(0)
	var res = int64(0)
	for r < n {
		sum += int64(nums[r])
		r++
		for int64(r-l)*sum >= k {
			sum -= int64(nums[l])
			l++
		}
		res += int64(r - l)
	}
	return res
}
