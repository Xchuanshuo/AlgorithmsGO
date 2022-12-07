package lcw316

func subarrayGCD(nums []int, k int) int {
	var n = len(nums)
	var res = 0
	for i := 0; i < n; i++ {
		j, m := i, nums[i]
		var pre = 0
		for ; j < n; j++ {
			m = min(m, getFac(nums[j], pre))
			if nums[j]%k != 0 || m < k {
				break
			}
			if m == k {
				res++
			}
			pre = nums[j]
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

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
