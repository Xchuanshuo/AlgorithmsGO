package test2258

/**
 * @Description https://leetcode.cn/problems/escape-the-spreading-fire/
 * idea: 二分+bfs+dfs
		1.可等待的较长时间能走到终点较短也一定能，显然可等待时间t是满足单调性的
		2.二分枚举等待时间t, 首先模拟着火格子往外扩散t秒，使用bfs即可；接着dfs模拟从起点出发走到终点，
		 由于火会随着时间一直扩散，所以bfs过程中可以计算t秒过后，到达剩余格子的最早时间。
		 若走到当前格子花费的时间比火势扩散到该格子的时间要早，则说明可以走到该格子，以此类推。
		 对于走到终点后火也刚好到达做一个特判即可。
 **/

var INF = int(1e9)
var dirs = [][]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}

func maximumMinutes(g [][]int) int {
	m, n := len(g), len(g[0])
	var check = func(t int) bool {
		var vis = make([][]bool, m)
		var gv = make([][]int, m)
		for i := 0; i < len(vis); i++ {
			vis[i] = make([]bool, n)
			gv[i] = make([]int, n)
		}
		var q = make([][]int, 0)
		for i := 0; i < m; i++ {
			for j := 0; j < n; j++ {
				if g[i][j] == 1 {
					q = append(q, []int{i, j})
					vis[i][j] = true
					gv[i][j] = -1
				}
			}
		}
		var level = 0
		// 1.模拟等待t秒
		for len(q) > 0 {
			level++
			var sz = len(q)
			for i := 0; i < sz; i++ {
				var cur = q[0]
				q = q[1:]
				for _, d := range dirs {
					nx, ny := cur[0]+d[0], cur[1]+d[1]
					if nx < 0 || nx >= m || ny < 0 || ny >= n ||
						g[nx][ny] == 2 || vis[nx][ny] {
						continue
					}
					if level > t {
						gv[nx][ny] = level - t
					} else {
						gv[nx][ny] = -1
					}
					vis[nx][ny] = true
					q = append(q, []int{nx, ny})
				}
			}
		}
		// 2.等待t秒后往终点走
		var dfs func(x, y, dis int) bool
		dfs = func(x, y, dis int) bool {
			if x == m-1 && y == n-1 {
				return true
			}
			vis[x][y] = true
			gv[x][y] = dis
			for _, d := range dirs {
				nx, ny := x+d[0], y+d[1]
				if nx < 0 || nx >= m || ny < 0 || ny >= n || g[nx][ny] == 2 {
					continue
				}
				if nx == m-1 && ny == n-1 && dis+1 == gv[nx][ny] { // 特判
					return true
				}
				if vis[nx][ny] && dis+1 >= gv[nx][ny] {
					continue
				}
				if dfs(nx, ny, dis+1) {
					return true
				}
			}
			return false
		}
		return dfs(0, 0, 0)
	}
	l, r := 0, INF
	for l <= r {
		var mid = l + (r-l)/2
		if !check(mid) {
			r = mid - 1
		} else {
			if mid == INF || !check(mid+1) {
				return mid
			}
			l = mid + 1
		}
	}
	return -1
}
