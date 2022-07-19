package test390

func lastRemaining1(n int) int {
	l2r := true
	left, step := 1, 1
	for n > 1 {
		if l2r || n%2 == 1 {
			left += step
		}
		n /= 2
		step = 2 * step
		l2r = !l2r
	}
	return left
}


func lastRemaining(n int) int {
	if n == 1 {
		return 1
	}
	return 2 * (n/2+1 - lastRemaining(n/2))
}
