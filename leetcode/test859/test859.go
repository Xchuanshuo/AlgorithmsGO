package test859


func buddyStrings(s string, goal string) bool {
	cnt := 0
	if len(s) != len(goal) {
		return false
	}
	var cnts [256]int
	var last = -1
	flag := false
	for i, c := range s {
		cnts[c]++
		if int32(goal[i]) != c {
			cnt++
			if last != -1 && (s[last] != goal[i] || s[i] != goal[last]) {
				return false
			}
			last = i
		}
		if cnts[c] >= 2 {
			flag = true
		}

	}
	return cnt == 2 || (cnt == 0 && flag)
}
