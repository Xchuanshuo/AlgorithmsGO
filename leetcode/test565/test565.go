package test565

func arrayNesting(nums []int) int {
	var n = len(nums)
	var visited = make([]bool, n)
	var res = 0
	for i := 0; i < n; i++ {
		cnt, c := 0, i
		for !visited[c] {
			visited[c] = true
			c = nums[c]
			cnt++
		}
		if cnt > res {
			res = cnt
		}
	}
	return res
}
