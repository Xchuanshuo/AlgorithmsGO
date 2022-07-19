package main

import "fmt"

var n, max = 0, 0

func main() {
	fmt.Scanln(&n)
	var grid = make([][]byte, n)
	for i := 0;i < n;i++ {
		fmt.Scanln(&grid[i])
	}
	visited := make([][]bool, n)
	for i := 0;i < n;i++ { visited[i] = make([]bool, n) }
	dfs(0, 0, 0, 0, ' ', grid, visited)
	fmt.Println(max)
}

func dfs(x, y, l, r int, last byte, grid [][]byte, visited [][]bool) {
	if x<0 || x>=n || y<0 || y>=n || visited[x][y] || r > l ||
		(last==')'&&grid[x][y]=='(') {
		return
	}
	if grid[x][y] == '(' {
		l++
	} else {
		r++
	}
	if l == r {
		if l+r > max { max = l + r }
	}
	visited[x][y] = true
	dfs(x + 1, y, l, r, grid[x][y], grid, visited)
	dfs(x, y + 1, l, r, grid[x][y], grid, visited)
	dfs(x - 1, y, l, r, grid[x][y], grid, visited)
	dfs(x, y - 1, l, r, grid[x][y], grid, visited)
	visited[x][y] = false
}
