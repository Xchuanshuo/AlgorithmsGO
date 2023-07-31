package lcwt109

import (
	"sort"
	"unicode"
)

func sortVowels(s string) string {
	var a = make([]byte, 0)
	var isValid = func(b byte) bool {
		var c = unicode.ToUpper(rune(b))
		return c == 'A' || c == 'E' || c == 'I' || c == 'O' || c == 'U'
	}
	for i := 0; i < len(s); i++ {
		if isValid(s[i]) {
			a = append(a, s[i])
		}
	}
	sort.Slice(a, func(i, j int) bool {
		return int(a[i]) < int(a[j])
	})
	var bs = []byte(s)
	var k = 0
	for i := 0; i < len(s); i++ {
		if isValid(s[i]) {
			bs[i] = a[k]
			k++
		}
	}
	return string(bs)
}
