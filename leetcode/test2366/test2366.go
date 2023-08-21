package test2366

func minimumReplacement(nums []int) int64 {
	var res int64
	var n = len(nums)
	var last = nums[n-1]
	for i := n - 2; i >= 0; i-- {
		if nums[i] > last {
			var cnt = (nums[i] + last - 1) / last
			res += int64(cnt - 1)
			last = nums[i] / cnt
		} else {
			last = nums[i]
		}
	}
	return res
}
