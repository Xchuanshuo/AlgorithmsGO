package test2049

var max, cnt, n int
var nodes [][]int

func countHighestScoreNodes(parents []int) int {
	max, cnt, n = 0, 0, len(parents)
	nodes = make([][]int, n)
	for i := 0; i < n; i++ {
		nodes[i] = make([]int, 0)
	}
	for i := 0; i < len(parents); i++ {
		p := parents[i]
		if p == -1 {
			continue
		}
		nodes[p] = append(nodes[p], i)
	}
	dfs(0)
	return cnt
}

func dfs(node int) int {
	score, size := 1, n-1
	for _, c := range nodes[node] {
		var t = dfs(c)
		score *= t
		size -= t
	}
	if node != 0 {
		score *= size
	}
	if score == max {
		cnt++
	} else if score > max {
		max = score
		cnt = 1
	}
	return n - size
}
