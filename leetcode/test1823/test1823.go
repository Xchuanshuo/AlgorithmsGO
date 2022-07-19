package test1823

func findTheWinner(n int, k int) int {
	if n == 1 {
		return 1
	}
	var val = (findTheWinner(n-1, k) + k) % n
	if val == 0 {
		return n
	}
	return val
}
