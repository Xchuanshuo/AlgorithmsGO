package main

import "fmt"

func main() {
	var n int
	fmt.Scanln(&n)
	var a = [3][3]int{}
	for i := 0;i < n;i++ {
		var f, s int
		fmt.Scanln(&f, &s)
		a[f-1][s-1]++
	}
	r1 := a[0][1] + a[1][2] + a[2][0]
	r2 := a[2][1] + a[1][0] + a[0][2]
	if r1 < r2 { r1 = r2 }
	fmt.Println(r1)
}