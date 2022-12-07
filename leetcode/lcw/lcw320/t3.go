package lcw320

/**
 * @Description https://leetcode.cn/problems/minimum-fuel-cost-to-report-to-the-capital/
 * idea:  贪心: 每到一个点, 看所有人数需要几个车，其它车扔掉即可
 **/

func minimumFuelCost(roads [][]int, seats int) int64 {
	var g = make(map[int][]int)
	for _, r := range roads {
		g[r[0]] = append(g[r[0]], r[1])
		g[r[1]] = append(g[r[1]], r[0])
	}
	var res int64 = 0
	var dfs func(f, c int) int
	dfs = func(f, c int) int {
		var cur = 1
		for _, nxt := range g[c] {
			if nxt == f {
				continue
			}
			cur += dfs(c, nxt)
		}
		if c > 0 {
			res += int64((cur + seats - 1) / seats)
		}
		return cur
	}
	dfs(-1, 0)
	return res
}

//[[0,1],[0,2],[3,2],[0,4],[1,5],[5,6],[3,7]]
//1
//输出：
//9
//预期：
//13

//[[0,1],[0,2],[1,3],[1,4]]
//5
//3
//4
