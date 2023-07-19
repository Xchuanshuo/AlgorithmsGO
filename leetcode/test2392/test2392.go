package test2392

func buildMatrix(k int, rows [][]int, cols [][]int) [][]int {
	// 根据限制条件 计算拓补序列
	var topSort = func(arr [][]int) []int {
		var g = make(map[int][]int)
		var degree = make([]int, k)
		for _, a := range arr {
			g[a[0]-1] = append(g[a[0]-1], a[1]-1)
			degree[a[1]-1]++
		}
		var q = make([]int, 0)
		for i, d := range degree {
			if d == 0 {
				q = append(q, i)
			}
		}
		var res = make([]int, 0)
		for len(q) > 0 {
			var cur = q[0]
			q = q[1:]
			res = append(res, cur)
			for _, nxt := range g[cur] {
				degree[nxt]--
				if degree[nxt] == 0 {
					q = append(q, nxt)
				}
			}
		}
		if len(res) != k {
			return []int{}
		}
		return res
	}
	var rs = topSort(rows)
	var cs = topSort(cols)
	if len(rs) == 0 || len(cs) == 0 {
		return [][]int{}
	}
	var res = make([][]int, k)
	for i := 0; i < k; i++ {
		res[i] = make([]int, k)
	}
	var mp = make(map[int]int)
	for i, c := range cs {
		mp[c] = i
	}
	// 行顺序+列顺序填充数据
	for i, v := range rs {
		res[i][mp[v]] = v + 1
	}
	return res
}
