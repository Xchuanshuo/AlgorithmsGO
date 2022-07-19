package test1614


func maxDepth(s string) int {
	max, cnt := 0, 0
	for _, v := range s {
		if v == '(' {
			cnt++
		} else if v == ')' {
			cnt--
		}
		if cnt > max { max = cnt }
	}
	return max
}