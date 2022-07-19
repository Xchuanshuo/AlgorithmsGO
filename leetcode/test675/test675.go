package test675

import "sort"

var m, n int

func cutOffTree(forest [][]int) int {
	var arr = make([][]int, 0)
	m, n = len(forest), len(forest[0])
	for i := 0;i < m;i++ {
		for j := 0;j < n;j++ {
			if forest[i][j] > 1 {
				arr = append(arr, []int{i, j})
			}
		}
	}
	sort.Slice(arr, func(i, j int) bool {
		x1, y1 := arr[i][0], arr[i][1]
		x2, y2 := arr[j][0], arr[j][1]
		return forest[x1][y1] < forest[x2][y2]
	})
	var res = 0
	var last = []int{0, 0}
	for i := 0;i < len(arr);i++ {
		var cur = bfs(last[0], last[1], arr[i][0], arr[i][1], forest)
		last = arr[i]
		if cur == -1 {
			return -1
		}
		res += cur
	}
	return res
}

var dirs = [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
func bfs(x, y int, tx, ty int, forest [][]int) int {
	if x == tx && y == ty {
		return 0
	}
	var visited = make([][]bool, m)
	for i := 0;i < m;i++ {
		visited[i] = make([]bool, n)
	}
	visited[x][y] = true
	var q = make([][]int, 0)
	q = append(q, []int{x, y})
	var step = 0
	for len(q) > 0 {
		step++
		var size = len(q)
		for k := 0;k < size;k++ {
			var cur = q[0]
			q = q[1:]
			for i := 0;i < len(dirs);i++ {
				nx, ny := dirs[i][0] + cur[0], dirs[i][1] + cur[1]
				if nx < 0 || nx >= m || ny < 0 || ny >= n || forest[nx][ny] == 0 || visited[nx][ny] {
					continue
				}
				if nx == tx && ny == ty {
					return step
				}
				visited[nx][ny] = true
				q = append(q, []int{nx, ny})
			}
		}
	}
	return -1
}