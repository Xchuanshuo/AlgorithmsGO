package main

import (
	"fmt"
	"sort"
)

func main() {
	var n,x,y,z int
	fmt.Scanln(&n, &x, &y, &z)
	var dif = make(map[int]int)
	var arr = []int{0}
	var tmp = map[int]bool{0: true}
	for i := 0;i < n;i++ {
		var l,r int
		fmt.Scanln(&l, &r)
		if _, exist := tmp[l];!exist {
			arr = append(arr, l)
			tmp[l] = true
		}
		if _, exist := tmp[r+1];!exist {
			arr = append(arr, r+1)
			tmp[r+1] = true
		}
		dif[0] += x
		dif[l] += y - x
		dif[r+1] += z - y
	}
	sort.Ints(arr)
	sum, res := 0, 0
	for _, a := range arr {
		sum += dif[a]
		if sum > res { res = sum }
	}
	fmt.Println(res)
}
