package lcwt99

/**
 * @Description https://leetcode.cn/problems/count-number-of-possible-root-nodes/
 * idea: 换根dp，统计每个节点为根时bob猜对的数量即可
 **/

func rootCount(edges [][]int, guesses [][]int, k int) int {
	var g = make(map[int][]int)
	for _, e := range edges {
		g[e[0]] = append(g[e[0]], e[1])
		g[e[1]] = append(g[e[1]], e[0])
	}
	var valid = make(map[T]bool)
	for _, c := range guesses {
		valid[T{c[0], c[1]}] = true
	}
	var dfs1 func(fa, cur int) int
	dfs1 = func(fa, cur int) int {
		var cnt = 0
		for _, nxt := range g[cur] {
			if nxt == fa {
				continue
			}
			if valid[T{cur, nxt}] {
				cnt++
			}
			cnt += dfs1(cur, nxt)
		}
		return cnt
	}
	var n = len(edges)
	var dp = make([]int, n+1)
	var dfs2 func(fa, cur, cnt int)
	dfs2 = func(fa, cur, cnt int) {
		dp[cur] = cnt
		for _, nxt := range g[cur] {
			if nxt == fa {
				continue
			}
			var sum = cnt
			// 换根 cur->nxt
			if valid[T{cur, nxt}] {
				sum--
			}
			if valid[T{nxt, cur}] {
				sum++
			}
			dfs2(cur, nxt, sum)
		}
	}
	// 1.第一遍dfs计算以0为根的答案
	var sum = dfs1(-1, 0)
	// 2.第二遍从节点0开始换根
	dfs2(-1, 0, sum)
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
