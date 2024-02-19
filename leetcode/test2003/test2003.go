package test2003

/**
 * @Description https://leetcode.cn/problems/smallest-missing-genetic-value-in-each-subtree
 * idea: 1.当前节点只受子节点影响，所以只要子节点不存在值为1的节点，最小基因值【一定】是1
		  转换思路，只需单独计算基因值为1的节点及其祖先节点的最小基因值即可
		 2.当访问1节点-根节点的路径时，需要查看路径上的每个节点可以访问到的后代节点，dfs遍历即可。
		  由于每个节点最多访问一次，所以整体时间复杂度为O(n)
 **/

func smallestMissingValueSubtree(parents []int, nums []int) []int {
	var n = len(parents)
	var g = make(map[int][]int)
	var res = make([]int, n)
	for i, p := range parents {
		g[p] = append(g[p], i)
		res[i] = 1
	}
	var idx = 0
	for i := 0; i < n; i++ {
		if nums[i] == 1 { // 1及其祖先节点需要计算，其它节点最小基因值都为1
			idx = i
			break
		}
	}
	var vis = make([]bool, n+2)
	var pos = 1
	var dfs func(cur int)
	dfs = func(cur int) {
		vis[nums[cur]] = true
		for _, nxt := range g[cur] {
			if vis[nums[nxt]] {
				continue
			}
			dfs(nxt)
		}
	}
	for c := idx; c != -1; c = parents[c] {
		if !vis[nums[c]] {
			dfs(c)
		}
		// 查找当前第一个未访问到的基因值
		for pos <= n+1 && vis[pos] {
			pos++
		}
		res[c] = pos
	}
	return res
}
