package main

import "fmt"

func main() {
	var n int
	fmt.Scanf("%d", &n)
	var str string
	fmt.Scanf("%s", &str)
	for l := 1; l <= n; l++ {
		var mp = make(map[string]bool)
		var valid = true
		for i := 0; i <= n-l; i++ {
			j := i + l - 1
			if mp[str[i:j+1]] {
				valid = false
				break
			}
			mp[str[i:j+1]] = true
		}
		if valid {
			fmt.Println(l)
			return
		}
	}
	fmt.Println(n)
}
