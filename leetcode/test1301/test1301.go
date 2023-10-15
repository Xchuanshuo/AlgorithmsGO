package test1301

var M = int(1e9) + 7
var dirs = [][]int{{1, 0}, {0, 1}, {1, 1}}

func pathsWithMaxScore(board []string) []int {
	m, n := len(board), len(board[0])
	var f = make([][]int, m+1)
	var g = make([][]int, m+1)
	for i := 0; i < len(f); i++ {
		f[i] = make([]int, n+1)
		g[i] = make([]int, n+1)
	}
	g[m-1][n-1] = 1
	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			if board[i][j] == 'X' || board[i][j] == 'S' {
				continue
			}
			var v = int(board[i][j] - '0')
			if board[i][j] == 'E' {
				v = 0
			}
			for _, d := range dirs {
				nx, ny := i+d[0], j+d[1]
				if nx < 0 || nx >= m || ny < 0 || ny >= n || board[nx][ny] == 'X' || g[nx][ny] == 0 {
					continue
				}
				if f[i][j] < f[nx][ny]+v {
					f[i][j] = f[nx][ny] + v
					g[i][j] = g[nx][ny]
				} else if f[i][j] == f[nx][ny]+v {
					g[i][j] = (g[i][j] + g[nx][ny]) % M
				}
			}
		}
	}
	return []int{f[0][0], g[0][0]}
}
