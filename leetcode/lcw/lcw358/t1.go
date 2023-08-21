package lcw358

func maxSum(nums []int) int {
	var n = len(nums)
	var cal = func(t int) int {
		var mx = 0
		for t != 0 {
			mx = max(mx, t%10)
			t /= 10
		}
		return mx
	}
	var res = 0
	for i, v := range nums {
		var m1 = cal(v)
		for j := i + 1; j < n; j++ {
			var m2 = cal(nums[j])
			if m1 == m2 {
				res = max(res, v+nums[j])
			}
		}
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
