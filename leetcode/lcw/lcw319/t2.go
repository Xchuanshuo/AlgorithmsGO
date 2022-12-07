package lcw319

func subarrayLCM(nums []int, k int) int {
	var n = len(nums)
	var res = 0
	for i := 0; i < n; i++ {
		j, m := i, nums[i]
		var pre = 1
		for ; j < n; j++ {
			m = max(m, (nums[j]*pre)/getFac(nums[j], pre))
			if m > k {
				break
			}
			if m == k {
				res++
			}
			pre = max(pre, nums[j])
		}
	}
	return res
}

func getFac(a, b int) int {
	if b == 0 {
		return a
	}
	return getFac(b, a%b)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
