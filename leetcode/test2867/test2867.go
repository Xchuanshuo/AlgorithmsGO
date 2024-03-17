package test2867

/**
 * @Description https://leetcode.cn/problems/count-valid-paths-in-a-tree
 * idea: 联通分量
		1.按质数拆分联通块，记录每个联通块的根节点防止重复计算，经过单个质数的所有路径都计入答案
		2.使用乘法原理计算，每个子树节点数两两相乘+质数节点的每个子树节点数量，如(a,b,c)，实际路径数为
			a*b + a*c + b*c + a + b + c，即 a*b + (a+b)*c + a + b + c
 **/

var isPrim []bool

func init() {
	isPrim = getPrims(int(1e5))
}

func countPaths(n int, edges [][]int) int64 {
	var g = make(map[int][]int)
	for _, e := range edges {
		g[e[0]] = append(g[e[0]], e[1])
		g[e[1]] = append(g[e[1]], e[0])
	}
	var sz = make([]int, n+1)
	var parent = make([]int, n+1)
	var dfs func(fa, cur int, root int) int
	dfs = func(fa, cur int, root int)int {
		parent[cur] = root
		var cnt = 1
		for _, nxt := range g[cur] {
			if nxt == fa || isPrim[nxt] {
				continue
			}
			cnt += dfs(cur, nxt, root)
		}
		return cnt
	}
	var res = int64(0)
	for i := 1; i <= n; i++ {
		if !isPrim[i] {
			continue
		}
		var sum = int64(0)
		for _, nxt := range g[i] {
			if isPrim[nxt] {
				continue
			}
			if parent[nxt] == 0 {
				sz[nxt] = dfs(-1, nxt, nxt)
			}
			res += sum * int64(sz[parent[nxt]])
			sum += int64(sz[parent[nxt]])
		}
		res += sum
	}
	return res
}

func getPrims(n int) []bool {
	var prim = make([]bool, n+1)
	for i := 2; i < n; i++ {
		prim[i] = true
	}
	for i := 2; i <= n; i++ {
		if !prim[i] {
			continue
		}
		for j := 2 * i; j <= n; j += i {
			prim[j] = false
		}
	}
	return prim
}
