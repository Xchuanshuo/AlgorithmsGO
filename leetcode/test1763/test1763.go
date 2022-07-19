package test1763

import "unicode"

func longestNiceSubstring(s string) string {
	res := ""
	for i := range s {
		low, up := 0, 0
		for j := i; j < len(s); j++ {
			if unicode.IsLower(rune(s[j])) {
				low |= 1 << (s[j] - 'a')
			} else {
				up |= 1 << (s[j] - 'A')
			}
			if low == up && j-i+1 > len(res) {
				res = s[i : j+1]
			}
		}
	}
	return res
}
