package test661

var dirs = [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}, {1, 1}, {1, -1}, {-1, 1}, {-1, -1}}

func imageSmoother(img [][]int) [][]int {
	m, n := len(img), len(img[0])
	var res = make([][]int, m)
	for i := 0; i < m; i++ {
		res[i] = make([]int, n)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			cnt, cur := 1, img[i][j]
			for k := 0; k < len(dirs); k++ {
				nx, ny := i+dirs[k][0], j+dirs[k][1]
				if nx < 0 || nx >= m || ny < 0 || ny >= n {
					continue
				}
				cnt++
				cur += img[nx][ny]
			}
			res[i][j] = cur / cnt
		}
	}
	return res
}
