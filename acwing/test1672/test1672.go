package main

import "fmt"

func main() {
	var n int
	fmt.Scanf("%d", &n)
	var a, b string
	fmt.Scanf("%s\n%s", &a, &b)
	var arr = make([]int, n)
	for i := 0; i < n; i++ {
		if a[i] != b[i] {
			arr[i] = 1
		}
	}
	last, res := arr[0], 0
	for i := 0; i < n; i++ {
		if (arr[i] == 0 && last == 1) || (arr[i] == 1 && i == n-1) {
			res++
		}
		last = arr[i]
	}
	fmt.Println(res)
}
