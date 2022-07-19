package test2038

func winnerOfGame(s string) bool {
	var cnt = 0
	for i := 0; i < len(s)-2; i++ {
		if s[i] == 'A' && s[i+1] == 'A' && s[i+2] == 'A' {
			cnt++
		}
		if s[i] == 'B' && s[i+1] == 'B' && s[i+2] == 'B' {
			cnt--
		}
	}
	return cnt > 0
}
