package main

import "fmt"

func main() {
	var n, m int
	fmt.Scanln(&n, &m)
	var mp = make([]map[rune]bool, m)
	for i := 0; i < n; i++ {
		var str string
		fmt.Scanln(&str)
		for idx, ch := range str {
			if mp[idx] == nil {
				mp[idx] = make(map[rune]bool)
			}
			mp[idx][ch] = true
		}
	}
	tmp := make([]bool, m)
	for i := 0; i < m; i++ {
		tmp[i] = true
	}
	for i := 0; i < n; i++ {
		var str string
		fmt.Scanln(&str)
		for idx, ch := range str {
			if mp[idx][ch] && tmp[idx] {
				tmp[idx] = false
			}
		}
	}
	res := 0
	for i := 0; i < m; i++ {
		if tmp[i] {
			res++
		}
	}
	fmt.Println(res)
}
