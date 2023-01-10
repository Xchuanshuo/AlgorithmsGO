package lcw324

func isPossible(n int, edges [][]int) bool {
	var hasEdge = make(map[int]map[int]bool)
	var degree = make([]int, n+1)
	for _, e := range edges {
		if _, exist := hasEdge[e[0]]; !exist {
			hasEdge[e[0]] = make(map[int]bool)
		}
		if _, exist := hasEdge[e[1]]; !exist {
			hasEdge[e[1]] = make(map[int]bool)
		}
		hasEdge[e[0]][e[1]] = true
		hasEdge[e[1]][e[0]] = true
		degree[e[0]]++
		degree[e[1]]++
	}
	var nodes = make([]int, 0)
	for i := 0; i < len(degree); i++ {
		if degree[i]%2 == 1 {
			nodes = append(nodes, i)
		}
	}
	if len(nodes) == 0 {
		return true
	}
	if len(nodes)%2 != 0 || len(nodes) > 4 {
		return false
	}
	var cals = make([][]int, 0)
	for i, v := range nodes {
		for j := i + 1; j < len(nodes); j++ {
			var w = nodes[j]
			if !hasEdge[v][w] {
				cals = append(cals, []int{v, w})
			}
		}
	}
	if len(nodes) == 2 && len(cals) == 0 {
		v, w := nodes[0], nodes[1]
		for i := 1; i <= n; i++ {
			if i != v && i != w && !hasEdge[i][v] && !hasEdge[i][w] {
				return true
			}
		}
		return false
	}
	var flag = false
	var dfs func(idx, cnt int)
	dfs = func(idx int, cnt int) {
		if cnt > 2 {
			return
		}
		var valid = true
		for _, node := range nodes {
			if degree[node]%2 != 0 {
				valid = false
				break
			}
		}
		if valid {
			flag = true
			return
		}
		for i := idx; i < len(cals); i++ {
			if flag {
				return
			}
			degree[cals[i][0]]++
			degree[cals[i][1]]++
			dfs(i+1, cnt+1)
			degree[cals[i][0]]--
			degree[cals[i][1]]--
		}
	}
	dfs(0, 0)
	return flag
}
