package lcw315

import "strconv"

func sumOfNumberAndReverse(num int) bool {
	if num == 0 {
		return true
	}
	for i := 1; i <= num; i++ {
		var s = []byte(strconv.Itoa(i))
		for j := 0; j < len(s)/2; j++ {
			s[j], s[len(s)-j-1] = s[len(s)-j-1], s[j]
		}
		v, _ := strconv.Atoi(string(s))
		if i+v == num {
			return true
		}
	}
	return false
}
