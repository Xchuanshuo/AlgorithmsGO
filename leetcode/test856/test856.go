package test856

func scoreOfParentheses(str string) int {
	cnt, res := 0, 0
	for i, c := range str {
		if c == '(' {
			cnt++
		} else {
			cnt--
			if str[i-1] == '(' {
				res += 1 << cnt
			}
		}
	}
	return res
}
