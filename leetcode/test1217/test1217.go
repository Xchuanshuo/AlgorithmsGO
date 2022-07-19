package test1217

func minCostToMoveChips(position []int) int {
	odd, even := 0, 0
	for _, c := range position {
		if c%2 == 0 {
			even += 1
		} else {
			odd += 1
		}
	}
	return min(even, odd)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
