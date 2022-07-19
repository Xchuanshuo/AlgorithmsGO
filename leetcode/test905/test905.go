package test905

func sortArrayByParity(nums []int) []int {
	var n = len(nums)
	l, r := 0, n-1
	for ; l < r; l++ {
		if nums[l]%2 == 0 {
			continue
		}
		for r > l && nums[r]%2 == 1 {
			r--
		}
		nums[l], nums[r] = nums[r], nums[l]
	}
	return nums
}
