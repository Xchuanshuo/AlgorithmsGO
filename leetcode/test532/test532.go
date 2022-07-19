package test532

func findPairs(nums []int, k int) int {
	var visited = make(map[int]bool)
	var res = make(map[int]bool)
	for _, v := range nums {
		if visited[v-k] {
			res[v-k] = true
		}
		if visited[v+k] {
			res[v] = true
		}
		visited[v] = true
	}
	return len(res)
}
