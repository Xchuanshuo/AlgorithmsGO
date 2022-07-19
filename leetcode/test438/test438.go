package test438


func findAnagrams(s string, p string) []int {
	m, n := len(s), 0
	wind := make([]int, 128)
	need := make([]int, 128)
	for _, v := range p {
		if need[v] == 0 { n++ }
		need[v]++
	}
	var result []int
	l, r, valid := 0, 0, 0
	for r < m {
		c := s[r]
		r++
		wind[c]++
		if wind[c] == need[c] {valid++ }
		for valid == n {
			if r - l == len(p) {
				result = append(result, l)
			}
			d := s[l]
			if wind[d] == need[d] {	valid-- }
			wind[d]--
			l++
		}
	}
	return result
}