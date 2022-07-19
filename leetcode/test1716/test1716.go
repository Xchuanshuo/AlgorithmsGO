package test1716

func totalMoney(n int) int {
	res, cur, next := 0, 1, 1
	for i := 1;i <= n;i++ {
		res += cur
		if i%7 == 0 {
			cur = next
			next += 1
		}
		cur += 1
	}
	return res
}