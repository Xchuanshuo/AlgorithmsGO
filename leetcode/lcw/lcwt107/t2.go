package lcwt107

func longestString(x int, y int, z int) int {
	var mem = make([][][][]int, x+1)
	for i := 0; i <= x; i++ {
		mem[i] = make([][][]int, y+1)
		for j := 0; j <= y; j++ {
			mem[i][j] = make([][]int, z+1)
			for k := 0; k <= z; k++ {
				mem[i][j][k] = make([]int, 4)
				for e := 0; e < len(mem[i][j][k]); e++ {
					mem[i][j][k][e] = -1
				}
			}
		}
	}
	var dfs func(x, y, z int, pre int) int
	dfs = func(x, y, z int, pre int) int {
		if mem[x][y][z][pre] != -1 {
			return mem[x][y][z][pre]
		}
		var res = 0
		if x > 0 && pre != 1 {
			res = max(res, 2+dfs(x-1, y, z, 1))
		}
		if y > 0 && pre != 2 && pre != 3 {
			res = max(res, 2+dfs(x, y-1, z, 2))
		}
		if z > 0 && pre != 1 {
			res = max(res, 2+dfs(x, y, z-1, 3))
		}
		mem[x][y][z][pre] = res
		return res
	}
	return dfs(x, y, z, 0)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
