package main

var INF = int(1e9)
var dirs = [][]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}
var pathMap = map[int]byte{
	0: 'v', 1: '>', 2: '^', 3: '<',
}

func conveyorBelt(matrix []string, start []int, end []int) int {
	m, n := len(matrix), len(matrix[0])
	var dis = make([][]int, m)
	for i := 0; i < m; i++ {
		dis[i] = make([]int, n)
		for j := 0; j < n; j++ {
			dis[i][j] = INF
		}
	}
	dis[start[0]][start[1]] = 0
	var q = make([][]int, 0)
	q = append(q, start)
	for len(q) > 0 {
		var cur = q[0]
		q = q[1:]
		for i := 0; i < len(dirs); i++ {
			nx, ny := cur[0]+dirs[i][0], cur[1]+dirs[i][1]
			if nx < 0 || ny < 0 || nx >= m || ny >= n {
				continue
			}
			var cost = 0 // 方向不同 花费加1
			if matrix[cur[0]][cur[1]] != pathMap[i] {
				cost = 1
			}
			// 松弛后的边才加入下一轮判定
			if dis[cur[0]][cur[1]]+cost < dis[nx][ny] {
				dis[nx][ny] = dis[cur[0]][cur[1]] + cost
				if cost == 0 {
					var nq = make([][]int, 0)
					nq = append(nq, []int{nx, ny})
					nq = append(nq, q...)
					q = nq
				} else {
					q = append(q, []int{nx, ny})
				}
			}
		}
	}
	return dis[end[0]][end[1]]
}
