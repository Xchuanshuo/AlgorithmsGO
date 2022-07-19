package main

import "fmt"

func findSubstring(s string, words []string) []int {
	var mp = make(map[string]int)
	for _, w := range words {
		mp[w]++
	}
	var res = make([]int, 0)
	m, n := len(words[0]), len(words)
	for i := 0; i < len(s); i++ {
		if i+m*n > len(s) {
			break
		}
		var cloneMap = clone(mp)
		r, valid := i+m*n, 0
		for j := i; j < r; {
			var sub = s[j : j+m]
			if cloneMap[sub] == 0 {
				break
			}
			cloneMap[sub]--
			j += m
			valid++
		}
		if valid == n {
			res = append(res, i)
		}
	}
	return res
}

func clone(mp map[string]int) map[string]int {
	var newMap = make(map[string]int)
	for k, v := range mp {
		newMap[k] = v
	}
	return newMap
}

func main() {
	var word = "goodgoodbestword"
	var arr = []string{"word", "good", "best", "good"}
	res := findSubstring(word, arr)
	fmt.Println(res)
}
