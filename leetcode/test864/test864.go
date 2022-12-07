package test864

var dirs = [][]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}

func shortestPathAllKeys(grid []string) int {
	m, n := len(grid), len(grid[0])
	var target = 0
	var s = []int{0, 0, 0, 0}
	for i, g := range grid {
		for j, v := range g {
			if v >= 'a' && v <= 'z' {
				target |= 1 << (v - 'a')
			}
			if v == '@' {
				s = []int{i, j, 0}
			}
		}
	}
	var visited = make([][][]bool, m)
	for i := 0; i < m; i++ {
		visited[i] = make([][]bool, n)
		for j := 0; j < n; j++ {
			visited[i][j] = make([]bool, 128)
		}
	}
	var q = make([][]int, 0)
	q = append(q, s)
	visited[s[0]][s[1]][0] = true
	var step = 0
	for len(q) > 0 {
		var size = len(q)
		step++
		for i := 0; i < size; i++ {
			var cur = q[0]
			q = q[1:]
			for j := 0; j < len(dirs); j++ {
				nx, ny := cur[0]+dirs[j][0], cur[1]+dirs[j][1]
				if nx < 0 || nx >= m || ny < 0 || ny >= n || grid[nx][ny] == '#' {
					continue
				}
				var status = cur[2]
				if grid[nx][ny] >= 'A' && grid[nx][ny] <= 'Z' {
					var idx = grid[nx][ny] - 'A'
					if (status & (1 << idx)) == 0 {
						continue
					}
				}
				if grid[nx][ny] >= 'a' && grid[nx][ny] <= 'z' {
					var idx = grid[nx][ny] - 'a'
					status |= 1 << idx
				}
				if visited[nx][ny][status] {
					continue
				}
				if status == target {
					return step
				}
				visited[nx][ny][status] = true
				q = append(q, []int{nx, ny, status})
			}
		}
	}
	return -1
}
