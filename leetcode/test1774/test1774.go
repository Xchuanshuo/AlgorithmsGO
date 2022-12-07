package test1774

func closestCost(b []int, t []int, target int) int {
	t = append(t, t...)
	m, n := len(b), len(t)
	var calc = func(s int) int {
		var cur = 0
		for i := 0; i < 31; i++ {
			if (1<<i)&s > 0 {
				cur += t[i]
			}
		}
		return cur
	}
	var dif = target
	var res = target
	for i := 1; i < 1<<m; i++ {
		if (b[i] - target) > dif {
			continue
		}
		for j := 0; j < 1<<n; j++ {
			var cur = b[i] + calc(j)
			var v = abs(cur - target)
			if v < dif || (v == dif && cur < res) {
				dif = v
				res = cur
			}
		}
	}
	return res
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}