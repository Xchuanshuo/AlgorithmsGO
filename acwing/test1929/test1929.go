package main

import "fmt"

var n, m, res int
var grid [][]byte

func main() {
	fmt.Scanf("%d %d", &n, &m)
	grid = make([][]byte, n)
	for i := 0;i < n;i++ {
		var s string
		fmt.Scanf("%s", &s)
		grid[i] = []byte(s)
	}
	for i := 0;i < m;i++ {
		dfs(0, i, 1, 2)
		dfs(n-1, i, 1, 0)
	}

	for i := 0;i < n;i++ {
		dfs(i, 0, 1, 1)
		dfs(i, m-1, 1, 3)
	}
	fmt.Println(res)
}

var dirs = [][]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}
func dfs(x, y, l, dir int) {

	res = max(res, l)
	var d int
	if grid[x][y] == '/' {
		d = dir^1
	} else {
		d = dir^3
	}
	nx, ny := x + dirs[d][0], y + dirs[d][1]
	if nx<0 || nx>=n || ny<0 || ny>=m {
		return
	}
	dfs(nx, ny, l + 1, d)
}

func max(a, b int) int {
	if a > b { return a }
	return b
}