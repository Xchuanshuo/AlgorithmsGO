package lcwt98

import "strconv"

func minMaxDifference(num int) int {
	var s = strconv.Itoa(num)
	var t = int(s[0] - '0')
	var t2 = int(s[0] - '0')
	if t == 9 {
		for _, s1 := range s {
			var v = int(s1 - '0')
			if v != 9 {
				t = v
				break
			}
		}
	}
	max, min := 0, 0
	for _, s1 := range s {
		var v = int(s1 - '0')
		if v == t {
			max = max*10 + 9
		} else {
			max = max*10 + v
		}
		if v == t2 {
			min = min*10 + 0
		} else {
			min = min*10 + v
		}
	}
	return max - min
}
