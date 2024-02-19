package lcw370

func maximumScoreAfterOperations(edges [][]int, values []int) int64 {
	var g = make(map[int][]int)
	for _, e := range edges {
		g[e[0]] = append(g[e[0]], e[1])
		g[e[1]] = append(g[e[1]], e[0])
	}
	// dp0-1个不保留，dp1-至少保留1个
	var dfs func(fa, cur int) (int64, int64)
	dfs = func(fa, cur int) (int64, int64) {
		var dp0 = int64(values[cur])
		var dp1 = int64(0)
		var pre = int64(0)
		var isLeaf = true
		for _, nxt := range g[cur] {
			if nxt == fa {
				continue
			}
			isLeaf = false
			t0, t1 := dfs(cur, nxt)
			pre += t1
			dp0 += t0 // 之前的全取
			dp1 += t0
		}
		if isLeaf {
			return dp0, dp1
		}
		// 保留1个。去掉当前和保留当前两者取最大值
		dp1 = maxI64(dp1, int64(values[cur])+pre)
		return dp0, dp1
	}
	_, dp1 := dfs(-1, 0)
	return dp1
}

func maxI64(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}
