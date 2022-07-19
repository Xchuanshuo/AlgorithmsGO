package main

import "fmt"

func main() {
	var n int
	fmt.Scanln(&n)
	l := n + 1
	var a = make([]int, l)
	var b = make([]int, l)
	var c = make([]int, l)
	for i := 1; i <= n; i++ {
		var v int
		fmt.Scanf("%d", &v)
		a[v] = i
	}
	for i := 1; i <= n; i++ {
		var v int
		fmt.Scanf("%d", &v)
		b[i] = v
	}
	for i := 1; i <= n; i++ {
		c[a[a[a[i]]]] = b[i]
	}
	for i := 1; i <= n; i++ {
		fmt.Println(c[i])
	}
}
