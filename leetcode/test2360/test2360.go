package test2360

// 1.跑一遍拓扑排序 2.剩余节点存在环 从一个节点一直往后寻找即可
func longestCycle(edges []int) int {
	var n = len(edges)
	var degree = make([]int, n)
	for _, e := range edges {
		if e == -1 {
			continue
		}
		degree[e]++
	}
	var q = make([]int, 0)
	for i := 0; i < n; i++ {
		if degree[i] == 0 {
			q = append(q, i)
		}
	}
	for len(q) > 0 {
		var cur = q[0]
		q = q[1:]
		var nxt = edges[cur]
		if nxt == -1 {
			continue
		}
		degree[nxt]--
		if degree[nxt] == 0 {
			q = append(q, nxt)
		}
	}
	var res = -1
	for i := 0; i < n; i++ {
		if degree[i] == 0 {
			continue
		}
		cur, cnt := i, 0
		for cur != -1 && degree[cur] > 0 {
			degree[cur]--
			cur = edges[cur]
			cnt++
		}
		res = max(res, cnt)
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
