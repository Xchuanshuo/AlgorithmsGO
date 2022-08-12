package test1413

func minStartValue(nums []int) int {
	var res = int(1e9)
	for i := 1; i <= 100; i++ {
		var cur = 0
		var isValid = true
		for j := 0; j < len(nums); j++ {
			cur += nums[j]
			if cur < 1 {
				isValid = false
				break
			}
		}
		if isValid && cur < res {
			res = cur
		}
	}
	return res
}
