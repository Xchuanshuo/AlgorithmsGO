package lcwt105

func maxStrength(nums []int) int64 {
	var n = len(nums)
	var m = 1 << n
	var res = int64(nums[0])
	for s := 1; s < m; s++ {
		var cur int64 = 1
		for i := 0; i < n; i++ {
			if (1<<i)&s == 0 {
				continue
			}
			cur *= int64(nums[i])
		}
		res = max(res, cur)
	}
	return res
}

func max(a, b int64) int64 {
	if a > b { return a }
	return b
}
