package main

import (
	"fmt"
	"strings"
)

func complexNumberMultiply(num1 string, num2 string) string {
	a1, a2 := strings.Split(num1, "+"), strings.Split(num2, "+")
	virtual, real := 0, 0
	for _, s1 := range a1 {
		for _, s2 := range a2 {
			v1, v2 := parseV(s1), parseV(s2)
			if strings.HasSuffix(s1, "i") && strings.HasSuffix(s2, "i") {
				real += v1 * v2 * -1
			} else if !strings.HasSuffix(s1, "i") && !strings.HasSuffix(s2, "i") {
				real += v1 * v2
			} else {
				virtual += v1 * v2
			}
		}
	}
	return fmt.Sprintf("%d+%di", real, virtual)
}

func parseV(s string) int {
	var t = 0
	var neg = s[0] == '-'
	for i := 0; i < len(s); i++ {
		if s[i] == 'i' {
			break
		}
		if s[i] == '-' {
			continue
		}
		t = t*10 + int(s[i]-'0')
	}
	if neg {
		return -t
	}
	return t
}
func main() {
	var s1 = "78+-76i"
	var s2 = "-86+72i"
	var res = complexNumberMultiply(s1, s2)
	fmt.Println(res)
}
