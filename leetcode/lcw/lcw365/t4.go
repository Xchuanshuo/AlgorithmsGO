package lcw365

/**
 * @Description https://leetcode.cn/contest/weekly-contest-365/problems/count-visited-nodes-in-a-directed-graph/
 * idea: 内向基环树 1.拓补排序去除叶子节点 2.找环 环内节点的答案即为环大小
				3.在环上建立反向图对树枝进行遍历，树枝答案为出发的环大小 + 树枝到环的路径长度
 **/

func countVisitedNodes(edges []int) []int {
	var n = len(edges)
	var degree = make([]int, n)
	var rg = make(map[int][]int)
	for i, f := range edges {
		degree[f]++
		rg[f] = append(rg[f], i) // 建立反图
	}
	var q = make([]int, 0)
	for i := 0; i < n; i++ {
		if degree[i] == 0 {
			q = append(q, i)
		}
	}
	// 1.拓扑排序去除叶子节点
	var res = make([]int, n)
	var level = 1
	for len(q) > 0 {
		var sz = len(q)
		for i := 0; i < sz; i++ {
			var cur = q[0]
			q = q[1:]
			var nxt = edges[cur]
			degree[nxt]--
			if degree[nxt] == 0 {
				q = append(q, nxt)
			}
		}
		level++
	}
	// 2.找环 入度>0的节点为环的路口
	var vis = make([]bool, n)
	for i := 0; i < n; i++ {
		if degree[i] == 0 {
			continue
		}
		ringSize, cur := 0, i
		var nodes = make([]int, 0)
		for degree[cur] > 0 {
			nodes = append(nodes, cur)
			vis[cur] = true
			degree[cur]--
			cur = edges[cur]
			ringSize++
		}
		var rq = make([]int, 0)
		for _, v := range nodes { // 环内每个节点的大小计入答案
			res[v] = ringSize
			for _, nxt := range rg[v] {
				if vis[nxt] {
					continue
				}
				rq = append(rq, nxt)
				vis[nxt] = true
			}
		}
		var level = 1
		for len(rq) > 0 {
			var sz = len(rq)
			for i := 0; i < sz; i++ {
				var c = rq[0]
				res[c] = level + ringSize
				rq = rq[1:]
				for _, nxt := range rg[c] {
					if vis[nxt] {
						continue
					}
					rq = append(rq, nxt)
					vis[nxt] = true
				}
			}
			level++
		}
	}
	return res
}
