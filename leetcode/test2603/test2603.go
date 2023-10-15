package test2603

/**
 * @Description https://leetcode.cn/problems/collect-coins-in-a-tree
 * idea: 贪心+拓扑排序, 1.不带金币的叶子节点无需访问 跑一遍拓扑排序进行删除
					 2.当前所有叶子节点都带金币，由于能收集距离为2的金币，所以最多只需要访问到叶子节点
					   的祖父节点，因此直接把它及父亲节点删除即可。
					 3.删除后的任意叶子节点实际都需要访问到，且最终要回到原点，此时其实在任意点出发都需要把剩余的边
					   遍历两遍，最终答案为 2*remain
		 边界: 两个带金币叶子节点连上一条边，可能重复减掉边后出现负数，实际无需走任何边，答案为0。
		 两种情况取max(2*remain, 0)即可

 **/

func collectTheCoins(coins []int, edges [][]int) int {
	var n = len(coins)
	var g = make(map[int][]int)
	var degree = make([]int, n)
	for _, e := range edges {
		g[e[0]] = append(g[e[0]], e[1])
		g[e[1]] = append(g[e[1]], e[0])
		degree[e[0]]++
		degree[e[1]]++
	}
	// 1.第一遍拓扑排序，去除无金币的叶子节点
	var q = make([]int, 0)
	for i := 0; i < n; i++ {
		if degree[i] == 1 && coins[i] == 0 {
			q = append(q, i)
		}
	}
	var remain = n - 1
	for len(q) > 0 {
		var cur = q[0]
		q = q[1:]
		degree[cur]--
		remain--
		for _, nxt := range g[cur] {
			degree[nxt]--
			if degree[nxt] == 1 && coins[nxt] == 0 {
				q = append(q, nxt)
			}
		}
	}
	// 2.第二遍, 去除有金币的叶子节点及其父节点(能搜集距离为2的金币)
	for x := 0; x < 2; x++ {
		var leaves = make([]int, 0)
		for i := 0; i < n; i++ {
			if degree[i] == 1 {
				leaves = append(leaves, i)
			}
		}
		for _, c := range leaves {
			remain--
			degree[c]--
			for _, nxt := range g[c] {
				degree[nxt]--
			}
		}
	}
	return max(2*remain, 0)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
