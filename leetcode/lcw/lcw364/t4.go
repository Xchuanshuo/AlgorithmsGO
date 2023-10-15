package lcw364

import "math"

/**
 * @Description https://leetcode.cn/problems/count-valid-paths-in-a-tree/
 * idea: 树型dp、分情况讨论
		dp[i][0]节点i有0个质数的路径数，i有1个质数的路径数
		1.当前为质数, 只能与不为质数的子树组成路径
		2.当前不为质数，能与只为1的子树和0,1组成的子树组成路径
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
	var res = int64(0)
	var dfs func(fa, cur int) (int64, int64)
	dfs = func(fa, cur int) (int64, int64) {
		// 当前数为素数 答案包含2部分
		//  1.当前节点->nxt节点下为质数为0的路径；   //  1->0
		//  2.nxt为0的路径->当前节点->前面为0的数枝  // 0->1->0
		if isPrim[cur] {
			var pre0 = int64(0)
			for _, nxt := range g[cur] {
				if nxt == fa {
					continue
				}
				nxt0, _ := dfs(cur, nxt)
				res += nxt0 * pre0
				pre0 += nxt0
			}
			res += pre0
			// 当前所有子树为1的路径总数+当前数本身
			return 0, pre0 + 1
		}
		// 当前数不为素数 答案包含3部分
		// 1.nxt包含1个质数的路径->cur节点->前面包含0个质数的路径 1->0->0
		// 2.nxt包含0个质数的路径->cur节点->前面包含1个质数的路径 0->0->1
		// 3.cur节点->到所有为1的路径                          0->1
		dp0, dp1 := int64(0), int64(0)
		for _, nxt := range g[cur] {
			if nxt == fa {
				continue
			}
			nxt0, nxt1 := dfs(cur, nxt)
			res += dp0*nxt1 + dp1*nxt0
			dp0 += nxt0
			dp1 += nxt1
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
