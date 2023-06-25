package lcw349

func minCost(nums []int, x int) int64 {
	var n = len(nums)
	var cost = make([]int64, n)
	var res = int64(0)
	for i := 0; i < n; i++ {
		cost[i] = int64(nums[i])
		res += cost[i]
	}
	for i := 1; i < n; i++ {
		var cur = int64(i * x)
		for j := 0; j < n; j++ {
			cost[j] = min(cost[j], int64(nums[(j+i)%n]))
			cur += cost[j]
		}
		res = min(res, cur)
	}
	return res
}

func min(a, b int64) int64 {
	if a < b {
		return a
	}
	return b
}
