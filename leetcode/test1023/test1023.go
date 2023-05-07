package test1023

func camelMatch(queries []string, pattern string) []bool {
	var check = func(s, t string) bool {
		i, j := 0, 0
		for i < len(s) && j < len(t) {
			if s[i] == t[j] {
				i++
				j++
				continue
			} else if s[i] >= 'A' && s[i] <= 'Z' {
				return false
			}
			i++
		}
		for i < len(s) {
			if s[i] >= 'A' && s[i] <= 'Z' {
				return false
			}
			i++
		}
		return j == len(t)
	}
	var res = make([]bool, 0)
	for _, q := range queries {
		res = append(res, check(q, pattern))
	}
	return res
}
