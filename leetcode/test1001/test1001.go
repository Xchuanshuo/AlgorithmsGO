package test1001

var dirs = [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}, {1, -1}, {-1, 1}, {1, 1}, {-1, -1}, {0, 0}}

func gridIllumination(n int, lamps [][]int, queries [][]int) []int {
	row, col := make(map[int]int), make(map[int]int)
	slash, reverseSlash := make(map[int]int), make(map[int]int)
	var points = make(map[int64]bool)
	for _, lamp := range lamps {
		h := hash(lamp[0], lamp[1])
		if points[h] {
			continue
		}
		row[lamp[0]]++
		col[lamp[1]]++
		slash[lamp[0]+lamp[1]]++
		reverseSlash[lamp[0]-lamp[1]]++
		points[h] = true
	}
	var res = make([]int, len(queries))
	for idx, q := range queries {
		x, y := q[0], q[1]
		var isLight = false
		if row[x] > 0 {
			isLight = true
		} else if col[y] > 0 {
			isLight = true
		} else if slash[x+y] > 0 {
			isLight = true
		} else if reverseSlash[x-y] > 0 {
			isLight = true
		}
		if isLight {
			res[idx] = 1
		}
		for i := 0; i < len(dirs); i++ {
			nx, ny := x+dirs[i][0], y+dirs[i][1]
			if nx < 0 || nx >= n || ny < 0 || ny >= n {
				continue
			}
			h := hash(nx, ny)
			if points[h] {
				row[nx]--
				col[ny]--
				slash[nx+ny]--
				reverseSlash[nx-ny]--
				points[h] = false
			}
		}
	}
	return res
}

func hash(x, y int) int64 {
	return int64(x) + int64(y)<<32
}
