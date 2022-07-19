package test1446

func maxPower(s string) int {
	var last int32 = -1
	cnt, res := 0, 0
	for _, c := range s {
		if last == -1 || c == last {
			cnt++
		} else {
			cnt = 1
		}
		last = c
		res = max(res, cnt)
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}