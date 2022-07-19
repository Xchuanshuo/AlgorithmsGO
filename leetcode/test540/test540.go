package test540

func singleNonDuplicate(nums []int) int {
	var n = len(nums)
	l, r := 0, n-1
	for l < r {
		mid := l + (r-l)/2
		if nums[mid] == nums[mid^1] {
			l = mid + 1
		} else {
			r = mid
		}
	}
	return nums[l]
}
