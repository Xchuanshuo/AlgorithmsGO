package test942

func diStringMatch(s string) []int {
	var n = len(s)
	l, r := 0, n
	var res = make([]int, 0)
	for _, c := range s {
		if c == 'I' {
			res = append(res, l)
			l++
		} else {
			res = append(res, r)
			r--
		}
	}
	res = append(res, l)
	return res
}
