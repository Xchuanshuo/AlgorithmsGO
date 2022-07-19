package main

import "fmt"

func main() {
	var n int
	fmt.Scanf("%d", &n)
	var arr = make([]int, n)
	for i := 0; i < n; i++ {
		var v int
		fmt.Scanf("%d", &v)
		arr[i] = v
	}
	for i := n - 2; i >= 0; i-- {
		if arr[i] > arr[i+1] {
			fmt.Println(i + 1)
			return
		}
	}
	fmt.Println(0)
}
