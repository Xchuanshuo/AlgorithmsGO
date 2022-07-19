package test713

func numSubarrayProductLessThanK(nums []int, k int) int {
	var n = len(nums)
	l, r := 0, 0
	val, res := 1, 0
	for r < n {
		var num = nums[r]
		r++
		val *= num
		for val > k {
			var pre = nums[l]
			l++
			val /= pre
		}
		res += (r - l) + 1
	}
	return res
}
