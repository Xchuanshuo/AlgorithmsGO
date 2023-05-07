package lcwt99

import (
	"sort"
	"strconv"
)

func splitNum(num int) int {
	var s = []byte(strconv.Itoa(num))
	var n = len(s)
	if n%2 == 1 {
		s = append(s, 0)
	}
	n = len(s)
	sort.Slice(s, func(i, j int) bool {
		return s[i] < s[j]
	})
	num1, num2 := 0, 0
	for i := 0; i < n; i += 2 {
		c1, c2 := int(s[i]-'0'), int(s[i-1]-'0')
		num1 = num1*10 + c1
		num2 = num2*10 + c2
	}
	return num1 + num2
}
