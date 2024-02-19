package lcwt124

func lastNonEmptyString(s string) string {
	var pos = make([][]int, 26)
	for i, v := range s {
		var p = int(v - 'a')
		pos[p] = append(pos[p], i)
	}
	for true {
		var tmp = make([]int, 0)
		for i := 0; i < 26; i++ {
			if len(pos[i]) > 0 {
				tmp = append(tmp, pos[i][0])
				pos[i] = pos[i][1:]
			}
		}
		var last = true
		for i := 0; i < 26; i++ {
			if len(pos[i]) > 0 {
				last = false
				break
			}
		}
		if last {
			var bs = make([]byte, 0)
			for _, p := range tmp {
				bs = append(bs, s[p])
			}
			return string(bs)
		}
	}
	return ""
}
