package main

import (
	"fmt"
	"strings"
)

func letterCasePermutation(s string) []string {
	s = strings.ToLower(s)
	var arr = make([]int, 0)
	for i := 0; i < len(s); i++ {
		if s[i] >= 'a' && s[i] <= 'z' {
			arr = append(arr, i)
		}
	}
	var n = len(arr)
	if n == 0 {
		return []string{s}
	}
	var bytes = []byte(s)
	var res = make([]string, 0)
	var dfs func(idx int, lower bool)
	dfs = func(idx int, lower bool) {
		if idx == n {
			res = append(res, string(bytes))
			return
		}
		var t = bytes[arr[idx]]
		if !lower {
			bytes[arr[idx]] = s[arr[idx]] - 32
		} else {
			bytes[arr[idx]] = s[arr[idx]]
		}
		dfs(idx+1, lower)
		if idx != n-1 {
			dfs(idx+1, !lower)
		}
		bytes[arr[idx]] = t
	}
	dfs(0, true)
	dfs(0, false)
	return res
}

func main() {
	var str = "a1b2"
	var res = letterCasePermutation(str)
	fmt.Println(res)
}
