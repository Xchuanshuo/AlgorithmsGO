package test1210

// 维度1.行、列，2.水平/竖直 0, 1
var dirs = [][]int{{0, 1, 0}, {1, 0, 0}, {0, 0, 1}}
func minimumMoves(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	var visited = make([][][]bool, m)
	for i := 0;i < len(visited);i++ {
		visited[i] = make([][]bool, n)
		for j := 0;j < len(visited[i]);j++ {
			visited[i][j] = make([]bool, 2)
		}
	}
	var paths = func(x, y, d int) [][]int {
		var list = make([][]int, 0)
		for i := 0;i < len(dirs);i++ {
			nx, ny, ns := x+dirs[i][0], y+dirs[i][1], d^dirs[i][2]
			tx, ty := nx+ns, ny+(ns^1) // 水平:nx,ny+1;竖直:nx+1,ny
			if nx >= m || ny >= n || grid[nx][ny]==1 || visited[nx][ny][ns] { continue }
			if tx >= m || ty >= n || grid[tx][ty]==1 { continue }
			if dirs[i][2]==1 && grid[nx+1][ny+1]==1 { continue }
			list = append(list, []int{nx, ny, ns})
			visited[nx][ny][ns] = true
		}
		//fmt.Println(x,y,d, list)
		return list
	}
	var q = make([][]int, 0)
	q = append(q, []int{0, 0, 0})
	visited[0][0][0] = true
	var step = 0
	for len(q) > 0 {
		var sz = len(q)
		step++
		for i := 0;i < sz;i++ {
			var cur = q[0]
			q = q[1:]
			for _, nxt := range paths(cur[0], cur[1], cur[2]) {
				if nxt[0] == m-1 && nxt[1] == n-2 {
					return step
				}
				q = append(q, nxt)
			}
		}
	}
	return -1
}
