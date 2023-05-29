package lcw345

var dirs = [][]int{{1, 1}, {0, 1}, {-1, 1}}
func maxMoves(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	var q = make([][]int, 0)
	var visited = make([][]bool, m)
	for i := 0;i < m;i++ {
		q = append(q, []int{i, 0, 0})
		visited[i] = make([]bool, n)
		visited[i][0] = true
	}
	var res = 0
	for len(q) > 0 {
		var cur = q[0]
		q = q[1:]
		res = max(res, cur[2])
		if res >= n-1 {
			break
		}
		for _, d := range dirs {
			nx, ny := cur[0]+d[0], cur[1]+d[1]
			if nx<0 || nx >= m || ny < 0 || ny >= n || visited[nx][ny] || grid[nx][ny] <= grid[cur[0]][cur[1]] {
				continue
			}
			visited[nx][ny] = true
			q = append(q, []int{nx, ny, cur[2]+1})
		}
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
