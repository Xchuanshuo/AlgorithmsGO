package test383


func canConstruct(ransomNote string, magazine string) bool {
	if len(ransomNote) > len(magazine) {
		return false
	}
	cntMap := make([]int, 26)
	for _, c := range magazine {
		cntMap[c-'a']++
	}
	for _, c := range ransomNote {
		cntMap[c-'a']--
		if cntMap[c-'a'] < 0 {
			return false
		}
	}
	return true
}
