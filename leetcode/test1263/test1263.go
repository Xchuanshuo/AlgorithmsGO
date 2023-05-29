package test1263

/**
 * @Description https://leetcode.cn/problems/minimum-moves-to-move-a-box-to-their-target-location/
 * idea:
		 1.0-1BFS, 将移动箱子作为权值为1的路径，不移动箱子作为权值为0的路径，1放对尾，0放队首；
		 2.所有状态 人移动*箱子移动, 所以人和箱子的坐标可以组成一个key, 用来判断是否访问过；
		 3.人移动的新位置是箱子时，箱子才可移动，且箱子新位置与人移动方向相同，其它流程同bfs，注意边界即可。
 **/

var dirs = [][]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}

func minPushBox(g [][]byte) int {
	m, n := len(g), len(g[0])
	tx, ty := 0, 0
	sx, sy := 0, 0
	bx, by := 0, 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if g[i][j] == 'S' {
				sx, sy = i, j
			} else if g[i][j] == 'T' {
				tx, ty = i, j
			} else if g[i][j] == 'B' {
				bx, by = i, j
			}
		}
	}
	var visited = make([][]bool, m*n)
	for i := 0; i < len(visited); i++ {
		visited[i] = make([]bool, m*n)
	}
	var check = func(x, y int) bool {
		return x >= 0 && x < m && y >= 0 && y < n && g[x][y] != '#'
	}
	var q = make([][]int, 0)
	q = append(q, []int{sx, sy, bx, by, 0})
	visited[sx*n+sy][bx*n+by] = true
	for len(q) > 0 {
		var cur = q[0]
		q = q[1:]
		x, y := cur[0], cur[1]
		bx, by = cur[2], cur[3]
		if bx == tx && by == ty {
			return cur[4]
		}
		for _, d := range dirs {
			nx, ny := x+d[0], y+d[1]
			if !check(nx, ny) || visited[nx*n+ny][bx*n+by] {
				continue
			}
			nbx, nby := bx+d[0], by+d[1]
			si, bi, nbi := nx*n+ny, bx*n+by, nbx*n+nby
			if si != bi {
				q = append([][]int{{nx, ny, bx, by, cur[4]}}, q...)
				visited[si][bi] = true
			} else if check(nbx, nby) && !visited[si][nbi] {
				q = append(q, []int{nx, ny, nbx, nby, cur[4] + 1})
				visited[si][nbi] = true
			}
		}
	}
	return -1
}
