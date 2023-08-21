package lcw357

/**
 * @Description https://leetcode.cn/problems/find-the-safest-path-in-a-grid/
 * idea: 1.多源BFS求每个格子的安全系数
		 2.找到一条路径，安全系数最大，即安全系数<limit的路径不可走，二分答案即可
 **/

var dirs = [][]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}

func maximumSafenessFactor(g [][]int) int {
	var n = len(g)
	var dis = make([][]int, n)
	var bfs = func(q [][]int) {
		for len(q) > 0 {
			x, y, d := q[0][0], q[0][1], q[0][2]
			q = q[1:]
			for _, dir := range dirs {
				nx, ny := x+dir[0], y+dir[1]
				if nx < 0 || nx >= n || ny < 0 || ny >= n || dis[nx][ny] <= d+1 {
					continue
				}
				dis[nx][ny] = d + 1
				q = append(q, []int{nx, ny, d + 1})
			}
		}
	}
	var q = make([][]int, 0)
	for i := 0; i < n; i++ {
		dis[i] = make([]int, n)
		for j := 0; j < n; j++ {
			if g[i][j] == 1 {
				dis[i][j] = 0
				q = append(q, []int{i, j, 0})
			} else {
				dis[i][j] = 2 * n
			}
		}
	}
	// 1.多源BFS求每个格子的安全系数
	bfs(q)
	var check = func(limit int) bool {
		if dis[0][0] < limit || dis[n-1][n-1] < limit {
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
			if x == n-1 && y == n-1 {
				return true
			}
			q = q[1:]
			for _, dir := range dirs {
				nx, ny := x+dir[0], y+dir[1]
				if nx < 0 || nx >= n || ny < 0 || ny >= n || dis[nx][ny] < limit || vis[nx][ny] {
					continue
				}
				vis[nx][ny] = true
				q = append(q, []int{nx, ny})
			}
		}
		return false
	}
	l, r := 0, max(dis[0][0], dis[n-1][n-1])
	// 2.二分答案
	for l <= r {
		var mid = l + (r-l)/2
		if !check(mid) {
			r = mid - 1
		} else {
			if mid == max(dis[0][0], dis[n-1][n-1]) || !check(mid+1) {
				return mid
			}
			l = mid + 1
		}
	}
	return 0
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
