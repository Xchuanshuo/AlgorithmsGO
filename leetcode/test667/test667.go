package test667

func constructArray(n int, k int) []int {
	var res = make([]int, 0)
	k = n - k
	for i := 1; i < k; i++ {
		res = append(res, i)
	}
	l, r := k, n
	var flag = true
	for i := k; i <= n; i++ {
		if flag {
			res = append(res, r)
			r--
		} else {
			res = append(res, l)
			l++
		}
		flag = !flag
	}
	return res
}
