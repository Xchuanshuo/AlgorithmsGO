package test1970

import (
	"math"
)

// 二分 + bfs

var dirs = [][]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}

func latestDayToCross(row int, col int, cells [][]int) int {
	var cell = make([][]int, row)
	for i := 0; i < len(cell); i++ {
		cell[i] = make([]int, col)
		for j := 0; j < len(cell[i]); j++ {
			cell[i][j] = math.MaxInt32
		}
	}
	for i, v := range cells {
		cell[v[0]-1][v[1]-1] = i + 1
	}
	var check = func(t int) bool {
		var q = make([][]int, 0)
		var vis = make([][]bool, row)
		for i := 0; i < len(vis); i++ {
			vis[i] = make([]bool, col)
		}
		for i := 0; i < col; i++ {
			if cell[0][i] <= t {
				continue
			}
			q = append(q, []int{0, i})
			vis[0][i] = true
		}
		for len(q) > 0 {
			x, y := q[0][0], q[0][1]
			q = q[1:]
			if x == row-1 {
				return true
			}
			for _, d := range dirs {
				nx, ny := x+d[0], y+d[1]
				if nx < 0 || nx >= row || ny < 0 || ny >= col || vis[nx][ny] || cell[nx][ny] <= t {
					continue
				}
				vis[nx][ny] = true
				q = append(q, []int{nx, ny})
			}
		}
		return false
	}
	l, r := 1, len(cells)
	for l <= r {
		var mid = l + (r-l)/2
		if !check(mid) {
			r = mid - 1
		} else {
			if mid == len(cells) || !check(mid+1) {
				return mid
			}
			l = mid + 1
		}
	}
	return 0
}
