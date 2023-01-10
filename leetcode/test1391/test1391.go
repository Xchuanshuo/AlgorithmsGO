package test1391

var pathTo = map[int][]map[int]bool{
	1: {{}, {}, {1: true, 4: true, 6: true}, {1: true, 3: true, 5: true}},
	2: {{2: true, 3: true, 4: true}, {2: true, 5: true, 6: true}, {}, {}},
	3: {{}, {2: true, 5: true, 6: true}, {1: true, 4: true, 6: true}, {}},
	4: {{}, {2: true, 5: true, 6: true}, {}, {1: true, 3: true, 5: true}},
	5: {{2: true, 3: true, 4: true}, {}, {1: true, 4: true, 6: true}, {}},
	6: {{2: true, 3: true, 4: true}, {}, {}, {1: true, 3: true, 5: true}},
}
var dirs = [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

func hasValidPath(grid [][]int) bool {
	m, n := len(grid), len(grid[0])
	var visited = make([][]bool, m)
	for i := 0; i < len(visited); i++ {
		visited[i] = make([]bool, n)
	}
	var q = [][]int{{0, 0}}
	visited[0][0] = true
	for len(q) > 0 {
		var cur = q[0]
		if cur[0] == m-1 && cur[1] == n-1 {
			return true
		}
		q = q[1:]
		for i := 0; i < len(dirs); i++ {
			nx, ny := cur[0]+dirs[i][0], cur[1]+dirs[i][1]
			if nx < 0 || nx >= m || ny < 0 || ny >= n || visited[nx][ny] {
				continue
			}
			curV, nxtV := grid[cur[0]][cur[1]], grid[nx][ny]
			if !pathTo[curV][i][nxtV] {
				continue
			}
			visited[nx][ny] = true
			q = append(q, []int{nx, ny})
		}
	}
	return false
}
