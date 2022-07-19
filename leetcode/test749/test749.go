package test749

var dirs = [][]int{{0, 1}, {1, 0}, {-1, 0}, {0, -1}}

func containVirus(isInfected [][]int) (ans int) {
	m, n := len(isInfected), len(isInfected[0])
	var bfs func() ([]int, []map[p]bool)
	bfs = func() ([]int, []map[p]bool) {
		var walls = make([]int, 0)
		var ns = make([]map[p]bool, 0)
		var id = 0
		for i := 0; i < m; i++ {
			for j := 0; j < n; j++ {
				if isInfected[i][j] != 1 {
					continue
				}
				id++
				var q = make([]p, 0)
				q = append(q, p{i, j})
				isInfected[i][j] = -id
				var neighbor = make(map[p]bool)
				var wal = 0
				for len(q) > 0 {
					var cur = q[0]
					q = q[1:]
					for k := 0; k < len(dirs); k++ {
						nx, ny := dirs[k][0]+cur.x, dirs[k][1]+cur.y
						if nx < 0 || nx >= m || ny < 0 || ny >= n {
							continue
						}
						if isInfected[nx][ny] == 1 {
							isInfected[nx][ny] = -id
							q = append(q, p{nx, ny})
						} else if isInfected[nx][ny] == 0 {
							neighbor[p{nx, ny}] = true
							wal++
						}
					}
				}
				walls = append(walls, wal)
				ns = append(ns, neighbor)
			}
		}
		return walls, ns
	}
	var res = 0
	for true {
		// 1.查找所有不同区域要安装的墙与邻居
		walls, ns := bfs()
		if len(ns) == 0 {
			break
		}
		// 2.找到最大邻居的区域
		var idx = 0
		for i := 1; i < len(ns); i++ {
			if len(ns[i]) > len(ns[idx]) {
				idx = i
			}
		}
		res += walls[idx]
		// 恢复感染区域
		for i := 0; i < m; i++ {
			for j := 0; j < n; j++ {
				if isInfected[i][j] >= 0 {
					continue
				}
				if isInfected[i][j] != -idx-1 {
					isInfected[i][j] = 1
				} else {
					isInfected[i][j] = 2
				}
			}
		}
		// 3.跳过需要安装墙的邻居 其它区域继续传染
		for id, v := range ns {
			if id == idx {
				continue
			}
			for k := range v {
				isInfected[k.x][k.y] = 1
			}
		}
	}
	return res
}

type p struct {
	x, y int
}
