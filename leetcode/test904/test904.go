package test904

func totalFruit(fruits []int) int {
	n, res := len(fruits), 0
	var wind = make(map[int]int)
	l, r := 0, 0
	for r < n {
		var cur = fruits[r]
		wind[cur]++
		for len(wind) > 2 {
			left := fruits[l]
			wind[left]--
			if wind[left] == 0 {
				delete(wind, left)
			}
			l++
		}
		res = max(res, r-l+1)
		r++
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
