package test2101

func maximumDetonation(bombs [][]int) int {
	var n = len(bombs)
	var check = func(i, j int) bool {
		x1, y1, r1 := bombs[i][0], bombs[i][1], bombs[i][2]
		x2, y2, _ := bombs[j][0], bombs[j][1], bombs[j][2]
		var dis = int64((x1-x2)*(x1-x2)) + int64((y1-y2)*(y1-y2))
		return dis <= int64(r1*r1)
	}
	var f = make([]int, 0)
	var cnt = 0
	var dfs func(cur, color int)
	dfs = func(cur, color int) {
		f[cur] = color
		cnt++
		for i := 0; i < n; i++ {
			if f[i] != 0 || !check(cur, i) {
				continue
			}
			dfs(i, color)
		}
	}
	var res = 0
	for i := 0; i < n; i++ {
		f = make([]int, n)
		cnt = 0
		dfs(i, 1)
		res = max(res, cnt)
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
