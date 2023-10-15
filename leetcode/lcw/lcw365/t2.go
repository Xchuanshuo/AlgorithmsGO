package lcw365

func maximumTripletValue(nums []int) int64 {
	var n = len(nums)
	var right = make([]int, n)
	right[n-1] = nums[n-1]
	for i := n - 2; i >= 0; i-- {
		right[i] = max(right[i+1], nums[i])
	}
	var res int64
	var lmx = nums[0]
	for i := 1; i < n-1; i++ {
		res = maxI64(res, int64(lmx-nums[i])*int64(right[i+1]))
		lmx = max(lmx, nums[i])
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxI64(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}
