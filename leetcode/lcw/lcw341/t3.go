package lcw341

func addMinimum(word string) int {
	var cnt = 0
	var res = 0
	var calc = func(v int32) {
		if v == 'a' {
			if cnt != 0 {
				res += 3 - cnt
			}
			cnt = 1
		} else if v == 'b' {
			if cnt == 2 {
				res += 2
			} else if cnt == 0 {
				res += 1
			}
			cnt = 2
		} else if v == 'c' {
			if cnt != 2 {
				res += 2 - cnt
			}
			cnt = 0
		}
	}
	for _, v := range word {
		calc(v)
	}
	if cnt != 0 {
		calc('a')
	}
	return res
}
