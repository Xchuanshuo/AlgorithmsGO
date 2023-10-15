package lcw362

func isReachableAtTime(sx int, sy int, fx int, fy int, t int) bool {
	dx, dy := abs(sx-fx), abs(sy-fy)
	if sx == fx && sy == fy {
		return t != 1
	}
	return t >= max(dx, dy)
}

func max(a, b int) int {
	if a > b {
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
