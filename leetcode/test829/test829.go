package test829

func consecutiveNumbersSum(n int) int {
	var res = 1
	for k := 1; k*k < n; k++ {
		if 2*n%k != 0 {
			continue
		}
		if (2*n/k-(k+1))%2 == 0 {
			res++
		}
	}
	return res
}
