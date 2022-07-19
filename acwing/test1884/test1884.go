package main

import "fmt"

func main() {
	var n int
	fmt.Scanf("%d", &n)
	var s string
	fmt.Scanf("%s", &s)
	c, o, w := 0, 0, 0
	for i := 0;i < n;i++ {
		if s[i] == 'C' {
			c++
		} else if s[i] == 'O' {
			if c == 0 { continue }
			o += c
		} else {
			w += o
		}
	}
	fmt.Println(w)
}
