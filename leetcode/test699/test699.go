package test699

func fallingSquares(positions [][]int) []int {
	var n = len(positions)
	var res = make([]int, n)
	for i, pos := range positions {
		var curW = pos[1]
		curL, curR := pos[0], pos[0]+curW-1
		res[i] = curW
		for j := 0; j < i; j++ {
			var preW = positions[j][1]
			preL, preR := positions[j][0], positions[j][0]+preW-1
			if curR >= preL && preR >= curL {
				res[i] = max(res[i], res[j]+curW)
			}
		}
	}
	for i := 1; i < n; i++ {
		res[i] = max(res[i], res[i-1])
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
