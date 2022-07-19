package test908

func smallestRangeI(nums []int, k int) int {
	max, min := 0, int(1e9)
	for i := 0; i < len(nums); i++ {
		if nums[i] > max {
			max = nums[i]
		}
		if nums[i] < min {
			min = nums[i]
		}
	}
	var dif = max - min - 2*k
	if dif < 0 {
		return 0
	}
	return dif
}
