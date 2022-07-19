package test2045


func secondMinimum(n int, edges [][]int, time int, change int) int {
	var graph = make(map[int][]int)
	for _, e := range edges {
		v, w := e[0]-1, e[1]-1
		graph[v] = append(graph[v], w)
		graph[w] = append(graph[w], v)
	}
	next := func(d int) int {
		if times := d/change;times%2 == 1 {
			return d + (change-d%change) + time
		}
		return d + time
	}
	dist := make([][]int, n)
	for i := 0;i < n;i++ {
		dist[i] = []int{1e9, 1e9}
	}
	q := make([][]int, 0)
	q = append(q, []int{0, 0})
	for len(q) > 0 {
		cur := q[0]
		q = q[1:]
		for _, w := range graph[cur[0]] {
			d := next(cur[1])
			if d < dist[w][0] {
				dist[w][0] = d
				q = append(q, []int{w, d})
			} else if d > dist[w][0] && d < dist[w][1] {
				dist[w][1] = d
				q = append(q, []int{w, d})
			}
		}
	}
	return dist[n-1][1]
}