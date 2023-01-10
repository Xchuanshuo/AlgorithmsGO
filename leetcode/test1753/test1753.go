package test1753

func maximumScore(a int, b int, c int) int {
	if a > b+c {
		return b + c
	}
	if b > a+c {
		return a + c
	}
	if c > a+b {
		return a + b
	}
	return (a + b + c) / 2
}
