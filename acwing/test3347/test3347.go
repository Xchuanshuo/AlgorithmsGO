package main

import "fmt"

func main() {
	var n int
	fmt.Scanf("%d", &n)
	var a = make([]int, 0)
	for i := 0; i < n; i++ {
		var v int
		fmt.Scanf("%d", &v)
		a = append(a, v)
	}
	var res = 0
	for i := 0; i < n; i++ {
		var mp = make(map[int]bool)
		sum, cnt := 0, 0
		for j := i; j < n; j++ {
			sum += a[j]
			cnt++
			mp[a[j]] = true
			if sum%cnt == 0 && mp[sum/cnt] {
				res++
			}
		}
	}
	fmt.Println(res)
}
