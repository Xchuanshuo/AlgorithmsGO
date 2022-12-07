package lcw322

/**
 * @Description https://leetcode.cn/problems/divide-nodes-into-the-maximum-number-of-groups
 * idea: 1.节点按连通性分组
		 2.分组后计算每个组内最大分层方案(枚举)
		 3.所有最大分层相加即是最大答案
 **/

func magnificentSets(n int, edges [][]int) int {
	var g = make(map[int][]int)
	for _, e := range edges {
		g[e[0]] = append(g[e[0]], e[1])
		g[e[1]] = append(g[e[1]], e[0])
	}
	var calc = func(t int) int {
		var level = 1
		var nodes = []int{t}
		var mp = make(map[int]int)
		mp[t] = level
		for len(nodes) > 0 {
			var tArr = make([]int, 0)
			for _, node := range nodes {
				for _, nxt := range g[node] {
					if mp[nxt] == 0 {
						tArr = append(tArr, nxt)
						mp[nxt] = mp[node] + 1
					} else if abs(mp[nxt]-level) != 1 {
						return -1
					}
				}
			}
			level++
			nodes = tArr
		}
		return level - 1
	}
	var idMap = make(map[int][]int)
	var visited = make(map[int]bool)
	var dfs func(cur, id int)
	dfs = func(cur int, id int) {
		visited[cur] = true
		idMap[id] = append(idMap[id], cur)
		for _, nxt := range g[cur] {
			if visited[nxt] {
				continue
			}
			dfs(nxt, id)
		}
	}
	var id = 0
	for i := 1; i <= n; i++ {
		if visited[i] {
			continue
		}
		id++
		dfs(i, id)
	}
	var res = 0
	for _, nodes := range idMap {
		var v = -1
		for _, node := range nodes {
			val := calc(node)
			if val == -1 {
				return -1
			}
			v = max(v, val)
		}
		res += v
	}
	return res
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
