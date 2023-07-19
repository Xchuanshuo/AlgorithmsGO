package s1

func findMinHeightTrees(n int, edges [][]int) []int {
	if n == 1 {
		return []int{0}
	}
	var g = make(map[int][]int)
	var degree = make([]int, n)
	for _, e := range edges {
		v, w := e[0], e[1]
		if _, exist := g[v]; !exist {
			g[v] = make([]int, 0)
		}
		if _, exist := g[w]; !exist {
			g[w] = make([]int, 0)
		}
		g[v] = append(g[v], w)
		g[w] = append(g[w], v)
		degree[v]++
		degree[w]++
	}
	var q = make([]int, 0)
	for i := 0; i < n; i++ {
		if degree[i] == 1 {
			q = append(q, i)
		}
	}
	var res []int
	for len(q) > 0 {
		var size = len(q)
		res = make([]int, 0)
		for i := 0; i < size; i++ {
			cur := q[0]
			res = append(res, cur)
			q = q[1:]
			for _, next := range g[cur] {
				degree[next]--
				if degree[next] == 1 {
					q = append(q, next)
				}
			}
		}
	}
	return res
}
