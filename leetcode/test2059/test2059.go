package test2059

// bfs 时间复杂度 入队元素最多1000, 每个元素与所有其它元素3次操作类型 整体复杂度为 O(m*n^2)
func minimumOperations(nums []int, start int, goal int) int {
	var step = 0
	var vis = make(map[int]bool)
	var q = []int{start}
	for len(q) > 0 {
		var sz = len(q)
		step++
		for i := 0; i < sz; i++ {
			var cur = q[0]
			q = q[1:]
			for _, v := range nums {
				var ds = []int{cur + v, cur - v, cur ^ v}
				for _, d := range ds {
					if d == goal {
						return step
					}
					if d < 0 || d > 1000 || vis[d] {
						continue
					}
					vis[d] = true
					q = append(q, d)
				}
			}
		}
	}
	return -1
}
