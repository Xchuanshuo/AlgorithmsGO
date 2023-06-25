package lcp41

var dirs = [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}, {1, 1}, {1, -1}, {-1, -1}, {-1, 1}}

func flipChess(board []string) int {
	m, n := len(board), len(board[0])
	var bfs = func(sx, sy int) int {
		var g = make([][]byte, m)
		for i, b := range board {
			g[i] = []byte(b)
		}
		g[sx][sy] = 'X'
		var check = func(x, y, dx, dy int) bool {
			x, y = x+dx, y+dy
			for x >= 0 && x < m && y >= 0 && y < n {
				if g[x][y] == '.' {
					return false
				}
				if g[x][y] == 'X' {
					return true
				}
				x, y = x+dx, y+dy
			}
			return false
		}
		var cnt = 0
		var q = [][]int{{sx, sy}}
		for len(q) > 0 {
			x, y := q[0][0], q[0][1]
			q = q[1:]
			for _, d := range dirs {
				if !check(x, y, d[0], d[1]) {
					continue
				}
				nx, ny := x+d[0], y+d[1]
				for nx >= 0 && nx < m && ny >= 0 && ny < n && g[nx][ny] != 'X' {
					g[nx][ny] = 'X'
					q = append(q, []int{nx, ny})
					cnt++
					nx, ny = nx+d[0], ny+d[1]
				}
			}
		}
		return cnt
	}
	var res = 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == '.' {
				res = max(res, bfs(i, j))
			}
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
