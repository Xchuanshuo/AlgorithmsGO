package test2044

func countMaxOrSubsets(nums []int) int {
	var n = len(nums)
	res, max := 0, 0
	for i := 1; i < 1<<n; i++ {
		var cur = 0
		for j := 0; j < n; j++ {
			if i&(1<<j) != 0 {
				cur |= nums[j]
			}
		}
		if cur > max {
			max = cur
			res = 1
		} else if cur == max {
			res++
		}
	}
	return res
}
