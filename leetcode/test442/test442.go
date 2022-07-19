package test442

func findDuplicates(nums []int) []int {
	var res = make([]int, 0)
	for _, num := range nums {
		var v = abs(num)
		if nums[v-1] < 0 {
			res = append(res, v)
		} else {
			nums[v-1] *= -1
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
