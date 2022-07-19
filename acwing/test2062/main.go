package main

import "fmt"

var n, m int

func main() {
	fmt.Scanln(&n, &m)
	data := make([][]byte, n)
	for i := 0;i < n;i++ {
		fmt.Scanln(&data[i])
	}
	//fmt.Println(data)
	fmt.Println(solve(data))
}

func dfs(x, y int, data [][]byte) {
	if x<0 || x>=n || y<0 || y>=m || data[x][y]!='X' {
		return
	}
	data[x][y] = '0'
	dfs(x+1, y, data)
	dfs(x, y+1, data)
	dfs(x-1, y, data)
	dfs(x, y-1, data)
}

var dirs = [][]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}

func solve(data [][]byte) int {
	isFindFirst := false
	for i := 0;i < n;i++ {
		for j := 0;j < m;j++ {
			if data[i][j] == 'X' {
				dfs(i, j, data)
				isFindFirst = true
				break
			}
		}
		if isFindFirst { break }
	}
	visited := make([][]bool, n)
	for i := 0;i < n;i++ {
		visited[i] = make([]bool, m)
	}
	q := make([][]int, 0)
	for i := 0;i < n;i++ {
		for j := 0;j < m;j++ {
			if data[i][j] == '0' {
				q = append(q, []int{i, j})
				visited[i][j] = true
			}
		}
	}
	step := 0
	for len(q) > 0 {
		size := len(q)
		step++
		for k := 0;k < size;k++{
			cur := q[0]
			q = q[1:]
			for i := 0;i < len(dirs);i++ {
				nx, ny := cur[0]+dirs[i][0], cur[1]+dirs[i][1]
				if nx<0 || nx >= n || ny<0 || ny>=m || visited[nx][ny] {
					continue
				}
				if data[nx][ny] == 'X' {
					return step - 1
				}
				visited[nx][ny] = true
				q = append(q, []int{nx, ny})
			}
		}
	}
	return -1
}
