package test1442

func maxScore(s string) int {
	var to = 0
	for _, v := range s {
		if v == '1' {
			to++
		}
	}
	zero, one := 0, 0
	var res = 0
	for idx, v := range s {
		if idx == len(s)-1 {
			break
		}
		if v == '0' {
			zero++
		} else {
			one++
		}
		var cur = zero + to - one
		if cur > res {
			res = cur
		}
	}
	return res
}
