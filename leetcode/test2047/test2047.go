package test2047

import (
	"strings"
	"unicode"
)

func countValidWords(sentence string) int {
	res := 0
	for _, s := range strings.Fields(sentence) {
		if isValid(s) {
			res++
		}
	}
	return res
}

func isValid(s string) bool {
	isExist := false
	for i, ch := range s {
		if ch>='0'&&ch<='9' || (strings.ContainsRune("!.,", ch)&&i!=len(s)-1) {
			return false
		}
		if ch == '-' {
			if isExist || i==0 || i==len(s)-1 || !unicode.IsLower(rune(s[i-1])) || !unicode.IsLower(rune(s[i+1])) {
				return false
			}
			isExist = true
		}
	}
	return true
}