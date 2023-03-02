package lcw332

func substringXorQueries(s string, queries [][]int) [][]int {
	var exists = make(map[int]bool)
	for _, q := range queries {
		exists[q[0]^q[1]] = true
	}
	var posMap = make(map[int][]int)
	for i := 0; i < len(s); i++ {
		var cur = 0
		if s[i] == '0' {
			if posMap[0] == nil {
				posMap[0] = []int{i, i}
			}
			continue
		}
		for j := i; j < min(j+30, len(s)); j++ {
			cur = cur*2 + int(s[j]-'0')
			if !exists[cur] || posMap[cur] != nil {
				continue
			}
			posMap[cur] = []int{i, j}
		}
	}
	var res = make([][]int, len(queries))
	for i, q := range queries {
		res[i] = []int{-1, -1}
		var t = q[0] ^ q[1]
		if v, exist := posMap[t]; exist {
			res[i] = v
		}
	}
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
