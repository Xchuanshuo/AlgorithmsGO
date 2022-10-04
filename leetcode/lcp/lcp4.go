package main

/**
 * @Description https://leetcode.cn/problems/broken-board-dominoes/
 * idea: 二分图的最大匹配 匈牙利算法
 **/

func domino(n int, m int, broken [][]int) int {
	var mp = make(map[int]bool)
	for _, b := range broken {
		mp[b[0]*m+b[1]] = true
	}
	var g = make(map[int][]int)
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			k1, k2, k3 := i*m+j, i*m+j+1, (i+1)*m+j
			if j+1 < m && !mp[k1] && !mp[k2] {
				g[k1] = append(g[k1], k2)
				g[k2] = append(g[k2], k1)
			}
			if i+1 < n && !mp[k1] && !mp[k3] {
				g[k1] = append(g[k1], k3)
				g[k3] = append(g[k3], k1)
			}
		}
	}
	var matching = make([]int, n*m)
	var visited = make(map[int]bool)
	for i := 0; i < len(matching); i++ {
		matching[i] = -1
	}
	var dfs func(v int) bool
	dfs = func(v int) bool {
		visited[v] = true
		for _, w := range g[v] {
			if visited[w] {
				continue
			}
			visited[w] = true
			if matching[w] == -1 || dfs(matching[w]) {
				matching[v], matching[w] = w, v
				return true
			}
		}
		return false
	}
	var res = 0
	for k, _ := range g {
		visited = make(map[int]bool)
		if matching[k] == -1 && dfs(k) {
			res++
		}
	}
	return res
}

//func main() {
//	n, m := 2, 3
//	var broken = [][]int{{1, 0}, {1, 1}}
//	var res = domino(n, m, broken)
//	fmt.Println(res)
//}
