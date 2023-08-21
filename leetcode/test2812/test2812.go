package test2812

import "math"

var dirs = [][]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}

func maximumSafenessFactor(g [][]int) int {
	var n = len(g)
	var dis = make([][]int, n)
	var q = make([][]int, 0)
	for i := 0; i < n; i++ {
		dis[i] = make([]int, n)
		for j := 0; j < n; j++ {
			if g[i][j] == 1 {
				q = append(q, []int{i, j, 0})
			} else {
				dis[i][j] = math.MaxInt32 / 2
			}
		}
	}
	for len(q) > 0 {
		x, y, d := q[0][0], q[0][1], q[0][2]
		q = q[1:]
		for _, dir := range dirs {
			nx, ny, nd := x+dir[0], y+dir[1], d+1
			if nx < 0 || nx >= n || ny < 0 || ny >= n || nd >= dis[nx][ny] {
				continue
			}
			dis[nx][ny] = nd
			q = append(q, []int{nx, ny, nd})
		}
	}
	var check = func(limit int) bool {
		if dis[0][0] < limit {
			return false
		}
		var q = make([][]int, 0)
		var vis = make([][]bool, n)
		for i := 0; i < n; i++ {
			vis[i] = make([]bool, n)
		}
		vis[0][0] = true
		q = append(q, []int{0, 0})
		for len(q) > 0 {
			x, y := q[0][0], q[0][1]
			q = q[1:]
			for _, dir := range dirs {
				nx, ny := x+dir[0], y+dir[1]
				if nx < 0 || nx >= n || ny < 0 || ny >= n || vis[nx][ny] || dis[nx][ny] < limit {
					continue
				}
				if nx == n-1 && ny == n-1 {
					return true
				}
				vis[nx][ny] = true
				q = append(q, []int{nx, ny})
			}
		}
		return false
	}
	l, r := 0, n+n
	for l <= r {
		var mid = l + (r-l)/2
		if !check(mid) {
			r = mid - 1
		} else {
			if mid == n+n || !check(mid+1) {
				return mid
			}
			l = mid + 1
		}
	}
	return 0
}
