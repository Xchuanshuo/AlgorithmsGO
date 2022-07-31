package test115

func sequenceReconstruction(nums []int, sequences [][]int) bool {
	var n = len(nums)
	var graph = make(map[int][]int)
	var inDegree = make([]int, n+1)
	for _, seq := range sequences {
		for i := 1; i < len(seq); i++ {
			pre, next := seq[i-1], seq[i]
			inDegree[next]++
			if _, exist := graph[pre]; !exist {
				graph[pre] = make([]int, 0)
			}
			graph[pre] = append(graph[pre], next)
		}
	}
	var q = make([]int, 0)
	for i := 1; i <= n; i++ {
		if inDegree[i] == 0 {
			q = append(q, i)
		}
	}
	for len(q) > 0 {
		if len(q) > 1 {
			return false
		}
		var c = q[0]
		q = q[1:]
		for _, nxt := range graph[c] {
			inDegree[nxt]--
			if inDegree[nxt] == 0 {
				q = append(q, nxt)
			}
		}
	}
	return true
}
