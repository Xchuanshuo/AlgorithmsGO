package test1519

func countSubTrees(n int, edges [][]int, labels string) []int {
	var g = make(map[int][]int)
	for _, e := range edges {
		g[e[0]] = append(g[e[0]], e[1])
		g[e[1]] = append(g[e[1]], e[0])
	}
	var res = make([]int, n)
	var dfs func(fa, cur int) []int
	dfs = func(fa, cur int) []int {
		var cnts = make([]int, 26)
		cnts[int(labels[cur]-'a')]++
		for _, nxt := range g[cur] {
			if fa == nxt {
				continue
			}
			for i, v := range dfs(cur, nxt) {
				cnts[i] += v
			}
		}
		res[cur] = cnts[int(labels[cur]-'a')]
		return cnts
	}
	dfs(-1, 0)
	return res
}
