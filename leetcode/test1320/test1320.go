package test1320

func minimumDistance(word string) int {
	var getDis = func(p1, p2 int) int {
		if p1 == 26 {
			return 0
		}
		x1, y1 := p1/6, p1%6
		x2, y2 := p2/6, p2%6
		return abs(x1-x2) + abs(y1-y2)
	}
	var mem = make([][27][27]int, len(word))
	for i := 0; i < len(mem); i++ {
		for j := 0; j < len(mem[i]); j++ {
			for k := 0; k < len(mem[i][j]); k++ {
				mem[i][j][k] = -1
			}
		}
	}
	var dfs func(k, i, j int) int
	dfs = func(k, i, j int) int {
		if k == len(word) {
			return 0
		}
		if mem[k][i][j] != -1 {
			return mem[k][i][j]
		}
		var pos = int(word[k] - 'A')
		var res = min(dfs(k+1, i, pos)+getDis(j, pos),
			dfs(k+1, pos, j)+getDis(i, pos))
		mem[k][i][j] = res
		return res
	}
	return dfs(0, 26, 26)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
