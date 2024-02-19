package lcw371

func maximumStrongPairXor1(nums []int) int {
	var n = len(nums)
	var res = 0
	for i, v := range nums {
		for j := i + 1; j < n; j++ {
			var w = nums[j]
			if abs(v-w) <= min(v, w) {
				res = max(res, v^w)
			}
		}
	}
	return res
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
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
