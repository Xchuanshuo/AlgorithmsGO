package test921

func minAddToMakeValid(str string) int {
	var s = make([]int32, 0)
	for _, v := range str {
		if v == '(' {
			s = append(s, v)
		} else {
			if len(s) == 0 || s[len(s)-1] != '(' {
				s = append(s, v)
			} else {
				s = s[0 : len(s)-1]
			}
		}
	}
	return len(s)
}
