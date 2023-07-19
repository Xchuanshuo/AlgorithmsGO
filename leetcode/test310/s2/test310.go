package s2

func findMinHeightTrees(n int, edges [][]int) []int {
	var g = make(map[int][]int)
	for _, e := range edges {
		g[e[0]] = append(g[e[0]], e[1])
		g[e[1]] = append(g[e[1]], e[0])
	}
	var h = make([]int, n)
	var dfs1 func(f, cur int) int
	dfs1 = func(fa, cur int) int {
		var res = 0
		for _, nxt := range g[cur] {
			if nxt == fa {
				continue
			}
			res = max(res, dfs1(cur, nxt))
		}
		h[cur] = res + 1
		return h[cur]
	}
	var dp = make([]int, n)
	var dfs2 func(fa, cur int, height int)
	dfs2 = func(fa, cur int, height int) {
		// 先序遍历，保存最新结果
		dp[cur] = height
		f, s := 0, 0
		for _, nxt := range g[cur] {
			if h[nxt] > f {
				s, f = f, h[nxt]
			} else if h[nxt] > s {
				s = h[nxt]
			}
		}
		for _, nxt := range g[cur] {
			if nxt == fa {
				continue
			}
			// cur->nxt 换根具体计算逻辑 就本题而言
			// 1.换根重新计算cur子树高度 2.新根nxt新增【cur子树】，需与原nxt的子树比较高度
			if h[nxt] == f {
				h[cur] = s + 1
			} else {
				h[cur] = f + 1
			}
			dfs2(cur, nxt, max(h[nxt], h[cur]+1))
		}
	}
	var sum = dfs1(-1, 0)
	dfs2(-1, 0, sum)
	var mn = n
	for _, v := range dp {
		mn = min(mn, v)
	}
	var res = make([]int, 0)
	for i, v := range dp {
		if v == mn {
			res = append(res, i)
		}
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
