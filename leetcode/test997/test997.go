package test997

func findJudge(n int, trust [][]int) int {
	if len(trust) == 0 {
		if n == 1 {
			return 1
		}
		return -1
	}
	m := 1
	inMap, outMap := make(map[int]int), make(map[int]int)
	for _, t := range trust {
		v, w := t[0], t[1]
		outMap[v]++
		inMap[w]++
		if v > m {
			m = v
		}
		if w > m {
			m = w
		}
	}
	for k, v := range inMap {
		if v == m - 1 && outMap[k] == 0 {
			return k
		}
	}
	return -1
}