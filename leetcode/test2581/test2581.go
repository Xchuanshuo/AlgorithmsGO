package test2581

/**
 * @Description https://leetcode.cn/problems/count-number-of-possible-root-nodes/
 * idea: 换根dp
 **/

func rootCount(edges [][]int, guesses [][]int, k int) int {
	var g = make(map[int][]int)
	for _, e := range edges {
		g[e[0]] = append(g[e[0]], e[1])
		g[e[1]] = append(g[e[1]], e[0])
	}
	var valid = make(map[T]int)
	for _, p := range guesses {
		valid[T{p[0], p[1]}] = 1
	}
	var dfs1 func(fa, cur int) int
	dfs1 = func(fa, cur int) int {
		var res = valid[T{fa, cur}]
		for _, nxt := range g[cur] {
			if nxt == fa {
				continue
			}
			res += dfs1(cur, nxt)
		}
		return res
	}
	var dp = make([]int, len(edges)+1)
	var dfs2 func(fa, cur int, sum int)
	dfs2 = func(fa, cur int, sum int) {
		dp[cur] = sum
		for _, nxt := range g[cur] {
			if nxt == fa {
				continue
			}
			var c = sum
			if valid[T{cur, nxt}] == 1 {
				c--
			}
			if valid[T{nxt, cur}] == 1 {
				c++
			}
			dfs2(cur, nxt, c)
		}
	}
	dfs2(-1, 0, dfs1(-1, 0))
	var res = 0
	for _, v := range dp {
		if v >= k {
			res++
		}
	}
	return res
}

type T struct {
	v, w int
}
