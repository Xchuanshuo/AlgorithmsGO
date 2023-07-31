package test2050

func minimumTime(n int, relations [][]int, time []int) int {
	var degree = make([]int, n)
	var g = make(map[int][]int)
	for _, r := range relations {
		g[r[0]-1] = append(g[r[0]-1], r[1]-1)
		degree[r[1]-1]++
	}
	var f = make([]int, n)
	var q = make([]int, 0)
	for i, v := range degree {
		f[i] = time[i]
		if v == 0 {
			q = append(q, i)
		}
	}
	for len(q) > 0 {
		var c = q[0]
		q = q[1:]
		for _, nxt := range g[c] {
			degree[nxt]--
			f[nxt] = max(f[nxt], f[c]+time[nxt])
			if degree[nxt] == 0 {
				q = append(q, nxt)
			}
		}
	}
	var res = 0
	for i := 0; i < n; i++ {
		res = max(res, f[i])
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
