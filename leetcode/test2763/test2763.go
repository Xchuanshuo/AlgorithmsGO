package test2763

func sumImbalanceNumbers(nums []int) int {
	var n = len(nums)
	var res = 0
	for i := 0; i < n; i++ {
		var vis = make(map[int]bool)
		vis[nums[i]] = true
		var cur = 0
		for j := i + 1; j < n; j++ {
			var x = nums[j]
			if vis[x-1] && vis[x+1] && !vis[x] {
				cur -= 1
			} else if !vis[x-1] && !vis[x+1] && !vis[x] {
				cur += 1
			}
			vis[x] = true
			cur = max(cur, 0)
			res += cur
		}
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
