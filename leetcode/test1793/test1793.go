package test1793

func maximumScore(src []int, k int) int {
	k += 1
	var n = len(src)
	var nums = append([]int{-1}, append(src, -1)...)
	var s = []int{0}
	var left = make([]int, n+2)
	for i := 1; i <= n; i++ {
		for len(s) > 0 && nums[s[len(s)-1]] >= nums[i] {
			s = s[0 : len(s)-1]
		}
		left[i] = s[len(s)-1] + 1
		s = append(s, i)
	}
	var res = 0
	s = []int{n + 1}
	for i := n; i >= 1; i-- {
		for len(s) > 0 && nums[s[len(s)-1]] >= nums[i] {
			s = s[0 : len(s)-1]
		}
		var rIdx = s[len(s)-1] - 1
		if left[i] <= k && k <= rIdx {
			res = max(res, (rIdx-left[i]+1)*nums[i])
		}
		s = append(s, i)
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
