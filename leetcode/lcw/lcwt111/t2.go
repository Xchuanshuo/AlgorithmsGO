package lcwt111

func canMakeSubsequence(s1 string, s2 string) bool {
	n1, n2 := len(s1), len(s2)
	if n1 < n2 {
		return false
	}
	i, j := 0, 0
	for i < n1 && j < n2 {
		var idx = (int(s2[j]-'a') - 1 + 26) % 26
		if s1[i] == s2[j] || s1[i] == byte('a'+idx) {
			j++
		}
		i++
	}
	return j == n2
}
