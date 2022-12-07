package lcw83

func uniqueLetterString(s string) int {
	var res = 0
	var idxMp = make([][]int, 26)
	for i := 0; i < len(idxMp); i++ {
		idxMp[i] = []int{-1}
	}
	for i, v := range s {
		var p = v - 'A'
		if len(idxMp[p]) == 1 {
			idxMp[p] = append(idxMp[p], i)
		} else {
			idxMp[p][0], idxMp[p][1] = idxMp[p][1], i
		}
		for j := 0; j < 26; j++ {
			if len(idxMp[j]) == 2 {
				res += idxMp[j][1] - idxMp[j][0]
			}
		}
	}
	return res
}
