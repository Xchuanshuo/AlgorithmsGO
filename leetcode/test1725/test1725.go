package test1725

func countGoodRectangles(rectangles [][]int) int {
	maxL := 0
	for _, rec := range rectangles {
		cur := min(rec[0], rec[1])
		maxL = max(maxL, cur)
	}
	res := 0
	for _, rec := range rectangles {
		cur := min(rec[0], rec[1])
		if cur == maxL {
			res++
		}
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
