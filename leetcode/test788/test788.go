package main

import (
	"fmt"
)

var sums = make([]int, 100001)

func rotatedDigits(n int) int {
	return sums[n]
}

func init() {
	var mp = map[int]int{
		2: 5, 5: 2, 6: 9, 0: 0, 1: 1, 9: 6, 8: 8,
	}
	for i := 1; i <= 10000; i++ {
		t, v := i, 0
		var list = make([]int, 0)
		var isValid = true
		for t != 0 {
			if r, exist := mp[t%10]; exist {
				list = append(list, r)
			} else {
				isValid = false
				break
			}
			t /= 10
		}
		for j := len(list) - 1; j >= 0; j-- {
			v = v*10 + list[j]
		}
		sums[i] = sums[i-1]
		if isValid && v != 0 && v != i {
			sums[i] += 1
		}
	}
}

func main() {
	var s = sums[10]
	fmt.Println(s)
}
