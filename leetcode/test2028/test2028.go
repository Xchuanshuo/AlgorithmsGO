package test2028

func missingRolls(rolls []int, mean int, n int) []int {
	var m = len(rolls)
	var sum = mean * (n + m)
	for i := 0; i < m; i++ {
		if rolls[i] < 0 || rolls[i] > 6 {
			return []int{}
		}
		sum -= rolls[i]
	}
	if sum > n*6 || sum < n {
		return []int{}
	}
	var res = make([]int, n)
	var remain = sum % n
	for i := 0; i < remain; i++ {
		res[i] = 1
	}
	for i := 0; i < n; i++ {
		res[i] += sum / n
	}
	return res
}
