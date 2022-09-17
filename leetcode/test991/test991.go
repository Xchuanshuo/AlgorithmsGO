package test991

func brokenCalc(x int, y int) int {
	if x >= y {
		return x - y
	}
	if y%2 == 0 {
		return 1 + brokenCalc(x, y/2)
	}
	return 1 + brokenCalc(x, y+1)
}
