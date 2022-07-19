package test1036


var (
	limit int
	sx,sy,tx,ty int
)

func isEscapePossible(blocked [][]int, source []int, target []int) bool {
	n := len(blocked)
	blockMap := make(map[int]bool)
	for _, b := range blocked {
		blockMap[b[0]*1e6+b[1]] = true
	}
	sx, sy, tx, ty = source[0], source[1], target[0], target[1]
	// 最多能包围的格子数(等差数列)
	limit = n*(n-1)/2
	res1 := dfs(source[0], source[1], 0, blockMap, make(map[int]bool), true)
	res2 := dfs(target[0], target[1], 0, blockMap, make(map[int]bool), false)
	return res1 && res2
}

func dfs(x, y, step int, block, visited map[int]bool, isFromSrc bool) bool {
	if x<0 || x>=1e6 || y<0 || y>=1e6 || block[x*1e6+y] || visited[x*1e6+y] {
		return false
	}
	visited[x*1e6+y] = true
	if step >= limit { return true }
	if isFromSrc {
		if x==tx && y==ty { return true}
	} else {
		if x==sx && y==sy { return true}
	}
	return dfs(x+1, y, step+1, block, visited, isFromSrc) ||
		dfs(x, y+1, step+1, block, visited, isFromSrc) ||
		dfs(x-1, y, step+1, block, visited, isFromSrc) ||
		dfs(x, y-1, step+1, block, visited, isFromSrc)
}