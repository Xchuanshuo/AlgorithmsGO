package lcwt124

func maxOperations(nums []int) int {
	var n = len(nums)
	var mem = make(map[int][][]int)
	var dfs func(i, j, score int) int
	dfs = func(i, j, score int) int {
		if i >= j {
			return 0
		}
		if mem[score][i][j] != -1 {
			return mem[score][i][j]
		}
		var res = 0
		if nums[i]+nums[j] == score {
			res = max(res, 1+dfs(i+1, j-1, score))
		}
		if i+1 < n && nums[i]+nums[i+1] == score {
			res = max(res, 1+dfs(i+2, j, score))
		}
		if j-1 >= 0 && nums[j-1]+nums[j] == score {
			res = max(res, 1+dfs(i, j-2, score))
		}
		mem[score][i][j] = res
		return res
	}
	var t = []int{nums[0] + nums[1], nums[n-1] + nums[n-2], nums[0] + nums[n-1]}
	for _, v := range t {
		var cur = make([][]int, n)
		for i := 0; i < n; i++ {
			cur[i] = make([]int, n)
			for j := 0; j < n; j++ {
				cur[i][j] = -1
			}
		}
		mem[v] = cur
	}
	return 1 + max(dfs(2, n-1, t[0]), dfs(0, n-3, t[1]), dfs(1, n-2, t[2]))
}
