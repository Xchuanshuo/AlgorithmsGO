package test961

func repeatedNTimes(nums []int) int {
	var n = len(nums)
	for i := 0; i <= n; i++ {
		if nums[i%n] == nums[(i+1)%n] {
			return nums[i%n]
		}
		if nums[i%n] == nums[(i+2)%2] {
			return nums[i%n]
		}
	}
	return -1
}
