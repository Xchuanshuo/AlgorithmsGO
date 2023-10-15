package lcw364

import "math"

// dp[i][0]节点i有0个质数的路径数，i有1个质数的路径数
// 1.当前为质数, 只能与不为质数的子树组成路径
// 2.当前不为质数，能与只为1的子树和0,1组成的子树组成路径

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
	var dp = make([][2]int64, n+1)
	var res = int64(0)
	var dfs func(fa, cur int) (int64, int64)
	dfs = func(fa, cur int) (int64, int64) {
		// 当前数为素数
		if isPrim[cur] {
			var pre0 = int64(0)
			for _, nxt := range g[cur] {
				if nxt == fa {
					continue
				}
				dfs(cur, nxt)
				// 包含两部分: 1.当前数->每个为0的数枝；2.前面为0的树枝->当前数->当前为0的数枝
				res += dp[nxt][0] * (pre0 + 1)
				pre0 += dp[nxt][0]
			}
			// 当前所有子树为1的路径总数+当前数本身
			return 0, pre0 + 1
		}
		// 当前数不为素数
		// 1.当前包含1个质数的路径->cur节点->前面包含0个质数的路径
		// 2.当前包含0个质数的路径->cur节点->前面包含1个质数的路径
		// 3.cur节点->到所有为1的路径
		dp0, dp1 := int64(0), int64(0)
		for _, nxt := range g[cur] {
			if nxt == fa {
				continue
			}
			dfs(cur, nxt)
			res += dp0*dp[nxt][1] + dp1*dp[nxt][0]
			dp0 += dp[nxt][0]
			dp1 += dp[nxt][1]
		}
		res += dp1
		// 当前所有子树为0的路径总数+当前数本身
		return dp0 + 1, dp1
	}
	dfs(-1, 1)
	return res
}

func getPrims(n int) []bool {
	var isPrim = make([]bool, n+1)
	for i := 2; i < len(isPrim); i++ {
		isPrim[i] = true
	}
	for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
		if isPrim[i] {
			for j := i * i; j <= n; j += i {
				isPrim[j] = false
			}
		}
	}
	return isPrim
}
