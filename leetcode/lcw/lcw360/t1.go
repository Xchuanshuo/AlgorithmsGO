package lcw360

func furthestDistanceFromOrigin(moves string) int {
	cntL, cntR := 0, 0
	var n = len(moves)
	for _, v := range moves {
		if v == 'L' {
			cntL++
		} else if v == 'R' {
			cntR++
		}
	}
	var o = n - (cntL + cntR)
	return max(cntL+o-cntR, cntR+o-cntL)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
