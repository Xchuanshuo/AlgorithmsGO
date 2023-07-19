package test2018

import "strings"

func placeWordInCrossword(board [][]byte, word string) bool {
	m, n := len(board), len(board[0])
	var bs = []byte(word)
	for i := 0; i < len(bs)/2; i++ {
		bs[i], bs[len(bs)-i-1] = bs[len(bs)-i-1], bs[i]
	}
	var rword = string(bs)
	var check = func(s1 string) bool {
		var valid = true
		for i := 0; i < len(s1); i++ {
			if s1[i] != ' ' && s1[i] != word[i] {
				valid = false
				break
			}
		}
		if valid {
			return true
		}
		valid = true
		for i := 0; i < len(s1); i++ {
			if s1[i] != ' ' && s1[i] != rword[i] {
				valid = false
				break
			}
		}
		return valid
	}
	for i := 0; i < m; i++ {
		ss := strings.Split(string(board[i]), "#")
		for _, s := range ss {
			if len(s) != len(word) {
				continue
			}
			if check(s) {
				return true
			}
		}
	}
	for j := 0; j < n; j++ {
		var ts = make([]byte, 0)
		for i := 0; i < m; i++ {
			ts = append(ts, board[i][j])
		}
		ss := strings.Split(string(ts), "#")
		for _, s := range ss {
			if len(s) != len(word) {
				continue
			}
			if check(s) {
				return true
			}
		}
	}
	return false
}
